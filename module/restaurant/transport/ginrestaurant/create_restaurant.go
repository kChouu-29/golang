package ginrestaurant

import (
	restaurantbiz "myCode/module/restaurant/biz"
	restaurantmodel "myCode/module/restaurant/model"
	restaurantstorage "myCode/module/restaurant/storage"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateRestaurant(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data restaurantmodel.ResturantCrete

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLRestore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
