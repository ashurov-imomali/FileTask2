package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	WriteToFile()
}

type Student struct {
	Name   string
	Age    int
	Course int
}

func NewStudent(name string, age int, course int) Student {
	return Student{
		Name:   name,
		Age:    age,
		Course: course,
	}
}

func WriteToFile() {
	file, err := os.OpenFile("./data.txt", os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
	FreshMan := NewStudent("Kurush", 20, 3)
	bytes, err1 := json.Marshal(FreshMan)
	if err1 != nil {
		log.Println(err1)
		return
	}
	_, err = file.Write(bytes)
	if err != nil {
		log.Println(err)
		return
	}
}
