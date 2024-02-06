package services

import (
	"context"

	"github.com/salmantaghooni/golang-car-web-api/src/api/dto"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/data/db"
	"github.com/salmantaghooni/golang-car-web-api/src/data/models"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
)

type CarModelImageService struct {
	base *BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.UpdateCarModelImageRequest, dto.CarModelImageResponse]
}

func NewCarModelImageService(cfg *config.Config) *CarModelImageService {
	return &CarModelImageService{
		base: &BaseService[models.CarModelImage, dto.CreateCarModelImageRequest, dto.UpdateCarModelImageRequest, dto.CarModelImageResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Image"},
			},
		},
	}
}

// Create
func (s *CarModelImageService) Create(ctx context.Context, req *dto.CreateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *CarModelImageService) Update(ctx context.Context, id int, req *dto.UpdateCarModelImageRequest) (*dto.CarModelImageResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *CarModelImageService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *CarModelImageService) GetById(ctx context.Context, id int) (*dto.CarModelImageResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *CarModelImageService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelImageResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
