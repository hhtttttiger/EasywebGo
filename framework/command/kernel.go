package command

import "ASPGo/framework/cobra-1.7.0"

// AddKernelCommands will add all command/* to root command
func AddKernelCommands(root *cobra.Command) {
	// 挂载AppCommand命令
	root.AddCommand(initAppCommand())
	root.AddCommand(initSwaggerCommand())
}
