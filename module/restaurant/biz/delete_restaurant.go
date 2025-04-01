package biz

import (
	"context"
	"errors"
	restaurantmodel "myCode/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FindDataWithCondition(
		ctx context.Context,
		condition map[string]interface{},
		morekey string,
	) (*restaurantmodel.Resturant, error)
	Delete(ctx context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id}, "") //Lấy dữ liệu cũ ra để kiểm tra xem có tồn tại hay không

	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("restaurant not found")
	}
	return nil
}
