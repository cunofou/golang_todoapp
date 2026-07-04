package users_service

import (
	"context"
	"fmt"

	"github.com/cunofou/golang-todoapp/internal/core/domain"
	"github.com/google/uuid"
)

// GetUser возвращает пользователя по ID, делегируя запрос репозиторию.
func (s *UsersService) GetUser(
	ctx context.Context,
	id uuid.UUID,
) (domain.User, error) {
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("get user from repository: %w", err)
	}

	return user, nil
}
