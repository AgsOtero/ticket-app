package ports

import (
	"context"

	"github.com/AgsOtero/event-ticket-api/internal/core/domain"
)

type UserService interface {
	Register(ctx context.Context, email, password, name, surname, phone string) (domain.User, error)
	GetById(ctx context.Context, id int64) (domain.User, error)
}

type EventService interface {
	CreateEvent(ctx context.Context, name, artist, dateTime, placeId string) (domain.Event, error)
	GetById(ctx context.Context, id int64) (domain.User, error)
}
