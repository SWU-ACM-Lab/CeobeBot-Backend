package main

import (
	"CeobeBot-Backend/controller/initializer"
	"fmt"
)

func main() {
	fmt.Println(initializer.Init())
	fmt.Println(initializer.SyncDatabase())
	fmt.Println(initializer.BindApiEngine())
	fmt.Println(initializer.StartApiEngine())
}