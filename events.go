package main

import (
	"fmt"
	"strings"
	"time"

	"gopkg.in/olivere/elastic.v3"
)

// TODO move all ElasticSearch client code to own file

// Event is a structure used for serializing/deserializing data in Elasticsearch.
type Event struct {
	Timestamp   time.Time             `json:"@timestamp"`
	Category    string                `json:"category"`
	Severity    string                `json:"severity"`
	User        string                `json:"user,omitempty"`
	Message     string                `json:"message"`
	Project     string                `json:"project,omitempty"`
	Action      string                `json:"action,omitempty"`
	Hostname    string                `json:"hostname,omitempty"`
	Environment string                `json:"environment,omitempty"`
	Origin      string                `json:"origin,omitempty"`
	Tags        []string              `json:"tags,omitempty"`
	Suggest     *elastic.SuggestField `json:"suggest_field,omitempty"`
}

func getIndex() (i string) {
	// TODO make configurable
	i = strings.Join([]string{"changelog-", time.Now().Format("2006.01.02")}, "")
	return
}

// TODO cleanup code
// Method that sends a new changelog event
func changeEvent(url string, event Event) {

	// Set timestamp for event
	event.Timestamp = time.Now()

	// Define index name
	index := getIndex()

	// TODO move all ES related code to it's own file
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	// Use the IndexExists service to check if the specified index exists.
	exists, err := client.IndexExists(index).Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex(index).Do()
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	// Index a changelog event (using JSON serialization)
	put1, err := client.Index().
		Index(index).
		Type("event").
		Id("1").
		BodyJson(event).
		Do()
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Send changelog event %s to index %s \n", put1.Id, put1.Index)

}
