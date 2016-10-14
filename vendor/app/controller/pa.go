package controller

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func LoginGET(w http.ResponseWriter, r *http.Request) {

	params := context.Get(r, "params").(httprouter.Params)
	fmt.Print(params.ByName("name"))
}

func LoginPOST(w http.ResponseWriter, r *http.Request) {

}
