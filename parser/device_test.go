package parser

import (
	"fmt"
	"log"
	"testing"
)

func TestMonitor(t *testing.T) {
	Monitor()
}

func TestParseDevices(t *testing.T) {
	devices, err := parseDevices()
	if err != nil {
		log.Fatal(err)
	}
	for _, d := range devices {
		ips := ""
		for _, a := range d.Addresses {
			ips = fmt.Sprint(ips, " ", a.IP)
		}
		t.Log(d.Name, ": ", ips)
	}
}
