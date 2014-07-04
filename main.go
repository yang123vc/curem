package main

import (
	"github.com/codegangsta/negroni"
	"github.com/joshsoftware/curem/config"
)

func main() {

	c := make(map[string]string)
	c["name"] = "test"
	c["url"] = "localhost"
	c["leads"] = "newlead"
	c["contacts"] = "newcontact"

	config.Configure(c)

	// r is a *mux.Router defined in contact_api.go

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":3000")
}
