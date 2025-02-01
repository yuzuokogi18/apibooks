package entities

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float32
}

func NewBook(title string, author string, price float32) *Book {
	return &Book{
		Title:  title,
		Author: author,
		Price:  price,
	}
}
func (b *Book) SetID(id int) {
	b.ID = id
}
