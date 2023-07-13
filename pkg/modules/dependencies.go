package modules

const (
	// All includes all modules necessary for Grafana to run as a standalone application.
	All string = "all"
	// BackgroundServices includes all Grafana services that run in the background
	BackgroundServices string = "background-services"
	// Provisioning sets up Grafana with preconfigured datasources, dashboards, etc.
	Provisioning string = "provisioning"
)

// dependencyMap defines Module Targets => Dependencies
var dependencyMap = map[string][]string{
	All: {BackgroundServices, Provisioning},
}
