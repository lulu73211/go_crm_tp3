package memory

import (
	"errors"
	"sync"

	"github.com/lulu73211/go_crm_tp3/internal/domain"
	"github.com/lulu73211/go_crm_tp3/internal/store"
)

var _ store.Storer = (*MemoryStore)(nil)

type MemoryStore struct {
	mu     sync.Mutex
	nextID uint
	items  map[uint]domain.Contact
}

func New() *MemoryStore {
	return &MemoryStore{
		nextID: 1,
		items:  make(map[uint]domain.Contact),
	}
}

func (m *MemoryStore) Init() error { return nil }

func (m *MemoryStore) CreateContact(c *domain.Contact) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	c.ID = m.nextID
	m.nextID++
	m.items[c.ID] = *c
	return nil
}

func (m *MemoryStore) ListContacts() ([]domain.Contact, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]domain.Contact, 0, len(m.items))
	for _, v := range m.items {
		out = append(out, v)
	}
	return out, nil
}

func (m *MemoryStore) UpdateContact(c domain.Contact) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	old, ok := m.items[c.ID]
	if !ok {
		return errors.New("contact not found")
	}
	if c.Name != "" {
		old.Name = c.Name
	}
	if c.Email != "" {
		old.Email = c.Email
	}
	if c.Phone != "" {
		old.Phone = c.Phone
	}
	m.items[c.ID] = old
	return nil
}

func (m *MemoryStore) DeleteContact(id uint) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.items[id]; !ok {
		return errors.New("contact not found")
	}
	delete(m.items, id)
	return nil
}
