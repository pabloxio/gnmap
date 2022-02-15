package mapper

import "net"

type Host struct {
	IP    net.IP
	Ports []*Port
}

func NewHost(ip net.IP) *Host {
	ports := make([]*Port, 0)
	return &Host{IP: ip, Ports: ports}
}

func (h *Host) Run(network string, ports []int) {
	for _, portNumber := range ports {
		p := NewPort(network, portNumber)
		p.Scan(h.IP)
		h.Ports = append(h.Ports, p)
	}
}

func (h *Host) String() string {
	return h.IP.String()
}
