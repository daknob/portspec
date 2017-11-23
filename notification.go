package main

import (
	log "github.com/sirupsen/logrus"
)

/*
notifyForHost sends a notification that a particular host's open ports
do not match the expected ones specified in the configuration file.
*/
func notifyForHost(host string, desiredPorts, openPorts []int) {
	log.Errorf("Host %s port mismatch! Expected: %v. Got: %v.", host, desiredPorts, openPorts)
}
