package parser

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
	"strings"
)

type Ifcs map[string]pcap.Interface

func (ifcs Ifcs) fmtPrint() {
	for _, ifc := range ifcs {
		addrsLenth := len(ifc.Addresses)
		if addrsLenth == 0 {
			continue
		}
		ips := make([]string, 0, addrsLenth)
		for _, addr := range ifc.Addresses {
			ips = append(ips, addr.IP.String())
		}
		fmt.Println(ifc.Name, strings.Join(ips, ","))
	}
}

func NewIfcs(pIfcs []pcap.Interface) Ifcs {
	ifcs := make(Ifcs, len(pIfcs))
	for _, ifc := range pIfcs {
		ifcs[ifc.Name] = ifc
	}
	return ifcs
}

func ParseAndPrint() {
	devices, err := parseDevices()
	if err != nil {
		log.Printf("解析网卡错误:%s", err.Error())
		return
	}
	if len(devices) == 0 {
		log.Printf("没有查找到网卡")
		return
	}
	ifcs := NewIfcs(devices)
	ifcs.fmtPrint()
}

func parseDevices() ([]pcap.Interface, error) {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		return nil, err
	}
	return devices, nil
}
