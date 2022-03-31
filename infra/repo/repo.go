package repo

import (
	"context"
	"fmt"

	"github.com/c-4u/pinned-guest/domain/entity"
	"github.com/c-4u/pinned-guest/infra/client/kafka"
	"github.com/c-4u/pinned-guest/infra/db"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

type Repository struct {
	Orm *db.DbOrm
	Kp  *kafka.KafkaProducer
}

func NewRepository(orm *db.DbOrm, kp *kafka.KafkaProducer) *Repository {
	return &Repository{
		Orm: orm,
		Kp:  kp,
	}
}

func (r *Repository) CreateGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Orm.Db.Create(guest).Error
	return err
}

func (r *Repository) FindGuest(ctx context.Context, guestID *string) (*entity.Guest, error) {
	var e entity.Guest

	r.Orm.Db.First(&e, "id = ?", *guestID)

	if e.ID == nil {
		return nil, fmt.Errorf("no guest found")
	}

	return &e, nil
}

func (r *Repository) SaveGuest(ctx context.Context, guest *entity.Guest) error {
	err := r.Orm.Db.Save(guest).Error
	return err
}

func (r *Repository) SearchGuests(ctx context.Context, searchGuest *entity.SearchGuests) ([]*entity.Guest, *string, error) {
	var e []*entity.Guest

	q := r.Orm.Db
	if *searchGuest.PageToken != "" {
		q = q.Where("token < ?", *searchGuest.PageToken)
	}
	err := q.Order("token DESC").
		Limit(*searchGuest.PageSize).
		Find(&e).Error
	if err != nil {
		return nil, nil, err
	}

	if len(e) == 0 {
		return nil, nil, nil
	}

	return e, e[len(e)-1].Token, nil
}

func (r *Repository) PublishEvent(ctx context.Context, topic, msg, key *string) error {
	message := &ckafka.Message{
		TopicPartition: ckafka.TopicPartition{Topic: topic, Partition: ckafka.PartitionAny},
		Value:          []byte(*msg),
		Key:            []byte(*key),
	}
	err := r.Kp.Producer.Produce(message, r.Kp.DeliveryChan)
	if err != nil {
		return err
	}
	return nil
}
