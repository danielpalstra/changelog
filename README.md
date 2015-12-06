# Changelog
Changelog is a tiny, simple application that sends "Change Log" events to elasticsearch by using the commandline. Inspired by the prezi changelog application https://github.com/prezi/changelog

# Elasticsearch (& Kibana)1
Changelog uses elasticsearch as the backend. The backend for changelog can be configured by using the `--url` parameter, but it's recommended to set the `ES_URL` environment variable. When both are not provided http://localhost:9200 will be used by default. Changelog uses the `@timestamp` field so that Kibana is supported

# Usage
Currently changelog only supports adding new events. Hit the following command for the available arguments
```
changelog add --help
```

# Building
Clone this repository
```
go get
go install
```

Add your `$GOPATH/bin` to the `$PATH` env variable and start using changelog from anywhere.


# Testing
Perform a minimal test by setting up a elasticsearch instance (we recommend this docker image) and performing the following commandline
```
go build && ./changelog add -m "Started deployment awesomeness"
```

# TODO
Too much todo
