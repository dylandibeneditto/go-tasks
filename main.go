// Golang program to illustrate the
// concept of parsing JSON to an array
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

type Task struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Importance int    `json:"importance"`
}

func main() {

	//defining the array of tasks
	var tasks Tasks

	jsonData, jsonErr := os.Open("/Users/dylandibeneditto/Desktop/new/go-todo/items.json")

	if jsonErr != nil {
		fmt.Println(jsonErr)
	}

	byteValue, _ := io.ReadAll(jsonData)

	//decoding JSON array to tasks array
	err := json.Unmarshal([]byte(byteValue), &tasks)

	if err != nil {
		fmt.Println("Error while decoding the data", err.Error())
	}

	//printing decoded array values one by one
	for i := 0; i < len(tasks.Tasks); i++ {

		fmt.Println("Task Title: " + tasks.Tasks[i].Desc)
	}

	defer jsonData.Close()

}
