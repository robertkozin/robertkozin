package website

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"robertkozin.com/services/imbox"
)

//embed:

type Service struct {
	Imbox *imbox.Service
	Router chi.Router
}

func New() *Service {
	return &Service{
		Imbox: MustValue(imbox.New("./tmp/imbox")),
	}
}

func (s *Service) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.Router.ServeHTTP(res, req)
}

func MustValue[V any](value V, err error) V {
	if err != nil {
		panic(err)
	}
	return value
}
