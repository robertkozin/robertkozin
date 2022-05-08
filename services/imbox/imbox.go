package imbox

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

// Endpoints to recieve
//    Files, links, multipart
//    TODO: figure out how to rec raw files

// Homepage
// Returns just the top item
// you have to decide what you want to do with it
//

// hard problems to solve:
// where to store images
// Where to store a little bit of metadata
// where to store image metadata

type Service struct {
	Dir string
}

func New(imboxDir string) (service *Service, err error) {
	err = os.MkdirAll(imboxDir, os.ModePerm)
	service = &Service{Dir: imboxDir}
	return
}

func (s *Service) Imbox() {
	// get latest img
	// render page
}

func (s *Service) File(res http.ResponseWriter, req *http.Request) {
	ct := req.Header.Get("Content-Type")

	fmt.Println(req.Header)

	now := time.Now().Unix()
	fname := fmt.Sprintf("%v.jpeg", now)

	out, err := os.Create(path.Join(s.Dir, fname))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	defer req.Body.Close()
	io.Copy(out, req.Body)

	fmt.Println(ct)
}

func (s *Service) Link(res http.ResponseWriter, req *http.Request) {
	link := req.FormValue("link")

	// get link
}

func (s *Service) Multipart() {

}
