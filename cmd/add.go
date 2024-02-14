package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/xnpltn/pScan/scan"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <host1>...<hostn>",
	Short: "Add new host(s) to list",
	Long: `A longer description of Add new host(s) to list`,
	Aliases: []string{"a"},
	Args: cobra.MinimumNArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string)error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err!= nil{
			return err
		}
		return addAction(os.Stdout, hostsFile, args)
	},
}
func addAction(out io.Writer, hostfile string, args []string) error{
	hl := &scan.HostsList{}
	if err:= hl.Load(hostfile); err!= nil{
		return err
	}
	for _, item := range args{
		err := hl.Add(item)
		if err!= nil{
			return err
		}
		fmt.Fprintln(out, "Added host:", item)
	}
	return hl.Save(hostfile)
}

func init() {
	hostsCmd.AddCommand(addCmd)
}
