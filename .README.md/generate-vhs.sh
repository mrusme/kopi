#/bin/sh

/bin/ls -1 ./*.tape | while read tape; do vhs "$tape"; done
