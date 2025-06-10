package ports

import (
	"context"

	"github.com/AgsOtero/event-ticket-api/internal/core/domain"
)

type UserRepository interface {
	Save(ctx context.Context, user domain.User) (domain.User, error)
	FindByID(ctx context.Context, id int64) (domain.User, error)
	FindByEmail(ctx context.Context, email string) (domain.User, error)
}
type EventRepository interface {
	Save(ctx context.Context, event domain.Event) (domain.Event, error)
	FindByID(ctx context.Context, id int64) (domain.Event, error)
	FindAll(ctx context.Context) ([]domain.Event, error)
}
