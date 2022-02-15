package cmd

import (
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

var (
	scan  string
	ports []int
	ips   []net.IP

	rootCmd = &cobra.Command{
		Short: "go Network Mapper.",
		Long:  `go Network Mapper (for personal learning purposes).`,
		Run:   rootHandler,
	}
)

func init() {
	rootCmd.Flags().IPSliceVar(&ips, "ips", []net.IP{{127, 0, 0, 1}}, "Target IPs")
	rootCmd.Flags().IntSliceVar(&ports, "ports", []int{21, 22, 23, 25, 80, 110, 139, 443, 445, 3389}, "Target ports")
	rootCmd.Flags().StringVar(&scan, "scan", "tcp", "Scan type (tcp or udp)")
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func rootHandler(cmd *cobra.Command, args []string) {
	if scan != "tcp" {
		fmt.Printf("Not a valid scan type: %s\n", scan)
		os.Exit(1)
	}

	fmt.Printf("Running %s scanning [WIP]\n", scan)
	fmt.Printf("ips: %v\n", ips)
	fmt.Printf("ports: %v\n", ports)
}
