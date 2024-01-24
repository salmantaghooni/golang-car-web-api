package models

type CarModelYear struct {
	BaseModel
	CarModel               CarModel    `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId             int         `gorm:"uniqueIndex:idx_CarModelId_PersianYearId"`
	PersianYear            PersianYear `gorm:"foreignKey:PersianYearId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	PersianYearId          int         `gorm:"uniqueIndex:idx_CarModelId_PersianYearId"`
	CarModelPriceHistories []CarModelPriceHistory
}
