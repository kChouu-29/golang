package biz

import (
	"context"
	"errors"
	restaurantmodel "myCode/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.ResturantCrete) error
}

type CreateRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *CreateRestaurantBiz {
	return &CreateRestaurantBiz{store: store}
}

func (biz *CreateRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.ResturantCrete) error {
	if data.Name == "" {
		return errors.New("name cannot empty")
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
