package controller

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func LogoutPOST(w http.ResponseWriter, r *http.Request) {

	params := context.Get(r, "params").(httprouter.Params)
	fmt.Print(params.ByName("name"))
}
