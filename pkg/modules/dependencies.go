package modules

const (
	// All includes all modules necessary for Grafana to run as a standalone application.
	All                string = "all"
	BackgroundServices string = "background-services"
	HTTPServer         string = "http-server"
)

// dependencyMap defines Module Targets => Dependencies
var dependencyMap = map[string][]string{
	All: {BackgroundServices, HTTPServer},
}
