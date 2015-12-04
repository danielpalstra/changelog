/*
	This file contains the available CLI commands that are supported
*/

package main

import "github.com/codegangsta/cli"

// Input for Flags
// Created     time.Time `json:"created"`
// Category    string    `json:"category"`
// Severity    string    `json:"severity"`
// User        string    `json:"user,omitempty"`
// Message     string    `json:"message"`
// Project     string    `json:"project,omitempty"`
// Action      string    `json:"action,omitempty"`
// Hostname    string    `json:"hostname,omitempty"`
// Environment string    `json:"environment,omitempty"`
// Origin      string    `json:"origin,omitempty"`
// Tags        []string  `json:"tags,omitempty"`

// List of available flags
var Flags = []cli.Flag{
	flagCategory,
	flagSeverity,
	flagUser,
	flagMessage,
	flagProject,
	flagAction,
	flagEnvironment,
	flagHostname,
}

var flagCategory = cli.StringFlag{
	Name:  "category, c",
	Usage: "category for the change event",
}

var flagSeverity = cli.StringFlag{
	Name:  "severity, s",
	Value: "5",
	Usage: "severity of the change event [1-5]",
}

var flagUser = cli.StringFlag{
	Name:  "user, u",
	Usage: "user that triggered the change event.",
}

var flagMessage = cli.StringFlag{
	Name:  "message, m",
	Usage: "message describing the change",
}

// var flag = cli.StringFlag{
// 	Name:  "",
// 	Usage: "",
// }

var flagProject = cli.StringFlag{
	Name:  "project, p",
	Usage: "project or application impacted by the change",
}

var flagAction = cli.StringFlag{
	Name:  "action, a",
	Usage: "action that is performed by the change. For example: restart, (undeployment)",
}

var flagEnvironment = cli.StringFlag{
	Name:  "environment, e",
	Usage: "environment on which the change has impact",
}

var flagHostname = cli.StringFlag{
	Name:  "hostname",
	Usage: "hostname where the change has impact on",
}
