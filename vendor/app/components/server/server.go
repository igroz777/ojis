package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	Hostname    string `json:Hostname`
	HTTPSPort   int    `json:HTTPSPort`
	TokenSecret string `TokenSecret`
}

func Start(httpHandlers http.Handler, s Server) {

	fmt.Println(time.Now().Format("2006-01-02 03:04:05 PM"), "Running HTTPS "+httpsAddress(s))

	//Start http server with httpHandlers being a httprouter.Router
	log.Fatal(http.ListenAndServe(httpsAddress(s), httpHandlers))

}

func httpsAddress(s Server) string {
	return s.Hostname + ":" + fmt.Sprintf("%d", s.HTTPSPort)
}
