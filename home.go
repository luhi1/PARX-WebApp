package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type StudentPageHandlers interface {
	GETStudentHandler(writer http.ResponseWriter, request *http.Request)
}
type PrizeData struct {
	Name      string
	Threshold int
}

type HomeData struct {
	U       *UserData
	Prizes  []Prize
	Winners Winners
}

func (h *HomeData) GETStudentHandler(writer http.ResponseWriter, request *http.Request) {
	h.Prizes = []Prize{}
	h.Winners = Winners{}
	insert, err := db.Query("select PrizeName, PointThreshold from prizes")
	if err != nil {
		fmt.Println(err)
		return
	}
	for insert.Next() {
		p := Prize{}
		insert.Scan(&p.PrizeName, &p.Points)
		h.Prizes = append(h.Prizes, p)
	}

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
			h.Winners.NinthWinners = append(h.Winners.NinthWinners, currentWinner.StudentName+"; Points: "+strconv.Itoa(currentWinner.Points))
		case 10:
			h.Winners.TenthWinners = append(h.Winners.TenthWinners, currentWinner.StudentName+"; Points: "+strconv.Itoa(currentWinner.Points))
		case 11:
			h.Winners.EleventhWinners = append(h.Winners.EleventhWinners, currentWinner.StudentName+"; Points: "+strconv.Itoa(currentWinner.Points))
		case 12:
			h.Winners.TwelvthWinners = append(h.Winners.TwelvthWinners, currentWinner.StudentName+"; Points: "+strconv.Itoa(currentWinner.Points))
		}
	}

	fmt.Println(h.U.Name)
	err = tplExec(writer, "home.gohtml", *h)
	//@TODO: REMOVE
	if err != nil {
		return
	}
}

func multiTplExec(w http.ResponseWriter, filename string, information any, filename2 string) error {
	//tplExec may be different now because /webpages ???
	temp := template.Must(template.ParseFiles("WebPages/"+filename, "WebPages/"+filename2))

	err := temp.Execute(w, information)
	//@TODO: REMOVE
	if err != nil {
		return err
	}
	return nil
}
