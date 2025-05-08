package repo

import (
	"context"

	"github.com/patricksferraz/pinned-guest/domain/entity"
)

type RepoInterface interface {
	CreateGuest(ctx context.Context, guest *entity.Guest) error
	FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error)
	SaveGuest(ctx context.Context, guest *entity.Guest) error
	SearchGuests(ctx context.Context, searchGuest *entity.SearchGuests) ([]*entity.Guest, *string, error)

	PublishEvent(ctx context.Context, topic, msg, key *string) error
}
