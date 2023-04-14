package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"strconv"
	"strings"
)

// EventInfo @TODO: Figure out how to pass a file -> struct -> SQL
type EventInfo struct {
	Points              int
	EventDescription    string
	EventDate           string
	RoomNumber          int
	AdvisorNames        string
	Location            string
	LocationDescription string
	Sport               string
	SportDescription    string
	EventImage          string
	StudentName         string
	StudentNumber       int
	StudentAttended     bool
	inputImage          fs.File
}

func (e *EventInfo) GETHandler(writer http.ResponseWriter, request *http.Request) {
	if (userInfo != UserData{}) {
		//SEMI-SCUFFED WAY OF MAKING THE USER NOT BE ABLE TO ACCESS HOME IF NOT LOGGED IN, CONSIDER USING COOKIES

		//Here we should populate the rest of the userInfo struct with sql queries and load whatever else we need for the home page.
		//Also, we need to find out how to get signup to upload to db and login to get
		//We can probably just do different interactions for get/post requests to the home, same way we did
		var Events []EventInfo
		for i := 0; i < 3; i++ {
			Events = append(Events, EventInfo{
				Points:              0,
				EventDescription:    "asdf",
				EventDate:           "2017-06-01",
				RoomNumber:          0,
				AdvisorNames:        "asdf",
				Location:            "asdf",
				LocationDescription: "asdf",
				Sport:               "asdf",
				SportDescription:    "asdf",
				EventImage:          "https://imgs.search.brave.com/ToRVheIVFOHdWRebW6v6BriMZf_slwrqoAXvU-I62CY/rs:fit:1200:1200:1/g:ce/aHR0cHM6Ly90aGV3/b3dzdHlsZS5jb20v/d3AtY29udGVudC91/cGxvYWRzLzIwMTUv/MDEvbmF0dXJlLWlt/YWdlcy4uanBn",
				StudentName:         "asdf",
				StudentNumber:       0,
				StudentAttended:     true,
			})
		}
		err := tplExec(writer, "teacher_events.gohtml", Events)
		//@TODO: REMOVE
		if err != nil {
			return
		}
	} else {
		http.Redirect(writer, request, "./login", 303)
	}
}

func (e *EventInfo) POSTHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "teacher_create_event.gohtml", nil)
	//@TODO: REMOVE
	if err != nil {
		return
	}
}

func (e *EventInfo) valHandler(writer http.ResponseWriter, request *http.Request) {
	//@todo: Implement Data Validation.
}

func (e *EventInfo) dataVal(requestMethod string) bool {
	//@todo: Implement Data Validation.
	return false
}

