package mapper

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

func TestPortStringer(t *testing.T) {
	port := NewPort("tcp", 80)
	expected := "80/tcp -- unknown"

	got := port.String()
	if got != expected {
		t.Errorf("got %s, want %s", got, expected)
	}
}

func TestScan(t *testing.T) {
	checkPort := func(t testing.TB, port *Port, want string) {
		t.Helper()

		if port.State != want {
			t.Errorf("got %s, but want %s port %s/%d", port.State, want, port.Network, port.Number)
		}
	}

	t.Run("closed", func(t *testing.T) {
		port := NewPort("tcp", 1122)
		localhost := net.IP{127, 0, 0, 1}

		port.Scan(localhost)
		checkPort(t, port, "closed")
	})

	t.Run("open", func(t *testing.T) {
		port := NewPort("tcp", 8183)
		server := fakeServer(port)
		defer server.Close()
		localhost := net.IP{127, 0, 0, 1}

		port.Scan(localhost)
		checkPort(t, port, "open")
	})

	t.Run("filtered", func(t *testing.T) {
		port := NewPort("tcp", 8185)
		localhost := net.IP{127, 0, 0, 1}

		port.scanWithTimeout(localhost, 1*time.Nanosecond)
		checkPort(t, port, "filtered")
	})
}

func fakeServer(port *Port) net.Listener {
	address := fmt.Sprintf("localhost:%d", port.Number)
	listener, err := net.Listen(port.Network, address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return listener
}
