package models

type User struct {
	BasicModel
	Email    string    `json:"email,omitempty" gorm:"not null;unique_index" validate:"required"`
	Name     string    `json:"name,omitempty" gorm:"not null" validate:"required,email,min=6"`
	Password string    `json:"password,omitempty" gorm:"not null" validate:"required"`
	Photo    string    `json:"photo,omitempty"  gorm:"not null" validate:"required"`
	Reviews  []*Review `json:"reviews,omitempty"  gorm:"foreignKey:UserID"`
}
