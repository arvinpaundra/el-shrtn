package link

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arvinpaundra/el-shrtn/config"
	"github.com/arvinpaundra/el-shrtn/internal/dto/request"
	"github.com/arvinpaundra/el-shrtn/internal/dto/response"
	"github.com/arvinpaundra/el-shrtn/internal/factory"
	"github.com/arvinpaundra/el-shrtn/internal/model"
	"github.com/arvinpaundra/el-shrtn/pkg/constant"
	"github.com/arvinpaundra/el-shrtn/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type LinkWriterRepository interface {
}

type LinkRepository interface {
	Insert(ctx context.Context, link model.Link, opts ...*options.InsertOneOptions) error
	FindOne(ctx context.Context, filter any, opts ...*options.FindOneOptions) (model.Link, error)
	Update(ctx context.Context, filter any, link interface{}, opts ...*options.UpdateOptions) error
	// FindAll(ctx context.Context, filter any, opts ...*options.FindOptions) ([]model.Link, error)
}

type Service struct {
	linkRepository LinkRepository
	logger         *zap.Logger
}

func NewService(f *factory.Factory) *Service {
	return &Service{
		linkRepository: f.LinkRepository,
		logger:         f.Logger.With(zap.String("domain", "link")),
	}
}

func (s *Service) CreateLink(ctx context.Context, payload request.CreateLink) (response.CreatedLink, error) {
	code, err := s.generateCodeLink(ctx, util.RandomString(8))
	if err != nil {
		s.logger.Error(err.Error())
		return response.CreatedLink{}, err
	}

	link := model.Link{
		ID:            primitive.NewObjectID(),
		OriginLink:    payload.OriginLink,
		ShortenedLink: fmt.Sprintf("%s/%s", config.GetBaseUrlApp(), code),
		Code:          code,
		VisitorAmount: 0,
		Status:        "active",
		ExpiredAt:     s.parseExpiredLink(payload.Duration),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err = s.linkRepository.Insert(ctx, link)
	if err != nil {
		s.logger.With(zap.Any("link_payload", link)).Error(err.Error())
		return response.CreatedLink{}, err
	}

	var expiredAt *string

	if link.ExpiredAt != nil {
		expiredAt = util.Address(link.ExpiredAt.Format(time.DateTime))
	}

	res := response.CreatedLink{
		ID:            link.ID,
		ShortenedLink: link.ShortenedLink,
		ExpiredAt:     expiredAt,
	}

	return res, nil
}

// TODO: refactor recursive with another approach
func (s *Service) generateCodeLink(ctx context.Context, code string) (string, error) {
	filter := bson.M{"code": code, "status": "active"}

	link, err := s.linkRepository.FindOne(ctx, filter)
	if err != nil && !errors.Is(err, constant.ErrLinkNotFound) {
		return "", err
	}

	if link == (model.Link{}) {
		return code, nil
	}

	code = util.RandomString(8)

	return s.generateCodeLink(ctx, code)
}

func (s *Service) parseExpiredLink(pattern string) *time.Time {
	now := time.Now()

	var parsedTime time.Time

	switch pattern {
	case "30m":
		parsedTime = now.Add(time.Minute * 30)
		return &parsedTime
	case "1h":
		parsedTime = now.Add(time.Hour * 1)
		return &parsedTime
	case "12h":
		parsedTime = now.Add(time.Hour * 12)
		return &parsedTime
	case "1d":
		parsedTime = now.AddDate(0, 0, 1)
		return &parsedTime
	case "5d":
		parsedTime = now.AddDate(0, 0, 5)
		return &parsedTime
	case "30d":
		parsedTime = now.AddDate(0, 0, 30)
		return &parsedTime
	}

	return nil
}

func (s *Service) AccessLink(ctx context.Context, code string) (string, error) {
	shortenedLink := fmt.Sprintf("%s/%s", config.GetBaseUrlApp(), code)

	filterLink := bson.M{
		"shortened_link": shortenedLink,
	}

	link, err := s.linkRepository.FindOne(ctx, filterLink)
	if err != nil {
		s.logger.Error(err.Error())
		return "", err
	}

	if link.Status == "expired" {
		s.logger.With(
			zap.String("shortened_link", link.ShortenedLink),
			zap.String("origin_link", link.OriginLink),
		).Info(constant.ErrLinkExpired.Error())

		return "", constant.ErrLinkExpired
	}

	now := time.Now()

	if link.ExpiredAt != nil && now.After(*link.ExpiredAt) {
		updatedLink := bson.M{
			"$set": bson.M{
				"status":     "expired",
				"updated_at": now,
			},
		}

		filterUpdateLink := bson.M{
			"_id": link.ID,
		}

		err = s.linkRepository.Update(ctx, filterUpdateLink, updatedLink)
		if err != nil {
			s.logger.With(zap.Any("updated_link_payload", updatedLink)).Error(err.Error())
			return "", err
		}

		return "", constant.ErrLinkExpired
	}

	updatedLink := bson.M{
		"$set": bson.M{
			"updated_at": now,
		},
		"$inc": bson.M{
			"visitor_amount": 1,
		},
	}

	filterUpdateLink := bson.M{
		"_id": link.ID,
	}

	err = s.linkRepository.Update(ctx, filterUpdateLink, updatedLink)
	if err != nil {
		s.logger.With(zap.Any("updated_link_payload", updatedLink)).Error(err.Error())
		return "", err
	}

	return link.OriginLink, nil
}
