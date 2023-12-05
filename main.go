package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//Using godotenv package to load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
	fmt.Println(os.Getenv("TEST"))
}
