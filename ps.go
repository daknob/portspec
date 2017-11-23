package main

import (
	"errors"
	"fmt"
	"net"
	"reflect"
	"sort"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

/*
scanHost will perform a full port scan against a host, and then compare the
desired open ports to the ports scanned. It will then send a notification if
anything is not in spec.
*/
func scanHost(host string, desiredPorts, scanPorts []int) {
	/* Store all host open ports here */
	var openPorts []int

	/* Check all ports to see if they're open */
	for _, port := range scanPorts {
		res, err := scanPort(host, port)
		if err != nil {
			log.Warnf("Port Scan Error (%s:%d): %v", host, port, err)
		}
		if res == true {
			openPorts = append(openPorts, port)
			log.Printf("%s:%d is open.", host, port)
		} else {
			log.Printf("%s:%d is closed.", host, port)
		}
	}

	/* Sort slices to ensure a match */
	sort.Ints(openPorts)
	sort.Ints(desiredPorts)

	/* If they're not as expected, send a notification */
	if !reflect.DeepEqual(desiredPorts, openPorts) {
		notifyForHost(host, desiredPorts, openPorts)
	} else {
		log.Printf("Host \"%s\" is within spec.", host)
	}
}

/*
scanPort will perform a TCP connection to the host on port and will return
whether this port is open or not.
*/
func scanPort(host string, port int) (bool, error) {

	/* Start a TCP Connection */
	conn, err := net.DialTimeout(
		"tcp4",
		fmt.Sprintf("%s:%d", host, port),
		2*time.Second,
	)

	/* Port is Open */
	if err == nil {
		/* Close the Connection */
		conn.Close()

		return true, nil
	}
	/* The Connection Timed Out */
	if err.(net.Error).Timeout() == true {
		return false, nil
	}
	/* Connection Refused */
	switch e := err.(type) {
	case *net.OpError:
		if e.Op == "read" || e.Op == "write" || e.Op == "dial" {
			return false, nil
		}
	case syscall.Errno:
		if e == syscall.ECONNREFUSED {
			return false, nil
		}
	}

	/* A Network Error Occured */
	if err != nil {
		return false, err
	}

	log.Warnf("Control Flow Error. Reached Point 0x01.")
	return false, errors.New("control flow error: reached point 0x01")
}
