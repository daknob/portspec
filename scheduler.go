package main

import (
	log "github.com/sirupsen/logrus"
)

/*
fullScan schedules a scan of all hosts against all ports, and checks the
results to ensure everything is within specifications.
*/
func fullScan() {
	for host, ports := range conf.Hosts {
		log.Printf("Scanning %s...", host)
		go scanHost(host, ports, conf.ScanPorts)
	}
}
