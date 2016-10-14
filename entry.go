package main

import (
	"encoding/json"
	"os"

	"app/components/jsonconfig"
	"app/components/server"
	"app/route"
)

type configuration struct {
	Server server.Server `json:Server`
}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

var config = &configuration{}

func main() {

	jsonconfig.Load("config"+string(os.PathSeparator)+"config.json", config)
	server.Start(route.GetRouter(), config.Server)
}
