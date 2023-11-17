// store.go
package InMemoryCRUDAPI

import (
	"errors"
	"sync"
)

type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

var (
	errNotFound = errors.New("item not found")
)

type Store struct {
	mu    sync.RWMutex
	items map[string]Item
}

func NewStore() *Store {
	return &Store{
		items: make(map[string]Item),
	}
}

func (s *Store) Create(item Item) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.items[item.ID] = item
}

func (s *Store) Get(id string) (Item, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	item, ok := s.items[id]
	if !ok {
		return Item{}, errNotFound
	}
	return item, nil
}

func (s *Store) Update(id string, newItem Item) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.items[id]
	if !ok {
		return errNotFound
	}
	s.items[id] = newItem
	return nil
}

func (s *Store) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.items[id]
	if !ok {
		return errNotFound
	}
	delete(s.items, id)
	return nil
}
