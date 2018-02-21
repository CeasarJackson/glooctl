package upstream

import (
	"fmt"

	storage "github.com/solo-io/gloo-storage"
	"github.com/solo-io/glooctl/pkg/util"
	"github.com/spf13/cobra"
)

func deleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete [name]",
		Short: "delete upstream",
		Args:  cobra.ExactArgs(1),
		Run: func(c *cobra.Command, args []string) {
			sc, err := util.GetStorageClient(c)
			if err != nil {
				fmt.Printf("Unable to create storage client %q\n", err)
				return
			}

			if len(args) < 1 {
				fmt.Println("missing name of upstream to delete")
				return
			}
			if err := runDelete(sc, args[0]); err != nil {
				fmt.Printf("Unable to delete upstream %s: %q\n", args[0], err)
				return
			}
			fmt.Printf("Upstream %s deleted\n", args[0])
		},
	}
	return cmd
}

func runDelete(sc storage.Interface, name string) error {
	if name == "" {
		return fmt.Errorf("missing name of upstream to delete")
	}

	return sc.V1().Upstreams().Delete(name)
}
