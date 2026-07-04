package request

import (
    "fmt"
    "net/http"
    "strconv"

    core_errors "github.com/cunofou/golang_todoapp/internal/core/errors"
    "github.com/google/uuid"
)

func GetUUIDPathValue(r *http.Request, key string) (uuid.UUID, error) {
    value := GetPathValue(r, key)
    if value == "" {
        return uuid.Nil, fmt.Errorf("path parameter %s is empty: %w", key, core_errors.ErrInvalidArgument)
    }

    id, err := uuid.Parse(value)
    if err != nil {
        return uuid.Nil, fmt.Errorf("invalid UUID format for %s: %w", key, core_errors.ErrInvalidArgument)
    }

    return id, nil
}

func GetIntPathValue(r *http.Request, key string) (int, error) {
    value := GetPathValue(r, key)
    if value == "" {
        return 0, fmt.Errorf("path parameter %s is empty: %w", key, core_errors.ErrInvalidArgument)
    }

    num, err := strconv.Atoi(value)
    if err != nil {
        return 0, fmt.Errorf("invalid integer format for %s: %w", key, core_errors.ErrInvalidArgument)
    }

    return num, nil
}

func GetPathValue(r *http.Request, key string) string {
    return ""
}