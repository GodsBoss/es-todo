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
		fmt.Println(" list")
		return nil
	}
	cmds := commands{
		"init":   initFile,
		"add":    task(toAddCommand),
		"reword": task(toRewordCommand),
		"cancel": task(toStateCommand(cancel)),
		"finish": task(toStateCommand(finish)),
		"remove": task(toStateCommand(removeOld)),
		"list":   list,
	}
	return cmds.run(args[0], args[1:])
}
