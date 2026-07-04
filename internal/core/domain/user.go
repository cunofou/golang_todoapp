package domain

import (
    "errors"
    "github.com/google/uuid"
)

type User struct {
    ID          uuid.UUID `json:"id"`
    Version     int       `json:"version"`
    FullName    string    `json:"full_name"`
    PhoneNumber *string   `json:"phone_number"`
}

func NewUser(
    id uuid.UUID,
    version int,
    fullName string,
    phoneNumber *string,
) User {
    return User{
        ID:          id,
        Version:     version,
        FullName:    fullName,
        PhoneNumber: phoneNumber,
    }
}

func CreateUser(fullName string, phoneNumber *string) User {
    return User{
        ID:          uuid.New(),
        Version:     1,
        FullName:    fullName,
        PhoneNumber: phoneNumber,
    }
}

func (u *User) Validate() error {
    if len(u.FullName) < 3 || len(u.FullName) > 100 {
        return errors.New("full name must be between 3 and 100 characters")
    }
    return nil
}

func (u *User) ApplyPatch(patch UserPatch) error {
    if patch.FullName.Set {
        if patch.FullName.Value == nil {
            return errors.New("full name cannot be null")
        }
        u.FullName = *patch.FullName.Value
    }
    if patch.PhoneNumber.Set {
        u.PhoneNumber = patch.PhoneNumber.Value
    }
    return nil
}