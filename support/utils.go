package support

import "net"

// Catch global catch func
func Catch(err error, hook func()) {
	if err != nil {
		panic(err)
	}
	if hook != nil {
		hook()
	}
}

func LocalIP() (ipv4 string, err error) {
	var (
		addrSet []net.Addr
		addr    net.Addr
		ipNet   *net.IPNet // IP地址
		isIpNet bool
	)
	// get all network cards
	if addrSet, err = net.InterfaceAddrs(); err != nil {
		return
	}
	// take the first non lo network card ip
	for _, addr = range addrSet {
		// this network address is an ip address ipv4 ipv6
		if ipNet, isIpNet = addr.(*net.IPNet); isIpNet && !ipNet.IP.IsLoopback() {
			// skip ipv6
			if ipNet.IP.To4() != nil {
				ipv4 = ipNet.IP.String()
				return
			}
		}
	}

	return
}
