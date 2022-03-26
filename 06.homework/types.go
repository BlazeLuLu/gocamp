package main

type Quanzi struct {
	Id      int64   `json:"id,omitempty" gorm:"primaryKey;column:id"`
	Name    string  `json:"name,omitempty" gorm:"column:name"`
	Sex     string  `json:"sex,omitempty" gorm:"column:sex"`
	Tall    float32 `json:"tall,omitempty" gorm:"column:tall"`
	Weight  float32 `json:"weight,omitempty" gorm:"column:weight"`
	Age     int64   `json:"age,omitempty" gorm:"column:age"`
	Fatrate float32 `json:"fatrate,omitempty" gorm:"column:fatrate"`
	Message string  `json:"message,omitempty" gorm:"column:message"`
}

func (*Quanzi) TableName() string {
	return "quanzi"
}
