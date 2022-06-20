package console

import (
	"github.com/spf13/cobra"
	"linkr-frame/app/console/commands/mysqlcommand"
)

var rootCmd = &cobra.Command{}

func init() {
	//注册需要执行的命令
	rootCmd.AddCommand(mysqlcommand.SqlCmd)
}

//Execute 执行相应命令
func Execute() error {
	return rootCmd.Execute()
}
