package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var (
	tData []teacher
	sData []student
	dData director
)

type (
	school interface {
		Add()
		Remove(int)
		FindByID(int)
	}
	person struct {
		ID      int    `json:"ID"`
		Name    string `json:"Name"`
		Surname string `json:"Surname"`
	}
	teacher struct {
		Person    person `json:"Person"`
		Subject   string `json:"Subject"`
		Classroom int    `json:"Classroom"`
	}
	student struct {
		Person person `json:"Person"`
		Class  int    `json:"Class"`
	}
	director struct {
		Person person `json:"Person"`
		Room   int    `json:"Room"`
		Phone  string `json:"Phone"`
	}
)

func (t teacher) Add() {
	file, _ := os.Open("../Files/TeacherData.json")
	byteSlice, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteSlice, &tData)
}

func (s student) Add() {
	file, _ := os.Open("../Files/StudentData.json")
	byteSlice, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteSlice, &sData)
}

func (d director) Add() {
	file, _ := os.Open("../Files/DirectorData.json")
	byteSlice, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteSlice, &dData)
}

func (teacher) Remove(IDDel int) {
	for n, c := range tData {
		if c.Person.ID == IDDel {
			copy(tData[n:], tData[n+1:])
			tData = tData[:len(tData)-1]
			return
		}
	}
}

func (student) Remove(IDDel int) {
	for n, c := range sData {
		if c.Person.ID == IDDel {
			copy(sData[n:], sData[n+1:])
			sData = sData[:len(sData)-1]
			return
		}
	}
}

func (director) Remove(IDDel int) {
	dData = director{Person: person{ID: 0, Name: "", Surname: ""}, Room: 0, Phone: ""}
}

func listEveryone() {
	listTeachers()
	listStudents()
	fmt.Println("Director:")
	fmt.Printf("ID: %d\nName: %s\nSurname: %s\nRoom: %d\nPhone: %s\n\n", dData.Person.ID, dData.Person.Name, dData.Person.Surname, dData.Room, dData.Phone)
}

func listStudents() {
	fmt.Println("students:")
	for _, s := range sData {
		fmt.Printf("ID: %d\nName: %s\nSurname: %s\nClass: %d\n\n", s.Person.ID, s.Person.Name, s.Person.Surname, s.Class)
	}
}

func listTeachers() {
	fmt.Println("teachers:")
	for _, t := range tData {
		fmt.Printf("ID: %d\nName: %s\nSurname: %s\nSubject: %s\nClassroom: %d\n\n", t.Person.ID, t.Person.Name, t.Person.Surname, t.Subject, t.Classroom)
	}
}

func (teacher) FindByID(IDFind int) {
	for _, t := range tData {
		if t.Person.ID == IDFind {
			fmt.Printf("ID: %d\nName: %s\nSurname: %s\nSubject: %s\nClassroom: %d\n\n", t.Person.ID, t.Person.Name, t.Person.Surname, t.Subject, t.Classroom)
			return
		}
	}
	fmt.Println("Teacher not found")
}

func (student) FindByID(IDFind int) {
	for _, s := range sData {
		if s.Person.ID == IDFind {
			fmt.Printf("ID: %d\nName: %s\nSurname: %s\nClass: %d\n\n", s.Person.ID, s.Person.Name, s.Person.Surname, s.Class)
			return
		}
	}
	fmt.Println("Student not found")
}

func (director) FindByID(IDFind int) {
	if IDFind == dData.Person.ID {
		fmt.Printf("ID: %d\nName: %s\nSurname: %s\nRoom: %d\nPhone: %s\n\n", dData.Person.ID, dData.Person.Name, dData.Person.Surname, dData.Room, dData.Phone)
	} else {
		fmt.Println("Director not found")
	}
}

func remove(t school, s school, d school) {
	idRem := 0
	fmt.Print("Enter id you want to delete: ")
	fmt.Scan(&idRem)
	if idRem <= 30 {
		t.Remove(idRem)
	} else if idRem <= 60 {
		s.Remove(idRem)
	} else {
		d.Remove(idRem)
	}
}

func find(t school, s school, d school) {
	idFind := 0
	fmt.Print("Enter id you want to find: ")
	fmt.Scan(&idFind)
	if idFind <= 30 {
		t.FindByID(idFind)
	} else if idFind <= 60 {
		s.FindByID(idFind)
	} else {
		d.FindByID(idFind)
	}
}

func main() {
	var (
		t   teacher
		s   student
		d   director
		num int
	)

	t.Add()
	s.Add()
	d.Add()
	for {
		fmt.Println("0) Exit\n1) List everyone\n2) Remove someone (1-30 teachers, 31-60 students, 61 - director)\n3) List everyone\n4) List students\n5) List teachers\n6) List director\n7) Show person by ID\n ")
		fmt.Print("Enter number of act you want to perform: ")
		fmt.Scan(&num)
		switch num {
		case 0:
			return
		case 1:
			listEveryone()
		case 2:
			remove(t, s, d)
		case 3:
			listEveryone()
		case 4:
			listStudents()
		case 5:
			listTeachers()
		case 6:
			fmt.Println("Director:")
			fmt.Printf("ID: %d\nName: %s\nSurname: %s\nRoom: %d\nPhone: %s\n\n", dData.Person.ID, dData.Person.Name, dData.Person.Surname, dData.Room, dData.Phone)
		case 7:
			find(t, s, d)
		}
	}
}
