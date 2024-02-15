package cmd

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/xnpltn/pScan/scan"
)

// scanrCmd represents the scanr command
var scanrCmd = &cobra.Command{
	Use:   "scanr",
	Short: "Scan in range of ports",
	Long: `A longer description of scan in range of ports`,
	RunE: func(cmd *cobra.Command, _ []string)error {
		hostFile, err := cmd.Flags().GetString("hosts-file")
		if err != nil {
			return err
		}
		rng, err:= cmd.Flags().GetIntSlice("range")
		if err!= nil{
			return err
		}
		return scanRangeAction(os.Stdout, hostFile, rng)
	},
}

func scanRangeAction(out io.Writer, hostsFile string, rng []int)error{
	hl := &scan.HostsList{}
	if err := hl.Load(hostsFile); err != nil {
		return err
	}
	ports := []int{}
	for i:= rng[0]; i<=rng[1]; i++{
		ports = append(ports, i)
	}
	results := scan.Run(hl, ports)
	return PrintResults(out, results)
}

func init() {
	rootCmd.AddCommand(scanrCmd)
	scanrCmd.Flags().IntSliceP("range", "r", []int{88, 20}, "range of ports to scan, ex: 1,10")
}
