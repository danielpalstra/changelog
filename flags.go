/*
	This file contains the available CLI flags for adding change events to changelog
*/

package main

import "github.com/codegangsta/cli"

var event Event

// Flags contains a list of available arguments to use
var Flags = []cli.Flag{
	flagURL,
	flagCategory,
	flagSeverity,
	flagUser,
	flagMessage,
	flagProject,
	flagAction,
	flagEnvironment,
	flagHostname,
}

var flagURL = cli.StringFlag{
	Name:   "url",
	Value:  "http://localhost:9200",
	Usage:  "URL of the ElasticSearch server",
	EnvVar: "ES_URL",
}

var flagCategory = cli.StringFlag{
	Name:        "category, c",
	Usage:       "category for the change event",
	Value:       "misc",
	Destination: &event.Category,
}

var flagSeverity = cli.StringFlag{
	Name:        "severity, s",
	Value:       "5",
	Usage:       "severity of the change event [1-5]",
	Destination: &event.Severity,
}

var flagUser = cli.StringFlag{
	Name:        "user, u",
	Value:       "unknown",
	Usage:       "user that triggered the change event.",
	Destination: &event.User,
}

var flagMessage = cli.StringFlag{
	Name:        "message, m",
	Usage:       "message describing the change",
	Destination: &event.Message,
}

var flagProject = cli.StringFlag{
	Name:        "project, p",
	Usage:       "project or application impacted by the change",
	Destination: &event.Project,
}

var flagAction = cli.StringFlag{
	Name:        "action, a",
	Usage:       "action that is performed by the change. For example: restart, (undeployment)",
	Destination: &event.Project,
}

var flagEnvironment = cli.StringFlag{
	Name:        "environment, e",
	Usage:       "environment on which the change has impact",
	Destination: &event.Environment,
}

var flagHostname = cli.StringFlag{
	Name:        "hostname",
	Usage:       "hostname where the change has impact on",
	Destination: &event.Hostname,
}
