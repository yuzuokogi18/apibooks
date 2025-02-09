package entities

import (
	"golang.org/x/crypto/bcrypt"
)

type Client struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Password      string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"` 
	
}

func NewClient(name string, email string, phone string, password string) (*Client, error) {
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Client{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: string(hashedPassword),
		
	}, nil
}

func (b *Client) SetID(id int) {
	b.ID = id
}

func (b *Client) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(b.Password), []byte(password))
	return err == nil
}
