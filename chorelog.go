package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Chore is the struct to hold each chore
type Chore struct {
	Title       string
	Description string
	Room        string
	Interval    int
	Priority    int
	Catagory    string
}

//ChoreList will hold the chores and other information
type ChoreList struct {
	Chores map[string]Chore
}

func main() {
	var masterList ChoreList
	masterList.Chores = map[string]Chore{}
	choreFile, err := os.Open("choreList.json")
	if err != nil {
		fmt.Println(err)
	}
	x := json.NewDecoder(choreFile)
	err = x.Decode(&masterList)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(masterList)
	fmt.Println(masterList.Chores[`dishes`])

}
