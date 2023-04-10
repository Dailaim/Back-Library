package models

type Review struct {
	BasicModel
	Comment string `json:"comment"`
	BookID  uint   `json:"book_id"`
	Book    Book   `gorm:"foreignKey:BookID" json:"book"`
	Score   uint   `json:"score"`
	UserID  uint   `gorm:"user_id"`
	User    User   `gorm:"foreignKey:UserID" json:"user"`
}
