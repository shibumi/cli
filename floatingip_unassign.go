package cli

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/spf13/cobra"
)

func newFloatingIPUnassignCommand(cli *CLI) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   "unassign [FLAGS] FLOATINGIP",
		Short:                 "Unassign a Floating IP",
		Args:                  cobra.ExactArgs(1),
		TraverseChildren:      true,
		DisableFlagsInUseLine: true,
		PreRunE:               cli.ensureActiveContext,
		RunE:                  cli.wrap(runFloatingIPUnassign),
	}
	return cmd
}

func runFloatingIPUnassign(cli *CLI, cmd *cobra.Command, args []string) error {
	floatingIPID, err := strconv.Atoi(args[0])
	if err != nil {
		return errors.New("invalid Floating IP ID")
	}
	floatingIP := &hcloud.FloatingIP{ID: floatingIPID}

	action, _, err := cli.Client().FloatingIP.Unassign(cli.Context, floatingIP)
	if err != nil {
		return err
	}
	errCh, _ := waitAction(cli.Context, cli.Client(), action)
	if err := <-errCh; err != nil {
		return err
	}
	fmt.Printf("Floating IP %d unassigned\n", floatingIP.ID)
	return nil
}
