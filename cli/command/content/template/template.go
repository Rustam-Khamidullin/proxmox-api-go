package template

import (
	"github.com/Rustam-Khamidullin/proxmox-api-go/cli/command/content"
	"github.com/spf13/cobra"
)

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "With this command you can manage Lxc container templates in proxmox",
}

func init() {
	content.ContentCmd.AddCommand(templateCmd)
}
