package useCase

import (
	"context"
	"fmt"
	"github.com/atadzan/learn-go-testing/test_with_mock/models"
	"github.com/atadzan/learn-go-testing/test_with_mock/repository"
	"github.com/sirupsen/logrus"
)

type UseCase struct {
	repo repository.OrderRepo
}

func New(repo repository.OrderRepo) *UseCase {
	return &UseCase{repo: repo}
}

func (s *UseCase) Save(ctx context.Context, log *logrus.Logger, order *models.Order) error {
	return s.repo.Save(ctx, log, order)
}

func (s *UseCase) Get(ctx context.Context, log *logrus.Logger, ids []uint) ([]models.Order, error) {
	ordersMap, err := s.repo.Get(ctx, log, ids)
	if err != nil {
		return nil, fmt.Errorf("err from orders_repository: %s", err.Error())
	}

	result := make([]models.Order, 0, len(ordersMap))
	for _, order := range ordersMap {
		result = append(result, order)
	}
	fmt.Println(result)
	return result, nil
}
