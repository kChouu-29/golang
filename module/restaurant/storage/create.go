package restaurantstorage

import (
	"context"
	restaurantmodel "myCode/module/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.ResturantCrete) error {
	s.db.Create(&data)

	return nil
}
