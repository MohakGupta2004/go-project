package main

import (
	"fmt"
	"os"
	"time"
)

type Task struct{
  ID int
  Task string
  Created time.Time
  Done bool
}
var Tasks = []Task{}


func add(task string){
  newTask:= Task{
    ID: len(Tasks)+1,
    Task: task,
    Created: time.Now(),
    Done: false,
  } 
  Tasks = append(Tasks, newTask)
  for _, t := range Tasks{
  fmt.Printf(`
    %-4s %-20s %-30s %v
    %-4d %-20s %-30v %v
    `,
		"ID", "Task", "Created", "Done", // Header
		t.ID, t.Task, t.Created.Hour(), t.Done)

  }
}
func main(){
  // Edge case of args
  if(len(os.Args)<3){
    fmt.Println("Not enough args")
    os.Exit(-1)
  }
  cmd:=os.Args[1]
  task:=os.Args[2]

  switch cmd{
    case "add":
      add(task)
  }
}
