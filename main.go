/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnvFiles() {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	LoadEnvFiles()

	shell := ishell.New()

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "greet",
		Help: "greet user",
		Func: func(c *ishell.Context) {

		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "username",
		Help: "get username",
		Func: func(c *ishell.Context) {
			c.Println(os.Getenv("USERNAME1"))
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "password",
		Help: "generate password",
		Func: func(c *ishell.Context) {
			c.Println(os.Getenv("PASSWORD"))
		},
	})

	shell.Run()

}
