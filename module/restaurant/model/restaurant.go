package restaurantmodel

type Resturant struct {
	Id   int    `json:"id" gorm:"column:id;"` // ten bien phai viet hoa
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Resturant) TableName() string {
	return "restaurants"
}

type ResturantCrete struct {
	Id   int    `json:"id" gorm:"column:id;"` // ten bien phai viet hoa
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (ResturantCrete) TableName() string {
	return Resturant{}.TableName()
}

type ResturantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr string  `json:"addr" gorm:"column:addr"`
}
