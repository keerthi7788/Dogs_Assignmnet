package service

import (
	"context"
	"dogs-service/models"
)

type DogRepositoryInterface interface {
	CreateDog(ctx context.Context, dog models.Dog) error
	GetAllDogs(ctx context.Context) ([]models.Dog, error)
	GetByID(ctx context.Context, id int) (models.Dog, error)
	Update(ctx context.Context, dog models.Dog) error
	Delete(ctx context.Context, id int) error
}

type DogService struct {
	repo DogRepositoryInterface
}

func NewDogService(repo DogRepositoryInterface) *DogService {
	return &DogService{repo: repo}
}

func (s *DogService) CreateDog(ctx context.Context, dog models.Dog) error {
	return s.repo.CreateDog(ctx, dog)
}

func (s *DogService) GetDogs(ctx context.Context) ([]models.Dog, error) {
	return s.repo.GetAllDogs(ctx)
}
func (s *DogService) GetDog(ctx context.Context, id int) (models.Dog, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *DogService) Update(ctx context.Context, dog models.Dog) error {
	return s.repo.Update(ctx, dog)
}

func (s *DogService) Delete(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
