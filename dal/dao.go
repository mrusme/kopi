package dal

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Row[T any] interface {
	PtrFields() []any
	*T
}

type Column interface {
	~byte | ~int16 | ~int32 | ~int64 | ~float64 |
		~string | ~bool | uuid.UUID | time.Time
}

func Create(ctx context.Context, database *sql.DB, q string, args ...any) (int64, error) {
	res, err := database.ExecContext(ctx, q, args...)
	if err != nil {
		return 0, fmt.Errorf("DAO#Create failed\n%s: %w", q, err)
	}
	return res.LastInsertId()
}

func GetRow[T any, PT Row[T]](ctx context.Context, database *sql.DB, q string, args ...any) (T, error) {
	row := database.QueryRowContext(ctx, q, args...)
	var t T
	ptr := PT(&t)
	if err := row.Scan(ptr.PtrFields()...); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return t, ErrNotFound
		}
		return t, fmt.Errorf("DAO#GetRow row.Scan error\n%s: %w", q, err)
	}
	return t, nil
}

func GetColumn[T Column](ctx context.Context, database *sql.DB, q string, args ...any) (T, error) {
	row := database.QueryRowContext(ctx, q, args...)
	var t T
	if err := row.Scan(&t); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return t, ErrNotFound
		}
		return t, fmt.Errorf("DAO#GetColumn row.Scan error\n%s: %w", q, err)
	}
	return t, nil
}

func FindRows[T any, PT Row[T]](ctx context.Context, database *sql.DB, q string, args ...any) ([]T, error) {
	rows, err := database.QueryContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("DAO#FindRows QueryContext failed\n%s: %w", q, err)
	}
	defer func() { _ = rows.Close() }()

	var result []T
	for rows.Next() {
		var t T
		ptr := PT(&t)
		if err := rows.Scan(ptr.PtrFields()...); err != nil {
			return nil, fmt.Errorf("DAO#FindRows row.Scan error\n%s: %w", q, err)
		}
		result = append(result, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("DAO#FindRows rows.Err()\n%s: %w", q, err)
	}
	return result, nil
}

func FindColumns[T Column](ctx context.Context, database *sql.DB, q string, args ...any) ([]T, error) {
	rows, err := database.QueryContext(ctx, q, args...)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("DAO#FindColumns QueryContext failed\n%s: %w", q, err)
	}
	defer func() { _ = rows.Close() }()

	var result []T
	for rows.Next() {
		var t T
		if err := rows.Scan(&t); err != nil {
			return nil, fmt.Errorf("DAO#FindColumns row.Scan error\n%s: %w", q, err)
		}
		result = append(result, t)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("DAO#FindColumns rows.Err()\n%s: %w", q, err)
	}
	return result, nil
}

func InArgs[T Column](tt []T) (string, []any) {
	args := make([]any, len(tt))
	for i, t := range tt {
		args[i] = t
	}
	return strings.Repeat("?,", len(args)-1) + "?", args
}
