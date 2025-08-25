package cmd

import (
	"flag"
	"fmt"
)

func Goodbye(args []string) {

	flagcmd := flag.NewFlagSet("goodbye", flag.ExitOnError)
	name := flagcmd.String("name", "user", "a name to say goodbye to")

	// Parse only the args for goodbye
	flagcmd.Parse(args)
	fmt.Println("Goodbye,", *name)
}
