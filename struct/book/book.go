package book

type Book struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author"`
}
