package utils

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var InvalidURI = fmt.Errorf("Invalid URI")

type RequestUtil struct {
	l *log.Logger
}

func NewRequestUtil(l *log.Logger) *RequestUtil {
	return &RequestUtil{l: l}
}

func (util *RequestUtil) GetId(r *http.Request) (int, error) {
	reg := regexp.MustCompile(`/([0-9]+)`)
	g := reg.FindAllStringSubmatch(r.URL.Path, -1)
	if len(g) != 1 {
		util.l.Println("Invalid URI more than one id")
		return -1, InvalidURI
	}

	if len(g[0]) != 2 {
		util.l.Println("Invalid URI more than one capture group")
		return -1, InvalidURI
	}

	idString := g[0][1]
	id, err := strconv.Atoi(idString)
	if err != nil {
		util.l.Println("Invalid URI unable to convert to numer", idString)
		return -1, InvalidURI
	}
	return id, nil
}
