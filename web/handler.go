package web

import (
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"wairimuian.com/GoReddit"
)

func NewHandler(store GoReddit.Store) *Handler {
	return &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}
}

type Handler struct {
	*chi.Mux
	store GoReddit.Store
}

func (h *Handler) ThreadsList() http.HandlerFunc {
	type data struct {
		Threads GoReddit.Thread
	}
	tmpl := template.Must(template.New("").Parse(``))
	return func(w http.ResponseWriter, r *http.Request) {
		tt, err := h.store.Threads()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := tmpl.Execute(w, data{Threads: tt}); err != nil {
			return
		}
	}
}
