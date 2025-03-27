package testpkg

import (
	"context"
	"fmt"
)

type MyStruct1 struct {
	ID int
	Context context.Context 
}

type MyStruct2 struct {
	Name string
	Context context.Context
}

func (m1 MyStruct1) PrintInfo() {
	fmt.Println("ID:", m1.ID)
}

func (m2 MyStruct2) PrintlnInfo() {
	fmt.Println("Name:", m2.Name)
}