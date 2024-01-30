package services

import (
	"context"

	"github.com/salmantaghooni/golang-car-web-api/src/api/dto"
	"github.com/salmantaghooni/golang-car-web-api/src/config"
	"github.com/salmantaghooni/golang-car-web-api/src/data/db"
	"github.com/salmantaghooni/golang-car-web-api/src/data/models"
	"github.com/salmantaghooni/golang-car-web-api/src/pkg/logging"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: &BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{string: "Category"}},
		},
	}
}

// Create
func (s *PropertyService) Create(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update
func (s *PropertyService) Update(ctx context.Context, id int, req *dto.UpdatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete
func (s *PropertyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// Get By Id
func (s *PropertyService) GetById(ctx context.Context, id int) (*dto.PropertyResponse, error) {
	return s.base.GetById(ctx, id)
}

// Get By Filter
func (s *PropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}