package restaurantstorage

import (
	"context"
	restaurantmodel "myCode/module/restaurant/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error { //Cần con trỏ ở đây vì sẽ mất memory nếu không có con trỏ

	if err := s.db.Table(restaurantmodel.Resturant{}.TableName()).
		Where("id=?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}

	return nil
}
