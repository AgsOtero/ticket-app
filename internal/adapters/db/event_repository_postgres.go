package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/AgsOtero/event-ticket-api/internal/core/domain"
	"github.com/AgsOtero/event-ticket-api/internal/core/ports"
)

type postgresEventRepository struct {
	db *sql.DB
}

func NewPostgresEventRepository(db *sql.DB) ports.EventRepository {
	return &postgresEventRepository{
		db: db,
	}
}

func (p *postgresEventRepository) Save(ctx context.Context, event domain.Event) (domain.Event, error) {
	query := `INSERT INTO events (name,artist,date_time,place_id) 
    		VALUES ($1,$2,$3,$4)
    		RETURNING id`
	err := p.db.QueryRowContext(ctx, query, event.Name, event.Artist, event.DateTime, event.PlaceID).Scan(&event.ID)
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (p *postgresEventRepository) FindByID(ctx context.Context, id int64) (domain.Event, error) {
	query := `SELECT id,name,artist,date_time,place_id
FROM events 
WHERE id=$1`

	var event domain.Event
	row := p.db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&event.ID,
		&event.Name,
		&event.Artist,
		&event.DateTime,
		&event.PlaceID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Event{}, nil
		}
		return domain.Event{}, err
	}

	return event, nil
}

func (p *postgresEventRepository) FindAll(ctx context.Context) ([]domain.Event, error) {
	query := `SELECT id,name,artist,date_time,place_id FROM events`
	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []domain.Event

	for rows.Next() {
		var event domain.Event
		err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Artist,
			&event.DateTime,
			&event.PlaceID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return events, nil
}
