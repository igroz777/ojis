package route

import (
	//"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"

	"app/controller"
	"app/route/middleware/auth"
	wrapper "app/route/middleware/httprouterwrapper"
)

func GetRouter() *httprouter.Router {
	r := httprouter.New()

	r.GET("/", wrapper.Handler(alice.
		New().
		ThenFunc(controller.IndexGET)))

	r.GET("/login/:name/:age", wrapper.Handler(alice.
		New().
		ThenFunc(controller.LoginGET)))

	r.POST("/login", wrapper.Handler(alice.
		New(auth.DisallowAuth).
		ThenFunc(controller.LoginPOST)))

	r.POST("/logout", wrapper.Handler(alice.
		New(auth.DissalowAnon).
		ThenFunc(controller.LogoutPOST)))

	//r.GET("/users/:username", hr.Handler(alice.
	//New().
	//ThenFunc(controller.UserProfileGET)))

	return r
}
