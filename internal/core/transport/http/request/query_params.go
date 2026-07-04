package request

import (
    "fmt"
    "net/http"
    "strconv"
    "time"

    core_errors "github.com/cunofou/golang_todoapp/internal/core/errors"
    "github.com/google/uuid"
)

func GetUUIDQueryParam(r *http.Request, key string) (uuid.UUID, error) {
    value := r.URL.Query().Get(key)
    if value == "" {
        return uuid.Nil, nil
    }

    id, err := uuid.Parse(value)
    if err != nil {
        return uuid.Nil, fmt.Errorf("invalid UUID format for query param %s: %w", key, core_errors.ErrInvalidArgument)
    }

    return id, nil
}

func GetIntQueryParam(r *http.Request, key string) (int, error) {
    value := r.URL.Query().Get(key)
    if value == "" {
        return 0, nil
    }

    num, err := strconv.Atoi(value)
    if err != nil {
        return 0, fmt.Errorf("invalid integer format for query param %s: %w", key, core_errors.ErrInvalidArgument)
    }

    return num, nil
}

func GetDateQueryParam(r *http.Request, key string) (time.Time, error) {
    value := r.URL.Query().Get(key)
    if value == "" {
        return time.Time{}, nil
    }

    t, err := time.Parse(time.RFC3339, value)
    if err != nil {
        return time.Time{}, fmt.Errorf("invalid date format for query param %s (expected RFC3339): %w", key, core_errors.ErrInvalidArgument)
    }

    return t, nil
}