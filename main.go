package main

import (
	"log"
	"myCode/module/restaurant/transport/ginrestaurant"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Resturant struct {
	Id   int    `json:"id" gorm:"column:id;"` // ten bien phai viet hoa
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Resturant) TableName() string {
	return "restaurants"
}

type ResturantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr string  `json:"addr" gorm:"column:addr"`
}

func (ResturantUpdate) TableName() string {
	return Resturant{}.TableName()
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3307)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	err := godotenv.Load()
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

		log.Fatalln(err)
	}
	log.Println(db)
	r := gin.Default() // đi lấy 1 server

	r.GET("/ping", func(c *gin.Context) { // đăng ký 1 đường đi trên server có dạng /ping

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// POST /restaurants
	v1 := r.Group("/v1")

	restaurants := v1.Group("/restaurants")

	restaurants.POST("", ginrestaurant.CreateRestaurant(db)) //func(c *gin.Context) {
	// 	var data Resturant

	// 	if err := c.ShouldBind(&data); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	db.Create(&data)

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": data,
	// 	})
	// })
	//GET
	// restaurants.GET("/:id", func(c *gin.Context) {

	// 	id, err := strconv.Atoi(c.Param("id"))

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	var data Resturant
	// 	db.Where("id =?", id).First(&data)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": data,
	// 	})
	// })
	restaurants.GET("", func(c *gin.Context) {
		var data []Resturant

		type Paging struct {
			Page  int `json:"page" form:"page"`
			Limit int `json:"limit" form:"limit"`
		}

		var paggingData Paging

		if err := c.ShouldBind(&paggingData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if paggingData.Page <= 0 {
			paggingData.Page = 1
		}
		if paggingData.Limit <= 0 {
			paggingData.Limit = 5
		}
		//Bỏ qua số dòng và lấy giới hạn dòng tiếp theo
		//ví dụ limit=3 và page=1 -> bỏ qua 0 dòng và lấy 3 dòng từ trên xuông
		db.Offset((paggingData.Page - 1) * paggingData.Limit).Order("id desc").Limit(paggingData.Limit).Find(&data)

		// if err := db.Find(&data).Error; err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": "Cannot fetch restaurants",
		// 	})
		// 	return
		// }
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})
	//PATCH

	restaurants.PATCH("/:id", func(c *gin.Context) {

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		var data ResturantUpdate
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		db.Where("id =?", id).Updates(&data)
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	})

	//DELETE
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db)) //func(c *gin.Context) {

	// 	id, err := strconv.Atoi(c.Param("id"))

	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	db.Table(Resturant{}.TableName()).Where("id =?", id).Delete(nil)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"data": 1,
	// 	})
	// })
	r.Run()
	// newRestaurant := Resturant{Name: "Tani", Addr: "9 Duong Lang"}

	// db.Create(&newRestaurant)

	// log.Println("New id:", newRestaurant.Id)
	// //Lay ra 1 du lieu co dieu kien
	// var myRestaurant Resturant
	// if err := db.Where("id = ?", 2).First(&myRestaurant).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(myRestaurant)
	// //Cap nhat du lieu co dieu kien
	// newName := ""
	// updateData := ResturantUpdate{Name: &newName}
	// if err := db.Where("id = ?", 1).Updates(&updateData).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(myRestaurant)
	// Xoas du lieu co dieu kien
	// if err := db.Table(Resturant{}.TableName()).Where("id = ?", 2).Delete(nil).Error; err != nil {
	// 	log.Println(err)
	// }
	// log.Println(myRestaurant)
}
