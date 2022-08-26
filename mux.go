package main

import (
	"technoabsurdist/digest/x/mux"
)

// glob to allow easy access throughout bot
var Router = mux.New()

func init() {
	// add handler that listens for and processes
	Session.AddHandler(Router.OnMessageCreate)
}
