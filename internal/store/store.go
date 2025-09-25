package store

import "github.com/lulu73211/go_crm_tp3/internal/domain"

type Storer interface {
	Init() error
	CreateContact(c *domain.Contact) error
	ListContacts() ([]domain.Contact, error)
	UpdateContact(c domain.Contact) error
	DeleteContact(id uint) error
}
