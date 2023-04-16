package AuthorModels

type Author struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Age       int    `json:"age,omitempty"`
	Books     []Book `json:"books,omitempty"`
}

type NewAuthor struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Age       int    `json:"age,omitempty"`
}

type Book struct {
	Title  string `json:"title,omitempty"`
	Image  string `json:"image,omitempty"`
	Resume string `json:"resume,omitempty"`
}

type Response struct {
	Data  Data   `json:"data"`
	Error *Error `json:"error"`
}

type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type Data interface{}

type SingleAuthorResponse struct {
	Data  *Author `json:"data"`
	Error *Error  `json:"error,omitempty"`
}

type MultipleAuthorsResponse struct {
	Data  []Author `json:"data"`
	Error *Error   `json:"error,omitempty"`
}
