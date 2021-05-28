/*
opfcli is a cli toolbox for maintaining an Operate First-style
Kubernetes configuration repository.

For more information, see <http://operate-first.cloud>
*/
package main

import (
	"github.com/operate-first/opfcli/cmd"
	"github.com/operate-first/opfcli/utils"
	"github.com/spf13/cobra"
)

func main() {
	utils.ConfigureLogging()

	root := cmd.NewCmdRoot()
	cobra.CheckErr(root.Execute())
}
