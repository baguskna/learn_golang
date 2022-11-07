package main

import (
	"fmt"
	"os"
	"strings"
)

type biodata struct {
	ID      int
	name    string
	address string
	title   string
	reason  string
}

func main() {
	names := []string{"bagus", "kurnia", "ahmad", "jokowi", "agus"}
	biodatas := []biodata{
		{ID: 1, name: "Bagus", address: "Bengkulu", title: "Farmer", reason: "live suck"},
		{ID: 2, name: "Kurnia", address: "Tangerang", title: "Teacher", reason: "live suck"},
		{ID: 3, name: "Ahmad", address: "Jakarta", title: "Student", reason: "live suck"},
		{ID: 4, name: "Jokowi", address: "Solo", title: "President", reason: "----"},
		{ID: 5, name: "Agus", address: "Bekasi", title: "CEO", reason: "live suck"},
	}

	getArgsCli := os.Args
	secondArgs := getArgsCli[1]

	var whichIndex int

	for i, name := range names {
		if name == strings.ToLower(secondArgs) {
			whichIndex = i
		}
	}

	for i, data := range biodatas {
		if i == whichIndex {
			fmt.Println("ID : ", data.ID)
			fmt.Println("Name : ", data.name)
			fmt.Println("Address : ", data.address)
			fmt.Println("Title : ", data.title)
			fmt.Println("Reason : ", data.reason)
		}
	}
}
