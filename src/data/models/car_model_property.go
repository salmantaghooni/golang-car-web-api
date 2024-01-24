package models

type CarModelProperty struct {
	BaseModel
	CarModel   CarModel `gorm:"foreignKey:CarModelId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	CarModelId int      `gorm:"uniqueIndex:idx_CarModelId_PropertyId"`
	Property   Property `gorm:"foreignKey:PropertyId;constraint:OnUpdate:NO ACTION;OnDelete:NO ACTION"`
	PropertyId int      `gorm:"uniqueIndex:idx_CarModelId_PropertyId"`
	Value      string   `gorm:"size:1000,type:string;not null"`
}
