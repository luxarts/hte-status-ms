package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"hte-status-ms/internal/defines"
	"hte-status-ms/internal/domain"
)

type StatusRepository interface {
	Update(p *domain.Status) error
}

type statusRepository struct {
	db             *mongo.Database
	collectionName string
}

func NewStatusRepository(client *mongo.Client) StatusRepository {
	return &statusRepository{
		db:             client.Database(defines.DatabasesHTE),
		collectionName: "status",
	}
}

func (r *statusRepository) Update(s *domain.Status) error {
	sBson, err := bson.Marshal(s)
	if err != nil {
		return err
	}

	ctx := context.Background()
	upsert := true
	_, err = r.db.Collection(r.collectionName).ReplaceOne(ctx, bson.M{"device_id": s.DeviceID}, sBson, &options.ReplaceOptions{Upsert: &upsert})
	return err
}
