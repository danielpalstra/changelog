/*
	This file contains the available CLI commands that are supported
*/

package main

import (
	"errors"

	"github.com/codegangsta/cli"
)

// Commands -  List of commands that are available
var Commands = []cli.Command{
	CommandAddChangeEvent,
}

// CommandAddChangeEvent - Command responsible for sending a change event to changelog.
var CommandAddChangeEvent = cli.Command{
	Name:        "add",
	Usage:       "changelog add [change event parameters]",
	Description: `Command to add a new change (event) to changelog. See help for the available parameters`,
	Action:      addChangeEvent,
	Flags:       Flags,
}

// Function that send the event
func addChangeEvent(c *cli.Context) {

	err := checkArguments(c)
	if err != nil {
		panic(err)
	}
	changeEvent(c.String("url"), event)

}

func checkArguments(c *cli.Context) (err error) {

	if c.String("message") == "" {
		err = errors.New("message argument cannot be empty")
	}

	return err
}
