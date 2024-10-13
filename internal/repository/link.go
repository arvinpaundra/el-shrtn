package repository

import (
	"context"
	"errors"

	"github.com/arvinpaundra/el-shrtn/internal/model"
	"github.com/arvinpaundra/el-shrtn/pkg/constant"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LinkRepository struct {
	coll *mongo.Collection
}

func NewLinkRepository(db *mongo.Database) *LinkRepository {
	return &LinkRepository{
		coll: db.Collection("links"),
	}
}

func (r *LinkRepository) Insert(ctx context.Context, link model.Link, opts ...*options.InsertOneOptions) error {
	_, err := r.coll.InsertOne(ctx, link, opts...)
	if err != nil {
		return err
	}

	return nil
}

func (r *LinkRepository) FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) (model.Link, error) {
	var link model.Link

	err := r.coll.FindOne(ctx, filter, opts...).Decode(&link)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Link{}, constant.ErrLinkNotFound
		}

		return model.Link{}, err
	}

	return link, nil
}

func (r *LinkRepository) Update(ctx context.Context, filter any, link interface{}, opts ...*options.UpdateOptions) error {
	_, err := r.coll.UpdateOne(ctx, filter, link, opts...)
	if err != nil {
		return err
	}

	return nil
}
