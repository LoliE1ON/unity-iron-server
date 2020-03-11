package Udp

import (
	"fmt"
	"net"
)

type Player struct {
	IP   net.IP
	Port int
}

var players []Player

func Connect() {

	ServerConn, _ := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: 3000, Zone: ""})
	defer ServerConn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, _ := ServerConn.ReadFromUDP(buf)

		candidate := Player{
			IP:   addr.IP,
			Port: addr.Port,
		}
		if !Contains(players, candidate) {
			players = append(players, candidate)
		}

		fmt.Println("Received ", string(buf[0:n]), " from ", addr)
		fmt.Println("Players ", players)

		for _, player := range players {

			_, err := ServerConn.WriteToUDP([]byte(string(buf[0:n])), &net.UDPAddr{IP: player.IP, Port: 5000})
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("Send ", string(buf[0:n]), " to ", player.IP)
		}

	}

}

// Search through each element in a slice
func Contains(slice []Player, item Player) bool {

	for _, s := range slice {
		if s.IP.String() == item.IP.String() && s.Port == item.Port {
			return true
		}
	}

	return false
}
