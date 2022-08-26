package main

var Router = mux.New()

func init() {
	// add handler that listens for and processes
	Session.AddHandler(Router.OnMessageCreate)
}
