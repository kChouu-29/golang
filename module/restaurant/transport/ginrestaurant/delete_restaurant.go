package ginrestaurant

import (
	restaurantbiz "myCode/module/restaurant/biz"
	restaurantstorage "myCode/module/restaurant/storage"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteRestaurant(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// Lấy ID từ URL
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID",
			})
			return
		}

		// Tạo storage để thao tác database
		store := restaurantstorage.NewSQLRestore(db)

		// Khởi tạo business logic để xóa nhà hàng
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		// Xóa nhà hàng
		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Trả về kết quả thành công
		c.JSON(http.StatusOK, gin.H{
			"data": 1,
		})
	}
}
