package main

import (
	"net/http"
	"strconv"
	"strings"
)

type UserData struct {
	Name         string
	Grade        int
	IdNumber     int
	passwordHash string
	valid        DisplayError
}

func (u *UserData) GETHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "login.gohtml", u.valid)
	//@TODO: REMOVE
	if err != nil {
		return
	}
	u.valid = DisplayError{""}
}

func (u *UserData) POSTHandler(writer http.ResponseWriter, request *http.Request) {
	err := tplExec(writer, "signup.gohtml", u.valid)
	//@TODO: REMOVE
	if err != nil {
		return
	}
	u.valid = DisplayError{""}
}

func (u *UserData) valHandler(writer http.ResponseWriter, request *http.Request) {
	var err error
	err = request.ParseForm()
	if err != nil {
		http.Redirect(writer, request, "./error", 303)
		return
	}
	u.Name = request.FormValue("name")
	u.Grade, err = strconv.Atoi(request.FormValue("grade"))
	u.IdNumber, err = strconv.Atoi(request.FormValue("IdNumber"))
	u.passwordHash = hashPswd(request.FormValue("password"))

	if err != nil || u.dataVal(strings.TrimPrefix(request.URL.Path, "/userValidation/")) {
		http.Redirect(writer, request, "../teacher_events", 307)
	} else {
		u.valid = DisplayError{"Invalid Credentials"}
		if strings.TrimPrefix(request.URL.Path, "/userValidation/") == "signup" {
			http.Redirect(writer, request, "../signup", 303)
		} else {
			http.Redirect(writer, request, "../login", 303)
		}
	}
}

func (u *UserData) dataVal(requestMethod string) bool {
	valid := false
	if (*u != UserData{}) &&
		(u.IdNumber > 0 &&
			u.IdNumber < 9999999 &&
			u.passwordHash != hashPswd("")) {

		valid = true
	}

	if requestMethod == "signup" && ((u.Grade != 9 && u.Grade != 10 &&
		u.Grade != 11 && u.Grade != 12) || u.Name == "") {
		valid = false
	}
	//here we can check if the password matches the one in our database
	if !valid {
		*u = UserData{}
	}
	return valid
}
