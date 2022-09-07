package main

import (
	"fmt"
	"os"
	"strconv"
)

type student struct {
	id     uint8
	name   string
	job    string
	reason string
}

func main() {
	var studentId int

	studentId, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("Student ID is not valid")
		return
	}

	studentList := createStudentList()

	if studentId <= 0 {
		fmt.Printf("Student Id %d must be at least 1\n", studentId)
		return
	}

	studentIdx := findStudentIndex(uint8(studentId), studentList)

	if studentIdx < 0 {
		fmt.Printf("Student with ID %d is not found\n", studentId)
	}

	student := studentList[studentIdx]
	fmt.Println("Student Found")
	fmt.Println("Name :", student.name)
	fmt.Println("Job :", student.job)
	fmt.Println("Reason to choose golang :", student.reason)
}

func findStudentIndex(id uint8, studentList []student) int {
	for i, std := range studentList {
		if std.id == id {
			return i
		}
	}
	return -1
}

func createStudentList() []student {
	return []student{
		{
			id:     1,
			name:   "Wahyu Setiawan",
			job:    "Mobile Application Engineer/Student",
			reason: "Want to work as Back End Engineer, because Front End is boring",
		},

		{
			id:     2,
			name:   "William Shakespeare",
			job:    "Poem Writer",
			reason: "Want to make a poem for programming devs",
		},
		{
			id:     3,
			name:   "Beben",
			job:    "QA Engineer",
			reason: "Want to be a backend engineer",
		},
		{
			id:     4,
			name:   "Ezio Auditore",
			job:    "Assassins",
			reason: "Now regret being assassins",
		},
	}
}
