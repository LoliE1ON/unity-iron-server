package Types

import "net"

type PlayerUdp struct {
	IP   net.IP
	Port int
}
