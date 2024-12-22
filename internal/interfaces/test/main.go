package main

import (
	"errors"
	"fmt"

	"github.com/r1005410078/meida-admin-server/internal/interfaces/shared"
)


func main() {
	// is is the main function
	eventBus := shared.NewEventBus()
	eventBus.Register(AddStudent)
	eventBus.Register(AddPersion)

	eventBus.Dispatch(Student{name: "Tom"})
	err := eventBus.Dispatch(Persion{name: "Rongts"})

	if err != nil {
		fmt.Println(err)
	}
}

type Student struct {
	name string
}

func AddStudent(s Student) error {
	// This is the Add function
	fmt.Printf("Add = %v ", s.name)
	return nil
}

type Persion struct  {
	name string
}

func AddPersion(s Persion) error {
	// This is the Add function
	fmt.Printf("Add = %v ", s.name)
	return errors.New("error in AddPersion") 
}