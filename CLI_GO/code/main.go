package main

import (
	"fmt"
	"time"
)

type item struct {
	Task       string    //description of task
	Status     bool      //Status of task ie done or not
	CreatedAt  time.Time //time of creation
	FinishedAt time.Time //time of completion
}

//now to store all the tasks we have to

func main() {
	fmt.Println("Hi")
}
