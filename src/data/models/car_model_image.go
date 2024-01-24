package models

type CarModelImage struct {
	BaseModel
	CarModel    CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId  int      `gorm:"uniqueIndex:idx_CarModelId_ImageId"`
	Image       File     `gorm:"foreignKey:ImageId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	ImageId     int      `gorm:"uniqueIndex:idx_CarModelId_ImageId"`
	IsMainImage bool
}
