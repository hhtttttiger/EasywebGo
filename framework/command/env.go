package command

import (
	"ASPGo/framework/cobra-1.7.0"
	"ASPGo/framework/contract"
	"fmt"
)

// envCommand 获取当前的 App 环境
var envCommand = &cobra.Command{
	Use:   "env",
	Short: "获取当前的 App 环境",
	Run: func(c *cobra.Command, args []string) {
		// 获取 env 环境
		container := c.GetContainer()
		envService := container.MustMake(contract.EnvKey).(contract.Env)
		// 打印环境
		fmt.Println("environment:", envService.AppEnv())
	},
}
