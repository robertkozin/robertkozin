package website

import "github.com/go-chi/chi/v5"

func (s *Service) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/imbox", func(r chi.Router) {
		r.Post("/file", s.Imbox.File)
		r.Post("/link", s.Imbox.Link)
	})

	return r
}