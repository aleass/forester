package pkg

import (
	"net"
	"strings"
)

// InternalIP return internal ip.
func InternalIP() string {
	inters, err := net.Interfaces()
	if err != nil {
		return ""
	}
	for _, inter := range inters {
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			adders, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range adders {
				if aspnet, ok := addr.(*net.IPNet); ok && !aspnet.IP.IsLoopback() {
					if aspnet.IP.To4() != nil {
						return aspnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}
