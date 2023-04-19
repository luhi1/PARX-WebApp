package main

import "net/http"

type HomeInfo struct {
}

func (e *HomeInfo) GETHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "home.gohtml", HomeInfo{})
	//@TODO: REMOVE
	if err != nil {
		return
	}
}
