package controller

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func LoginGET(w http.ResponseWriter, r *http.Request) {
	v := view.New("login", r)
	v.Repopulate()
	v.Render(w)
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {

}
