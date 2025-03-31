package restaurantstorage

import (
	"context"
	restaurantmodel "myCode/module/restaurant/model"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	morekey string,
) (*restaurantmodel.Resturant, error) { //Cần con trỏ ở đây vì sẽ mất memory nếu không có con trỏ
	var data restaurantmodel.Resturant
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}
