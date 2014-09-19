package main

import (
	"fmt"
	"net"
	"os"
)

/*
 * 将string类型的ip地址转换为IP对象
 */
/*func main() {
	name := "192.168.52.47"

	ip := net.ParseIP(name)

	if ip == nil {
		fmt.Fprintf(os.Stderr, "Err:无效的地址")
		return
	}

	fmt.Fprintf(os.Stdout, "IP: %s %s\n", ip, ip.String())
	defaultMask

	:= ip.DefaultMask()
	fmt.Fprintf(os.Stdout, "DefaultMask: %s %s\n", defaultMask, defaultMask.String())

	ones, bits := defaultMask.Size()
	fmt.Fprintf(os.Stdout, "ones: %d bits: %d\n", ones, bits)

	ip4 := ip.To4()
	fmt.Fprintf(os.Stdout, "to4: %s\n", ip4)
	ip16 := ip.To16()
	fmt.Fprintf(os.Stdout, "to16: %s", ip16)
}
*/

/**
 * 根据IP和掩码获得网络
 */
/*func main() {
	name := "192.168.52.47"

	ip := net.ParseIP(name)

	mask := ip.DefaultMask()

	network := ip.Mask(mask)

	fmt.Fprintf(os.Stdout, "network: %s", network.String())
}*/

/**
 *	将域名解析Ip地址
 *  获得域名对应的所有Ip地址
 */
/*func main() {
	// domain := "data.5054399.net"
	domain := "www.baidu.com"
	ipAddr, err := net.ResolveIPAddr("ip", domain)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "%s IP: %s Network: %s Zone: %s\n", ipAddr.String(), ipAddr.IP, ipAddr.Network(), ipAddr.Zone)

	// 百度 虽然只有一个域名 但实际上 其对应电信 网通 联通等又有多个IP地址
	ns, err := net.LookupHost(domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	for _, n := range ns {
		fmt.Fprintf(os.Stdout, "%s\n", n)
	}
}
*/

/**
 * 查看主机服务器占用的端口，这些服务，都是tcp或者udp的
 */
/*func main() {
	port, err := net.LookupPort("tcp", "telnet")

	if err != nil {
		fmt.Fprintf(os.Stderr, "未找到指定服务")
		return
	}

	fmt.Fprintf(os.Stdout, "telnet port: %d", port)
}*/

/*func main() {
	pTCPAddr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}

	fmt.Fprintf(os.Stdout, "www.baidu.com:80 IP: %s PORT: %d", pTCPAddr.IP.String(), pTCPAddr.Port)
}
*/
