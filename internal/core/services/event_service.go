package services

import (
	"context"
	"errors"

	"github.com/AgsOtero/event-ticket-api/internal/core/domain"
	"github.com/AgsOtero/event-ticket-api/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

type eventService struct {
	eventRepository ports.EventRepository
}

func NewEventService(eventRepository ports.EventRepository) ports.EventService {
	return &eventService{eventRepository: eventRepository}
}

func (e eventService) CreateEvent(ctx context.Context, name, artist, dateTime, placeId string) (domain.Event, error) {
	existingEvent, err := e.
		panic("implement me")
}
