package models

import "time"

type CarModelPriceHistory struct {
	BaseModel
	CarModelYear   CarModelYear `gorm:"foreignKey:CarModelYearId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelYearId int
	Price          float64   `gorm:"type:decimal(10,2);not null"`
	PriceAt        time.Time `gorm:"type:TIMESTAMP with time zone;not null"`
}
