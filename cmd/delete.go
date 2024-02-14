package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/xnpltn/pScan/scan"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <host1>...<hostn>",
	Short: "Remove host from the list",
	Long: `A longer description of remove host from the list`,
	SilenceUsage: true,
	Args:         cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		hostsFile, err := cmd.Flags().GetString("hosts-file")
		if err!= nil{
			return err
		}
		return deleteAction(os.Stdout, hostsFile, args)
	},
}

func deleteAction(out io.Writer, hostsFile string, args []string)error{
	hl := &scan.HostsList{}
	err:= hl.Load(hostsFile)
	if err!= nil{
		return err
	}

	for _, item := range args{
		err:=hl.Remove(item)
		if err!= nil{
			return err
		}
		fmt.Fprintln(out, "Host deleted:", item)
	}

	return hl.Save(hostsFile)
}


func init() {
	hostsCmd.AddCommand(deleteCmd)
}
