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
	CommandAddChangeEvent,
}

// Command responsible for sending a change event to changelog.
var CommandAddChangeEvent = cli.Command{
	Name:        "add",
	Usage:       "changelog add [change event parameters]",
	Description: `Command to add a new change (event) to changelog. See help for the available parameters`,
	Action:      handleAddChangeEvent,
}

// Function that send the event
func handleAddChangeEvent(c *cli.Context) {
	fmt.Print("Sending change event to changelog.")
}
