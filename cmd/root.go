package cmd

import (
	"errors"
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.io/pabloxio/gnmap/pkg/mapper"
)

const minConcurrentHosts = 1
const maxConcurrentHosts = 8
const concurrentHostsName = "concurrent-hosts"

var (
	ips             []net.IP
	tcpPorts        []int
	concurrentHosts int
	rootCmd         = newRootCmd()
)

func newRootCmd() *cobra.Command {
	return &cobra.Command{
		Short: "go Network Mapper",
		Long:  `go Network Mapper`,
		RunE:  rootHandler,
	}
}

func init() {
	rootCmd.Flags().IPSliceVar(&ips, "ips", []net.IP{{127, 0, 0, 1}}, "Target IPs")
	rootCmd.Flags().IntSliceVar(&tcpPorts, "tcp-ports", []int{21, 22, 23, 25, 80, 110, 139, 443, 445, 3389}, "TCP ports")
	rootCmd.Flags().IntVar(&concurrentHosts, concurrentHostsName, 5, fmt.Sprintf("Concurrent hosts (min %d max %d)", minConcurrentHosts, maxConcurrentHosts))
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func rootHandler(cmd *cobra.Command, args []string) error {
	if concurrentHosts < minConcurrentHosts || concurrentHosts > maxConcurrentHosts {
		return errors.New(fmt.Sprintf("%s should be %d <= %s <= %d\n", concurrentHostsName, minConcurrentHosts, concurrentHostsName, maxConcurrentHosts))
	}

	m := mapper.New(ips, tcpPorts, concurrentHosts)
	m.Run()

	fmt.Print(m)

	return nil
}
