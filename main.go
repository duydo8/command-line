/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/abiosoft/ishell/v2"
	"strings"
)

func main() {

	shell := ishell.New()

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "greet",
		Help: "greet user",
		Func: func(c *ishell.Context) {
			shell.AddCmd(&ishell.Cmd{
				Name: "greetchild",
				Help: "greet user",
				Func: func(c *ishell.Context) {
					c.Println("Hello", strings.Join(c.Args, " "))
				},
			})
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "greet1",
		Help: "greet user",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})

	shell.Run()

}
