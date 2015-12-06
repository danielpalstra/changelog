package main

import (
	"os"
	//"encoding/json"
	// "net/http"

	"github.com/codegangsta/cli"
)

// ES 1.6
// Kibana 4.0.2

// machine=default; ssh -i ~/.docker/machine/machines/$machine/id_rsa docker@$(docker-machine ip $machine) -L 9200:localhost:9200

func main() {

	app := cli.NewApp()
	app.Name = "changelog"
	app.Version = "0.1.0"
	app.Usage = "changelog [action] [arguments]"
	app.Author = "Daniel Palstra"
	app.Email = "daniel@danielpalstra.com"
	app.Commands = Commands
	app.Run(os.Args)

}
