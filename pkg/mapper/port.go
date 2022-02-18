package mapper

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Port struct {
	Network string
	Number  int
	State   string
}

func NewPort(network string, number int) *Port {
	return &Port{
		Network: network,
		Number:  number,
		State:   "unknown",
	}
}

func (p *Port) Scan(ip net.IP) {
	p.scanWithTimeout(ip, 3*time.Second)
}

func (p *Port) scanWithTimeout(ip net.IP, timeout time.Duration) {
	addr := net.TCPAddr{IP: ip, Port: p.Number}

	d := net.Dialer{Timeout: timeout}
	conn, err := d.Dial(p.Network, addr.String())

	if err != nil {
		switch t := err.(type) {
		case *net.OpError:
			if t.Timeout() {
				p.State = "filtered"
			}
			if strings.Contains(t.Error(), "connection refused") {
				p.State = "closed"
			}
			// WIP
			// if strings.Contains(t.Error(), "host is down") {
			// 	p.State = "WIP"
			// }
		default:
			p.State = err.Error()
		}

		return
	}
	defer conn.Close()

	p.State = "open"
}

func (p *Port) String() string {
	return fmt.Sprintf("%d/%s -- %s", p.Number, p.Network, p.State)
}
