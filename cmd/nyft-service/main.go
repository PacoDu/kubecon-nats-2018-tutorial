package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/nats-io/go-nats"
	"github.com/wallyqs/kubecon-nats-2018-tutorial/pkg/nyft-service"
	"github.com/wallyqs/kubecon-nats-2018-tutorial/pkg/component"
)

func main() {
	var showHelp bool
	var showVersion bool
	var natsServers string
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: nyft-service [options...]\n\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}

	// Setup default flags
	flag.BoolVar(&showHelp, "h", false, "Show help")
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.StringVar(&natsServers, "s", nats.DefaultURL, "List of NATS Servers to connect")
	flag.StringVar(&natsServers, "servers", nats.DefaultURL, "List of NATS Servers to connect")
	flag.Parse()

	switch {
	case showHelp:
		flag.Usage()
		os.Exit(0)
	case showVersion:
		fmt.Fprintf(os.Stderr, "NYFT Service v%s\n", service.Version)
		os.Exit(0)
	}
	log.Printf("Starting NYFT Service v%s", service.Version)

	// Register component
	comp := component.NewComponent("nyft-service")

	err := comp.SetupConnectionToNATS(natsServers)
	if err != nil {
		log.Fatal(err)
	}

	s := service.Server{
		Component: comp,
	}
	err = s.SetupSubscriptions()
	if err != nil {
		log.Fatal(err)
	}

	runtime.Goexit()
}
