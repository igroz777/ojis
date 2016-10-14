package view

import (
	"net/http"
)

type View struct {
	Name    string
	Vars    map[string]string
	request http.Request
}

func New(name string, r http.Request) *View {
	v := &View{Name: name, Vars: {}, request: r}
}

func (v *View) Render(w http.ResponseWriter) error {

}

func (v *View) Repopulate() error {

}
