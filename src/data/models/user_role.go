package models

type UserRole struct {
	BaseModel
	User   User `gorm:"foreignKey:UserId;constraint:OnUpdate:No Action;OnDelete:No Action"`
	Role   Role `gorm:"foreignKey:RoleId;constraint:OnUpdate:No Action;OnDelete:No Action"`
	UserId int
	RoleId int
}
