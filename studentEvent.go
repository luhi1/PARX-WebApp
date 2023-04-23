package main

import (
	"net/http"
)

func (e *EventInfo) GETStudentHandler(writer http.ResponseWriter, request *http.Request) {
	events := []EventInfo{}
	//SEMI-SCUFFED WAY OF MAKING THE USER NOT BE ABLE TO ACCESS HOME IF NOT LOGGED IN, CONSIDER USING COOKIES

	//Here we should populate the rest of the userInfo struct with sql queries and load whatever else we need for the home page.
	//Also, we need to find out how to get signup to upload to db and login to get
	//We can probably just do different interactions for get/post requests to the home, same way we did
	/*for i := 0; i < 2; i++ {
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
	}*/
	err := tplExec2(writer, "studentEvents.gohtml", events, "home.gohtml")
	//@TODO: REMOVE
	if err != nil {
		return
	}
}
