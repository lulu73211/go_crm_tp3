package json

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	dom "github.com/lulu73211/go_crm_tp3/internal/domain"
	"github.com/lulu73211/go_crm_tp3/internal/store"
)

var _ store.Storer = (*JSONStore)(nil)

type JSONStore struct {
	path   string
	nextID uint
	items  []dom.Contact
}

func New(path string) *JSONStore { return &JSONStore{path: path, nextID: 1} }

func (s *JSONStore) Init() error {
	if s.path == "" {
		return fmt.Errorf("json path is empty")
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	b, err := os.ReadFile(s.path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s.flush()
		}
		return err
	}
	if len(b) == 0 {
		return s.flush()
	}
	var box struct {
		NextID uint          `json:"next_id"`
		Items  []dom.Contact `json:"items"`
	}
	if err := json.Unmarshal(b, &box); err != nil {
		return err
	}
	s.nextID = box.NextID
	s.items = box.Items
	if s.nextID == 0 {
		var max uint
		for _, it := range s.items {
			if it.ID > max {
				max = it.ID
			}
		}
		s.nextID = max + 1
	}
	return nil
}

func (s *JSONStore) flush() error {
	box := struct {
		NextID uint          `json:"next_id"`
		Items  []dom.Contact `json:"items"`
	}{NextID: s.nextID, Items: s.items}
	b, err := json.MarshalIndent(box, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.path, b, 0o644)
}

func (s *JSONStore) CreateContact(c *dom.Contact) error {
	c.ID = s.nextID
	s.nextID++
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	s.items = append(s.items, *c)
	return s.flush()
}

func (s *JSONStore) ListContacts() ([]dom.Contact, error) {
	return append([]dom.Contact(nil), s.items...), nil
}

func (s *JSONStore) UpdateContact(c dom.Contact) error {
	for i := range s.items {
		if s.items[i].ID == c.ID {
			if c.Name != "" {
				s.items[i].Name = c.Name
			}
			if c.Email != "" {
				s.items[i].Email = c.Email
			}
			if c.Phone != "" {
				s.items[i].Phone = c.Phone
			}
			s.items[i].UpdatedAt = time.Now()
			return s.flush()
		}
	}
	return errors.New("contact not found")
}

func (s *JSONStore) DeleteContact(id uint) error {
	for i := range s.items {
		if s.items[i].ID == id {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return s.flush()
		}
	}
	return errors.New("contact not found")
}
