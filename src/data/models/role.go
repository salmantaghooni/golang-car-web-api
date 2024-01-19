package models

type Role struct {
	BaseModel
	Name      string `gorm:"size:30;type:string;not null;unique;"`
	UserRoles *[]UserRole
}
