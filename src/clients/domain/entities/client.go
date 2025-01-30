package entities

type Client struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
}

func NewClient(name string, email string, phone string) *Client {
	return &Client{Name: name, Email: email, Phone: phone}
}
func (b *Client) SetID(id int) {
	b.ID = id
}
