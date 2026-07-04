package domain

import "github.com/cunofou/golang_todoapp/internal/core/transport/http/types"

type UserPatch struct {
    FullName    types.Nullable[string]
    PhoneNumber types.Nullable[string]
}

func NewUserPatch(
    fullName types.Nullable[string],
    phoneNumber types.Nullable[string],
) UserPatch {
    return UserPatch{
        FullName:    fullName,
        PhoneNumber: phoneNumber,
    }
}