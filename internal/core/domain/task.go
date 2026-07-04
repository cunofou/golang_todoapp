package domain

import (
    "errors"
    "time"
    "github.com/google/uuid"
)

type Task struct {
    ID           uuid.UUID  `json:"id"`
    Version      int        `json:"version"`
    Title        string     `json:"title"`
    Description  *string    `json:"description"`
    Completed    bool       `json:"completed"`
    CreatedAt    time.Time  `json:"created_at"`
    CompletedAt  *time.Time `json:"completed_at"`
    AuthorUserID uuid.UUID  `json:"author_user_id"`
}

func NewTask(
    id uuid.UUID,
    version int,
    title string,
    description *string,
    completed bool,
    createdAt time.Time,
    completedAt *time.Time,
    authorUserID uuid.UUID,
) Task {
    return Task{
        ID:           id,
        Version:      version,
        Title:        title,
        Description:  description,
        Completed:    completed,
        CreatedAt:    createdAt,
        CompletedAt:  completedAt,
        AuthorUserID: authorUserID,
    }
}

func CreateTask(title string, description *string, authorUserID uuid.UUID) Task {
    now := time.Now().UTC()
    return Task{
        ID:           uuid.New(),
        Version:      1,
        Title:        title,
        Description:  description,
        Completed:    false,
        CreatedAt:    now,
        CompletedAt:  nil,
        AuthorUserID: authorUserID,
    }
}

func (t *Task) Validate() error {
    if len(t.Title) < 1 || len(t.Title) > 100 {
        return errors.New("title must be between 1 and 100 characters")
    }
    return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
    if patch.Title.Set {
        if patch.Title.Value == nil {
            return errors.New("title cannot be null")
        }
        t.Title = *patch.Title.Value
    }
    if patch.Description.Set {
        t.Description = patch.Description.Value
    }
    if patch.Completed.Set {
        if patch.Completed.Value == nil {
            return errors.New("completed cannot be null")
        }
        t.Completed = *patch.Completed.Value
        if t.Completed {
            now := time.Now().UTC()
            t.CompletedAt = &now
        } else {
            t.CompletedAt = nil
        }
    }
    return nil
}

func (t *Task) CompletionDuration() *time.Duration {
    if !t.Completed || t.CompletedAt == nil {
        return nil
    }
    duration := t.CompletedAt.Sub(t.CreatedAt)
    return &duration
}