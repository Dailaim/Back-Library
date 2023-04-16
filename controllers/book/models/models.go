package models

type Book struct {
	Title         string      `json:"title,omitempty"`
	Image         string      `json:"image,omitempty"`
	Resume        string      `json:"resume,omitempty"`
	AuthorsIDs    []uint      `json:"authors_id,omitempty"`
	CategoriesIDs []uint      `json:"categories_id,omitempty"`
	Authors       []*Author   `json:"authors,omitempty"`
	Categories    []*Category `json:"categories,omitempty"`
	Reviews       []Review    `json:"reviews,omitempty"`
}


