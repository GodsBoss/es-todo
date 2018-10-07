package main

import (
	"fmt"
)

type commands map[string]func([]string) error

func (cmds commands) run(sub string, args []string) error {
	cmd, ok := cmds[sub]
	if !ok {
		return fmt.Errorf("unknown command %s", sub)
	}
	return cmd(args)
}

func notImplemented(sub string) func([]string) error {
	return func(_ []string) error {
		return fmt.Errorf("command %s is not yet implemented", sub)
	}
}
