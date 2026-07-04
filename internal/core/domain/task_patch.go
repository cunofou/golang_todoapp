package domain

import "github.com/cunofou/golang_todoapp/internal/core/transport/http/types"

type TaskPatch struct {
    Title       types.Nullable[string]
    Description types.Nullable[string]
    Completed   types.Nullable[bool]
}

func NewTaskPatch(
    title types.Nullable[string],
    description types.Nullable[string],
    completed types.Nullable[bool],
) TaskPatch {
    return TaskPatch{
        Title:       title,
        Description: description,
        Completed:   completed,
    }
}