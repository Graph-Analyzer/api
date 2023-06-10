package main

import "graph-analyzer/api/web"

// Main entry point
//
//	@title			graph-analyzer API
//	@version		2.0.0
//	@description	Analyze a network based on graph properties
//	@contact.name	API Support
//	@contact.url	https://gitlab.ost.ch/graph-analyzer/api/-/issues
//	@BasePath		/api
func main() {
	web.Serve()
}
