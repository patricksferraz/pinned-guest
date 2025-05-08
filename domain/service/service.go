package service

import (
	"context"

	"github.com/patricksferraz/pinned-guest/domain/entity"
	"github.com/patricksferraz/pinned-guest/domain/repo"
	"github.com/patricksferraz/pinned-guest/infra/client/kafka/topic"
	"github.com/patricksferraz/pinned-guest/utils"
)

type Service struct {
	Repo repo.RepoInterface
}

func NewService(repo repo.RepoInterface) *Service {
	return &Service{
		Repo: repo,
	}
}

func (s *Service) CreateGuest(ctx context.Context, name, doc *string) (*string, error) {
	guest, err := entity.NewGuest(name, doc)
	if err != nil {
		return nil, err
	}

	if err = s.Repo.CreateGuest(ctx, guest); err != nil {
		return nil, err
	}

	// TODO: adds retry
	event, err := entity.NewEvent(guest)
	if err != nil {
		return nil, err
	}

	eMsg, err := event.ToJson()
	if err != nil {
		return nil, err
	}

	err = s.Repo.PublishEvent(ctx, utils.PString(topic.NEW_GUEST), utils.PString(string(eMsg)), guest.ID)
	if err != nil {
		return nil, err
	}

	return guest.ID, nil
}

func (s *Service) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	guest, err := s.Repo.FindGuest(ctx, guestID)
	if err != nil {
		return nil, err
	}

	return guest, nil
}

func (s *Service) SearchGuests(ctx context.Context, pageToken *string, pageSize *int) ([]*entity.Guest, *string, error) {
	pagination, err := entity.NewPagination(pageToken, pageSize)
	if err != nil {
		return nil, nil, err
	}

	searchGuests, err := entity.NewSearchGuests(pagination)
	if err != nil {
		return nil, nil, err
	}

	guests, nextPageToken, err := s.Repo.SearchGuests(ctx, searchGuests)
	if err != nil {
		return nil, nil, err
	}

	return guests, nextPageToken, nil
}
