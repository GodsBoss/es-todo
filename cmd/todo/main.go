package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		fmt.Println("Possible subcommands:")
		fmt.Println(" init")
		fmt.Println(" add <task>")
		fmt.Println(" reword <id> <task>")
		fmt.Println(" cancel <id>")
		fmt.Println(" finish <id>")
		fmt.Println(" remove <id>")
		return nil
	}
	cmds := commands{
		"init":   notImplemented("init"),
		"add":    notImplemented("add"),
		"reword": notImplemented("reword"),
		"cancel": notImplemented("cancel"),
		"finish": notImplemented("finish"),
		"remove": notImplemented("remove"),
	}
	return cmds.run(args[0], args[1:])
}
