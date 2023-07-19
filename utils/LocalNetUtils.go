package utils

import (
	"errors"
	"fmt"
	"net"
	"os"
)

func GetFirstNonLoopbackHostInfo() (string, string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		return hostname, "", err
	}
	ips, err := net.LookupIP(hostname)
	if err != nil {
		fmt.Println(err)
		return hostname, "", err
	}

	for _, ip := range ips {
		if !ip.IsLoopback() && ip.To4() != nil {
			fmt.Println(ip)
			return hostname, ip.String(), nil
		}
	}
	return hostname, "", errors.New("no interface")
}

func PrintIps2() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

}

func PrintIps() {
	PrintIps2()

	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Failed to get network interfaces:", err)
		return
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			fmt.Println("Failed to get addresses for interface", iface.Name, ":", err)
			continue
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
				fmt.Println("Interface:", iface.Name)
				fmt.Println("IP Address:", ipNet.IP)
			}
		}
	}
}
