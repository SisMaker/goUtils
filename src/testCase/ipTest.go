package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		println("no ip found")
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				gInnerIP := ipnet.IP.String()
				println("IMY**************1111  ", gInnerIP)
			}
		}
	}
	//这种实现方式忽略了一个网卡可用性的问题，导致获取出来的IP可能不一定是想要的。需要通过判断net.FlagUp标志进行确认，排除掉无用的网卡。优化后的实现方式：

	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						println("IMY**************22222 ")
						fmt.Println(ipnet.IP.String())
					}
				}
			}
		}
	}

}
