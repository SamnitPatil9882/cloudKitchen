package main

import (
	"fmt"

	"github.com/SamnitPatil9882/cloudKitchen/repository"
)

func main() {

	//context creation
	// ctx := context.Background()

	//database creation
	_, err := repository.InitializeDatabase()
	if err != nil {
		fmt.Println("error occured in creation of db")
	}

	// r := mux.NewRouter()
	// app.Routes(r, database)

}
