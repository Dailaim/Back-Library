package models

type User struct {
	BasicModel// ID, CreatedAt, UpdatedAt, DeletedAt
	Email       string    `gorm:"not null;unique_index" validate:"required"`
	Name        string    `gorm:"not null" validate:"required,email,min=6"`
	Password    string    `gorm:"not null" validate:"required"`
	Photo       string    `gorm:"not null" validate:"required"`
	Reviews     []*Review `gorm:"foreignKey:UserID"`
}
