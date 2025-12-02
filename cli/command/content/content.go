package content

import (
	"github.com/Rustam-Khamidullin/proxmox-api-go/cli"
	"github.com/spf13/cobra"
)

var ContentCmd = &cobra.Command{
	Use:   "content",
	Short: "With this command you can manage storage content",
}

func init() {
	cli.RootCmd.AddCommand(ContentCmd)
}
