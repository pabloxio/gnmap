package mapper

import (
	"fmt"
	"net"
	"strings"
	"testing"
)

func TestMapperRun(t *testing.T) {
	tests := []struct {
		port *Port
		want string
	}{
		{port: NewPort("tcp", 2385), want: "closed"},
		{port: NewPort("tcp", 2392), want: "open"},
		{port: NewPort("tcp", 2400), want: "open"},
		{port: NewPort("tcp", 5508), want: "closed"},
		{port: NewPort("tcp", 8765), want: "closed"},
		{port: NewPort("tcp", 2394), want: "open"},
	}

	portNumbers := make([]int, 0)

	for _, test := range tests {
		if test.want == "open" {
			server := fakeServer(test.port)
			defer server.Close()
		}
		portNumbers = append(portNumbers, test.port.Number)
	}

	localhost := net.IP{127, 0, 0, 1}
	m := New("tcp", []net.IP{localhost}, portNumbers)

	m.Run()

	for _, test := range tests {
		expected := fmt.Sprintf("%d/%s -- %s", test.port.Number, test.port.Network, test.want)
		if !strings.Contains(m.String(), expected) {
			t.Errorf("got %s/%d %s, but want %s", test.port.Network, test.port.Number, test.port.State, test.want)
		}
	}
}
