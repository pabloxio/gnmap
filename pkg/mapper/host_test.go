package mapper

import (
	"net"
	"testing"
)

func TestHostStringer(t *testing.T) {
	var tests = []struct {
		ip   net.IP
		want string
	}{
		{ip: net.IP{127, 0, 0, 1}, want: "127.0.0.1"},
		{ip: net.IP{8, 8, 8, 8}, want: "8.8.8.8"},
	}

	for _, test := range tests {
		host := NewHost(test.ip)
		got := host.String()

		if got != test.want {
			t.Errorf("got %s want %s", got, test.want)
		}
	}
}

func TestHostRun(t *testing.T) {
	localhost := net.IP{127, 0, 0, 1}
	host := NewHost(localhost)
	portNumbers := make([]int, 0)

	tests := []struct {
		port *Port
		want string
	}{
		{port: NewPort("tcp", 2125), want: "closed"},
		{port: NewPort("tcp", 8100), want: "open"},
		{port: NewPort("tcp", 2200), want: "closed"},
		{port: NewPort("tcp", 2219), want: "open"},
	}

	for _, test := range tests {
		if test.want == "open" {
			server := fakeServer(test.port)
			defer server.Close()
		}
		portNumbers = append(portNumbers, test.port.Number)
	}

	host.Run("tcp", portNumbers)
	for i, port := range host.Ports {
		if port.State != tests[i].want {
			t.Errorf("got %s/%d %s, but want %s", port.Network, port.Number, port.State, tests[i].want)
		}
	}
}
