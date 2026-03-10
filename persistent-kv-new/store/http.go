package store

import (
	"encoding/json"
	"net/http"
)

func (s *Store) HandleSet(rw http.ResponseWriter, req *http.Request) {

	key := req.URL.Query().Get("key")
	val := req.URL.Query().Get("value")

	err := s.Set(key, val)
	if err != nil {
		http.Error(rw, err.Error(), 500)
		return
	}

	rw.Write([]byte("OK"))
}

func (s *Store) HandleGet(rw http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")

	val, ok := s.Get(key)
	if !ok {
		http.NotFound(rw, req)
		return
	}

	json.NewEncoder(rw).Encode(map[string]string{
		"value": val,
	})
}
