package service

import (
	"fmt"

	"github.com/lulu73211/go_crm_tp3/internal/domain"
	"github.com/lulu73211/go_crm_tp3/internal/store"
)

type CRM struct{ S store.Storer }

func New(st store.Storer) *CRM { return &CRM{S: st} }

func (c *CRM) Add(name, email, phone string) (*domain.Contact, error) {
	co := &domain.Contact{Name: name, Email: email, Phone: phone}
	if err := c.S.CreateContact(co); err != nil {
		return nil, err
	}
	return co, nil
}

func (c *CRM) List() ([]domain.Contact, error) {
	return c.S.ListContacts()
}

func (c *CRM) Update(id uint, name, email, phone string) error {
	if id == 0 {
		return fmt.Errorf("id is required")
	}
	return c.S.UpdateContact(domain.Contact{ID: id, Name: name, Email: email, Phone: phone})
}

func (c *CRM) Delete(id uint) error {
	if id == 0 {
		return fmt.Errorf("id is required")
	}
	return c.S.DeleteContact(id)
}
