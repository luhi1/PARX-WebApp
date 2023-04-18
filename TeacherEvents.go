package main

import (
	"net/http"
)

// EventInfo @TODO: Figure out how to pass a file -> struct -> SQL
type StudentAttendance struct {
	StudentName   string
	StudentNumber int
}

type EventInfo struct {
	EventName           string
	Points              int
	EventDescription    string
	EventDate           string
	RoomNumber          int
	AdvisorNames        string
	Location            string
	LocationDescription string
	Sport               string
	SportDescription    string
	Attendance          []StudentAttendance
}

func (e *EventInfo) GETHandler(writer http.ResponseWriter, request *http.Request) {
	events := []EventInfo{}
	//SEMI-SCUFFED WAY OF MAKING THE USER NOT BE ABLE TO ACCESS HOME IF NOT LOGGED IN, CONSIDER USING COOKIES

	//Here we should populate the rest of the userInfo struct with sql queries and load whatever else we need for the home page.
	//Also, we need to find out how to get signup to upload to db and login to get
	//We can probably just do different interactions for get/post requests to the home, same way we did
	for i := 0; i < 3; i++ {
		events = append(events, EventInfo{
			EventName:           "asdfasd",
			Points:              0,
			EventDescription:    "asdf",
			EventDate:           "2017-06-01",
			RoomNumber:          0,
			AdvisorNames:        "asdf",
			Location:            "asdf",
			LocationDescription: "asdf",
			Sport:               "asdf",
			SportDescription:    "asdf",
			Attendance: []StudentAttendance{
				{"asdfasdf",
					1010},
			},
		})
	}
	err := tplExec(writer, "teacher_events.gohtml", events)
	//@TODO: REMOVE
	if err != nil {
		return
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
