package useCase

import (
	"context"
	"errors"
	"fmt"
	mock_repository "github.com/atadzan/learn-go-testing/test_with_mock/mocks"
	"github.com/atadzan/learn-go-testing/test_with_mock/models"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_repository.NewMockOrderRepo(ctl)

	ctx := context.Background()
	log := logrus.New()
	in := []uint{1, 2, 3}
	mockResp := map[uint]models.Order{
		1: {
			Id:          1,
			PaymentType: 1,
			Items: []models.Item{
				{Id: 1, OrderId: 2, Amount: 4},
				{Id: 2, OrderId: 3, Amount: 4},
			},
		},
		2: {
			Id:          2,
			PaymentType: 2,
			Items: []models.Item{
				{Id: 2, OrderId: 3, Amount: 5},
			},
		},
	}
	expected := []models.Order{
		{
			Id:          1,
			PaymentType: 1,
			Items: []models.Item{
				{Id: 1, OrderId: 2, Amount: 4},
				{Id: 2, OrderId: 3, Amount: 4},
			},
		},
		{
			Id:          2,
			PaymentType: 2,
			Items: []models.Item{
				{Id: 2, OrderId: 3, Amount: 5},
			},
		},
	}
	repo.EXPECT().Get(ctx, log, in).Return(mockResp, nil).Times(1)

	useCase := New(repo)
	orders, err := useCase.Get(ctx, log, in)

	require.NoError(t, err)
	require.ElementsMatch(t, expected, orders)
}

func TestGetError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_repository.NewMockOrderRepo(ctl)

	repoErr := errors.New("db is down")
	ctx := context.Background()
	log := logrus.New()
	in := []uint{1, 2, 3}
	repo.EXPECT().Get(ctx, log, in).Return(nil, repoErr).Times(1)

	useCase := New(repo)
	orders, err := useCase.Get(ctx, log, in)
	require.Error(t, err)
	require.EqualError(t, fmt.Errorf("err from orders_repository: %s", repoErr.Error()), err.Error())
	require.Nil(t, orders)
}

func TestSave(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	repo := mock_repository.NewMockOrderRepo(ctl)

	ctx := context.Background()
	log := logrus.New()
	in := &models.Order{}
	repo.EXPECT().Save(ctx, log, in).Return(nil).Times(1)

	useCase := New(repo)
	err := useCase.Save(ctx, log, in)
	require.NoError(t, err)
}
