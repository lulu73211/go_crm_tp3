package gormstore

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	dom "github.com/lulu73211/go_crm_tp3/internal/domain"
	"github.com/lulu73211/go_crm_tp3/internal/store"
)

var _ store.Storer = (*GORMStore)(nil)

type GORMStore struct {
	path string
	db   *gorm.DB
}

func New(path string) *GORMStore { return &GORMStore{path: path} }

func (s *GORMStore) Init() error {
	if s.path == "" {
		return fmt.Errorf("db path is empty")
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return err
	}
	db, err := gorm.Open(sqlite.Open(s.path), &gorm.Config{})
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&dom.Contact{}); err != nil {
		return err
	}
	s.db = db
	return nil
}

func (s *GORMStore) CreateContact(c *dom.Contact) error {
	c.CreatedAt = time.Now()
	c.UpdatedAt = c.CreatedAt
	return s.db.Create(c).Error
}

func (s *GORMStore) ListContacts() ([]dom.Contact, error) {
	var out []dom.Contact
	return out, s.db.Order("id ASC").Find(&out).Error
}

func (s *GORMStore) UpdateContact(c dom.Contact) error {
	var existing dom.Contact
	if err := s.db.First(&existing, c.ID).Error; err != nil {
		return err
	}
	if c.Name != "" {
		existing.Name = c.Name
	}
	if c.Email != "" {
		existing.Email = c.Email
	}
	if c.Phone != "" {
		existing.Phone = c.Phone
	}
	existing.UpdatedAt = time.Now()
	return s.db.Save(&existing).Error
}

func (s *GORMStore) DeleteContact(id uint) error {
	res := s.db.Delete(&dom.Contact{}, id)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("contact not found")
	}
	return nil
}
