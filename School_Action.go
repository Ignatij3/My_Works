package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"
)

type (
	school interface {
		create(actionData)
		update(actionData)
		read(actionData)
		delete(actionData)
	}
	actionData struct {
		action string `json:"action"`
		pData data `json:"data"`
	}
	bsData struct {
		name string `json:"name"`
		surname string `json:"surname"`
		personalCode string `json:"personalCode"`
	}
	data struct {
		id int `json:"id"`
		subject string `json:"subject"`
		salary float32 `json:"salary"`
		classrooms []string `json:"classrooms"`
		person bsData `json:"person"`
	}
	tData []data
)

func (d *tData) create(a actionData) {
	newID := 0
	for range *d {newID++}
	*d = append(*d, a.pData)
	(*d)[len(*d) - 1].id = newID
}

func (d *tData) update(a actionData) {
	for n, c := range *d {
		if c.id == a.pData.id {
			(*d)[n] = a.pData
			return
		}
	}
}

func (d *tData) read(a actionData) {
	for _, c := range *d {
		fmt.Printf("id - %d\nname, surname - %s, %s (%s)\nsubject - %s\nsalary - %f\nclassrooms - %v\n\n", c.id, c.person.name, c.person.surname, c.subject, c.salary, c.classrooms)
	}
}

func (d *tData) delete(a actionData) {
	for n, c := range *d {
		if c.id == a.pData.id {
			copy((*d)[n:], (*d)[n+1:])
			*d = (*d)[:len(*d)-1]
			return
		}
	}
}

func commitActions(actions []actionData) tData {
	var teachers tData
	for _, c := range actions {
		switch c.action {
			case "create":
				teachers.create(c)
			case "update":
				teachers.update(c)
			case "read":
				teachers.read(c)
			case "delete":
				teachers.delete(c)
		}
	}
	return teachers
}

func main() {
	var aData []actionData
	
	file, _ := os.Open("../Files/ActionData.json")
	byteSlice, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteSlice, &aData)
	
	fmt.Println(aData)
	teachers := commitActions(aData)
	fmt.Println(teachers)
}
