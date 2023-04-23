package main

import (
	"html/template"
	"net/http"
)

type StudentPageHandlers interface {
	GETStudentHandler(writer http.ResponseWriter, request *http.Request)
}
type PrizeData struct {
	Name      string
	Threshold int
}
type HomeData struct {
	Name   string
	Grade  int
	Points int
	//the following can be changed later, it was just convenient and I liked it as a solution, but it may not work
	Grade9Points         []int
	Grade10Points        []int
	Grade11Points        []int
	Grade12Points        []int
	NinthWinners         []string
	TenthWinners         []string
	EleventhWinners      []string
	TwelfthWinners       []string
	RandomNinthWinner    string
	RandomTenthWinner    string
	RandomEleventhWinner string
	RandomTwelfthWinner  string
	Prizes               []PrizeData
}

func (e *HomeData) GETStudentHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "home.gohtml", *e)
	//@TODO: REMOVE
	if err != nil {
		return
	}
}

func tplExec2(w http.ResponseWriter, filename string, information any, filename2 string) error {
	//tplExec may be different now because /webpages ???
	temp := template.Must(template.ParseFiles(filename, filename2))

	err := temp.Execute(w, information)
	//@TODO: REMOVE
	if err != nil {
		return err
	}
	return nil
}
