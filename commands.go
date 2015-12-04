/*
	This file contains the available CLI commands that are supported
*/

package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

// List of commands that are available
var Commands = []cli.Command{
	CommandSend,
}

// Command responsible for sending a change event to changelog.
var CommandSend = cli.Command{
	Name:        "add",
	Usage:       "changelog add [change parameters]",
	Description: `Command to add a new change (event) to changelog.`,
	Action:      SendEvent,
}

// Function that send the event
func SendEvent(c *cli.Context) {
	fmt.Print("Sending change event to changelog.")
}
