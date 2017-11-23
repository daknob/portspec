package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func main() {

	/* Check log output type */
	logOut := flag.Bool("j", false, "Output logs in JSON")

	/* Check configuration file path */
	configPath := flag.String("c", "/etc/portspec/portspec.yml", "Configuration File Path")

	/* Parse Command Line Arguments */
	flag.Parse()

	/* Set proper log format */
	if *logOut == true {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{})
	}

	log.Printf("Starting portspec...")
	log.Printf("Log JSON: %v", *logOut)
	log.Printf("Configuration File: %s", *configPath)

	/* Parse Configuration File */
	log.Printf("Parsing Configuration File...")

	conf := config{}

	confbytes, err := ioutil.ReadFile(*configPath)
	if err != nil {
		log.Fatalf("Could not open configuration file: %s", err.Error())
	}

	err = yaml.Unmarshal(confbytes, &conf)
	if err != nil {
		log.Fatalf("Could not parse configuration file: %s", err.Error())
	}

	log.Printf("Configuration File Loaded")

	/* Check if ports are provided */
	if len(conf.ScanPorts) == 0 {
		log.Fatalf("Did not provide ports to scan in the configuration file.")
	}

	/* Check the interval in the config */
	if conf.Interval == 0 {
		log.Fatalf("Did not provide a port scanning interval in the configuration file.")
	}
	if conf.Interval < 3600 {
		log.Warningf("The scanning interval is less than one hour. This may be too often.")
	}

	/* Check the hosts */
	if len(conf.Hosts) == 0 {
		log.Fatalf("Did not provide hosts to scan ports on in the configuration file.")
	}
	log.Printf("Loaded hosts:")
	for host, ports := range conf.Hosts {
		log.Printf("Host: %s, Ports: %v", host, ports)
	}

	/* Check the parallel scans */
	if conf.ParallelScans <= 0 {
		log.Fatalf("There must be at least one scan set in the configuration file.")
	}

	/* Convert Interval to Time.Duration */
	inter, err := time.ParseDuration(fmt.Sprintf("%ds", conf.Interval))
	if err != nil {
		log.Fatalf("Could not convert set interval to time duration. Maybe interval is not right in configuration file.")
	}

	/* Main Loop that runs the Port Scans */
	for {
		log.Printf("Running full port scan against all hosts...")
		go fullScan(conf)
		time.Sleep(inter)
	}
}
