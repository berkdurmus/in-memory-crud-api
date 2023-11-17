// handler.go
package InMemoryCRUDAPI

import (
	"encoding/json"
	"net/http"
)

func (s *Store) createItemHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.Create(item)
	w.WriteHeader(http.StatusCreated)
}

func (s *Store) getItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	item, err := s.Get(id)
	if err != nil {
		if err == errNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (s *Store) updateItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	var item Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.Update(id, item); err != nil {
		if err == errNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *Store) deleteItemHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if err := s.Delete(id); err != nil {
		if err == errNotFound {
			http.NotFound(w, r)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
