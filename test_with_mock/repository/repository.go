package repository

import (
	"context"
	"github.com/atadzan/learn-go-testing/test_with_mock/models"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{db: pool}
}

type OrderRepo interface {
	Save(ctx context.Context, log *logrus.Logger, order *models.Order) error
	Get(ctx context.Context, log *logrus.Logger, ids []uint) (map[uint]models.Order, error)
}
