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

	conn, err := net.DialTimeout(p.Network, addr.String(), timeout)
	if err != nil {
		switch t := err.(type) {
		case *net.OpError:
			if t.Timeout() {
				p.State = "filtered"
			}
			if strings.Contains(t.Error(), "connection refused") {
				p.State = "closed"
			}
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
