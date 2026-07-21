package ocr

import (
	"image"
	"image/png"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/spf13/viper"
)

func TestGetDataFromPhotoMockBackend(t *testing.T) {
	var gotBody string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		gotBody = string(b)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, `{
			"id": "chatcmpl-mock",
			"object": "chat.completion",
			"created": 0,
			"model": "llama3.2-vision",
			"choices": [{
				"index": 0,
				"message": {
					"role": "assistant",
					"content": "[{\"coffee\":\"La Loma\",\"roaster\":\"Glitch Coffee\",\"drink\":\"Cappuccino\",\"rating\":\"5/5\"}]"
				},
				"finish_reason": "stop"
			}]
		}`)
	}))
	defer srv.Close()

	dir := t.TempDir()
	photo := filepath.Join(dir, "sample.png")
	f, err := os.Create(photo)
	if err != nil {
		t.Fatal(err)
	}
	if err := png.Encode(f, image.NewRGBA(image.Rect(0, 0, 4, 4))); err != nil {
		t.Fatal(err)
	}
	f.Close()

	viper.Reset()
	defer viper.Reset()
	viper.Set("LLM.Enabled", true)
	viper.Set("LLM.Host", srv.URL+"/v1")
	viper.Set("LLM.APIKey", "test")
	viper.Set("LLM.Model", "llama3.2-vision")

	od, err := GetDataFromPhoto(photo)
	if err != nil {
		t.Fatal(err)
	}

	if len(od) != 1 {
		t.Fatalf("expected 1 record, got %d", len(od))
	}
	if od[0].Coffee != "La Loma" || od[0].Roaster != "Glitch Coffee" {
		t.Fatalf("unexpected parse: %+v", od[0])
	}
	if !strings.Contains(gotBody, "image_url") || !strings.Contains(gotBody, "data:image/png;base64,") {
		t.Fatalf("request did not carry a data-URI image part:\n%s", gotBody)
	}
	t.Logf("request body carried image part; parsed %+v", od[0])
}
