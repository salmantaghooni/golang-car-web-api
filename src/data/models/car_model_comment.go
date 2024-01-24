package models

type CarModelComment struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId int
	User       User `gorm:"foreignKey:UserId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	UserId     int
	Message    string `gorm:"size:500,type:string;not null"`
}
