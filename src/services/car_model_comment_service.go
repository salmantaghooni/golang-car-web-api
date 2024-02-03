package services

import (
	"context"

	"github.com/salmantaghooni/golang-car-web-api/src/api/dto"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/constants"
	"github.com/salmantaghooni/golang-car-web-api/src/data/db"
	"github.com/salmantaghooni/golang-car-web-api/src/data/models"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
)

type CarModelCommentService struct {
	base *BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]
}

func NewCarModelCommentService(cfg *config.Config) *CarModelCommentService {
	return &CarModelCommentService{
		base: &BaseService[models.CarModelComment, dto.CreateCarModelCommentRequest, dto.UpdateCarModelCommentRequest, dto.CarModelCommentResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "User"},
			},
		},
	}
}

// Create
func (s *CarModelCommentService) Create(ctx context.Context, req *dto.CreateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	req.UserId = int(ctx.Value(constants.UserIdKey).(float64))
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelCommentService) Update(ctx context.Context, id int, req *dto.UpdateCarModelCommentRequest) (*dto.CarModelCommentResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelCommentService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelCommentService) GetById(ctx context.Context, id int) (*dto.CarModelCommentResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelCommentService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelCommentResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
