package main

import "time"

// Event is a structure used for serializing/deserializing data in Elasticsearch.
type Event struct {
	Created     time.Time `json:"created"`
	Category    string    `json:"category"`
	Severity    string    `json:"severity"`
	User        string    `json:"user,omitempty"`
	Message     string    `json:"message"`
	Project     string    `json:"project,omitempty"`
	Action      string    `json:"action,omitempty"`
	Hostname    string    `json:"hostname,omitempty"`
	Environment string    `json:"environment,omitempty"`
	Origin      string    `json:"origin,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	// Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}
