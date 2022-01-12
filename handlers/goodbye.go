package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger
}

func NewGoodBye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(rw, "Erro", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", d)
}
