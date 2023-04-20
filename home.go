package main

import (
	"html/template"
	"net/http"
)

type StudentPageHandlers interface {
	GETStudentHandler(writer http.ResponseWriter, request *http.Request)
}

type HomeInfo struct {
}

func (e *HomeInfo) GETStudentHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "home.gohtml", HomeInfo{})
	//@TODO: REMOVE
	if err != nil {
		return
	}
}

func tplExec2(w http.ResponseWriter, filename string, information any, filename2 string) error {
	temp := template.Must(template.ParseFiles(filename, filename2))

	err := temp.Execute(w, information)
	//@TODO: REMOVE
	if err != nil {
		return err
	}
	return nil
}
