package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"
	"github.io/pabloxio/gnmap/pkg/mapper"
)

var (
	tcpPorts []int
	ips      []net.IP

	rootCmd = &cobra.Command{
		Short: "go Network Mapper",
		Long:  `go Network Mapper`,
		Run:   rootHandler,
	}
)

func init() {
	rootCmd.Flags().IPSliceVar(&ips, "ips", []net.IP{{127, 0, 0, 1}}, "Target IPs")
	rootCmd.Flags().IntSliceVar(&tcpPorts, "tcp-ports", []int{21, 22, 23, 25, 80, 110, 139, 443, 445, 3389}, "TCP ports")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func rootHandler(cmd *cobra.Command, args []string) {
	m := mapper.New(ips, tcpPorts)
	m.Run()

	fmt.Print(m)
}
