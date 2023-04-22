package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// EventInfo @TODO: Figure out how to pass a file -> struct -> SQL
type Winners struct {
	RandomNinthWinner    string
	RandomTenthWinner    string
	RandomEleventhWinner string
	RandomTwelvthWinner  string
	NinthWinners         []string
	TenthWinners         []string
	EleventhWinners      []string
	TwelvthWinners       []string
}

type Winner struct {
	StudentName string
	Points      int
	GradeLevel  int
}

func (w *Winners) GETHandler(writer http.ResponseWriter, request *http.Request) {
	*w = Winners{}
	rows, err := db.Query("select StudentName, Points, GradeLevel from users left join grades on users.GradeID = grades.ID order by GradeLevel, Points desc;")
	if err != nil {
		fmt.Println(err)
		return
	}
	for rows.Next() {
		currentWinner := Winner{}
		rows.Scan(&currentWinner.StudentName, &currentWinner.Points, &currentWinner.GradeLevel)
		switch currentWinner.GradeLevel {
		case 9:
			w.NinthWinners = append(w.NinthWinners, currentWinner.StudentName+"; Points: "+strconv.Itoa(currentWinner.Points))
		case 10:
			w.TenthWinners = append(w.TenthWinners, currentWinner.StudentName+"; Points: "+strconv.Itoa(currentWinner.Points))
		case 11:
			w.EleventhWinners = append(w.EleventhWinners, currentWinner.StudentName+"; Points: "+strconv.Itoa(currentWinner.Points))
		case 12:
			w.TwelvthWinners = append(w.TwelvthWinners, currentWinner.StudentName+"; Points: "+strconv.Itoa(currentWinner.Points))
		}
	}
	//select RandomWinner from grades left join users on users.UserID = grades.RandomWinner;
	randWinners, err := db.Query("select RandomWinner, grades.GradeLevel from grades left join users on users.UserID = grades.RandomWinner")
	if err != nil {
		fmt.Println(err)
		return
	}
	for randWinners.Next() {
		currentWinner := Winner{}
		randWinners.Scan(&currentWinner.StudentName, &currentWinner.GradeLevel)
		switch currentWinner.GradeLevel {
		case 9:
			w.RandomNinthWinner = currentWinner.StudentName
		case 10:
			w.RandomTenthWinner = currentWinner.StudentName
		case 11:
			w.RandomEleventhWinner = currentWinner.StudentName
		case 12:
			w.RandomTwelvthWinner = currentWinner.StudentName
		}
	}
	tplExec(writer, "winners.gohtml", *w)
}

func (w *Winners) POSTHandler(writer http.ResponseWriter, request *http.Request) {
	//@todo: Implement.
}

func (w *Winners) valHandler(writer http.ResponseWriter, request *http.Request) {
	//@todo: Implement Data Validation.
}

func (w *Winners) dataVal(requestMethod string) bool {
	//@todo: Implement Data Validation.
	return false
}
