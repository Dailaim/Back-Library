package models
type Book struct {
	BasicModel
	Title         string      `json:"title,omitempty" gorm:"not null"`
	Image         string      `json:"image,omitempty" gorm:"not null"`
	Resume        string      `json:"resume,omitempty"`
	AuthorsIDs    []uint      `gorm:"-" json:"authors_id,omitempty"`
	CategoriesIDs []uint      `gorm:"-" json:"categories_id,omitempty"`
	Authors       []*Author   `gorm:"many2many:book_authors" json:"authors,omitempty"`
	Categories    []*Category `gorm:"many2many:book_categories" json:"categories,omitempty"`
	Reviews       []Review    `gorm:"foreignKey:BookID" json:"reviews,omitempty"`
}
