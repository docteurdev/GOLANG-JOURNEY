package main

import "fmt"

func main() {
	newSalary := Humain{
		id:   "12324",
		name: "Josef",
		age:  32,
		Job: Job{
			name:     "backend devloper",
			salary:   1000000,
			bonus:    50000,
			position: "head of engeneering",
		},
	}

	fmt.Println(newSalary.position)

}

type Job struct {
	name     string
	position string
	salary   int
	bonus    int
}

type Humain struct {
	id   string
	name string
	age  int
	Job
}
