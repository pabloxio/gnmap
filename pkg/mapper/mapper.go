package mapper

import (
	"bytes"
	_ "embed"
	"fmt"
	"net"
	"os"
	"text/template"
	"time"
)

//go:embed mapper.tmpl
var mapperTemplate string

// Mapper handles scanning for multiple Hosts/Ports
type Mapper struct {
	Network   string
	Hosts     []*Host
	Ports     []int
	StartTime time.Time
}

func New(network string, ips []net.IP, ports []int) *Mapper {
	hosts := make([]*Host, len(ips))
	for i, ip := range ips {
		hosts[i] = NewHost(ip)
	}

	return &Mapper{
		Network:   network,
		Hosts:     hosts,
		Ports:     ports,
		StartTime: time.Now(),
	}
}

func (m *Mapper) Run() {
	for _, host := range m.Hosts {
		host.Run(m.Network, m.Ports)
	}
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