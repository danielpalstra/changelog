package main

import (
	"time"
	//"encoding/json"
	// "net/http"
	"fmt"
	"strings"

	"gopkg.in/olivere/elastic.v3"
)

// ES 1.6
// Kibana 4.0.2

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

// machine=default; ssh -i ~/.docker/machine/machines/$machine/id_rsa docker@$(docker-machine ip $machine) -L 9200:localhost:9200

func main() {

	index := strings.Join([]string{"changelog-", time.Now().Format("2006.01.02")}, "")

	fmt.Print(index)
	// Obtain a client and connect to the default Elasticsearch installation
	// Disable sniffing for https://github.com/olivere/elastic/issues/57
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.99.100:9200"), elastic.SetSniff(false))
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
	event := Event{User: "olivere", Message: "Take Five"}
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
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

}
