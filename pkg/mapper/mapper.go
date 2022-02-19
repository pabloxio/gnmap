package mapper

import (
	"bytes"
	_ "embed"
	"fmt"
	"net"
	"os"
	"sync"
	"text/template"
	"time"
)

//go:embed mapper.tmpl
var mapperTemplate string

// Mapper handles scanning for multiple Hosts/Ports
type Mapper struct {
	StartTime       time.Time
	Hosts           []*Host
	tcpPortNumbers  []int
	concurrentHosts int
}

func New(ips []net.IP, tcpPortNumbers []int, concurrentHosts int) *Mapper {
	hosts := make([]*Host, len(ips))
	for i, ip := range ips {
		hosts[i] = NewHost(ip)
	}

	return &Mapper{
		StartTime:       time.Now(),
		Hosts:           hosts,
		tcpPortNumbers:  tcpPortNumbers,
		concurrentHosts: concurrentHosts,
	}
}

func (m *Mapper) Run() {
	var wg sync.WaitGroup

	// Creating hostWorkers based on m.concurrentHosts
	hostRunChannel := make(chan hostRunConfig, m.concurrentHosts) //buffered
	for i := 0; i < m.concurrentHosts; i++ {
		wg.Add(1)
		go hostWorker(&wg, hostRunChannel)
	}

	// Sending host
	for _, host := range m.Hosts {
		hostRunChannel <- hostRunConfig{host, "tcp", m.tcpPortNumbers}
	}

	close(hostRunChannel)
	wg.Wait()
}

func (m *Mapper) String() string {
	funcMap := template.FuncMap{"since": time.Since}
	tmpl := template.Must(template.New("mapper").Funcs(funcMap).Parse(mapperTemplate))
	var output bytes.Buffer

	if err := tmpl.Execute(&output, m); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return output.String()
}
