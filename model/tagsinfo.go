package model

type TagsInfo struct {
	Tid    int    `gorm:"type:int;primary_key"`
	Class  string `gorm:"type:varchar(255)"`
	Bg     string `gorm:"type:varchar(255)"`
	Dob    string `gorm:"type:varchar(255)"`
	Gender int    `gorm:"type:int"`
}
