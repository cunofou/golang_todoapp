package request

import (
    "encoding/json"
    "fmt"
    "net/http"

    core_errors "github.com/cunofou/golang_todoapp/internal/core/errors"
    "github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Validatable interface {
    Validate() error
}

func DecodeAndValidateRequest(r *http.Request, dest any) error {
    if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
        return fmt.Errorf("decode json: %v: %w", err, core_errors.ErrInvalidArgument)
    }

    var err error
    v, ok := dest.(Validatable)
    if ok {
        err = v.Validate()
    } else {
        err = validate.Struct(dest)
    }

    if err != nil {
        return fmt.Errorf("request validation: %v: %w", err, core_errors.ErrInvalidArgument)
    }

    return nil
}