/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package todo

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var todoCmd = &cobra.Command{
	Use:   "todo",
	Short: "一个记录、修改任务列表的工具",
	Long:  `用go实现的一个命令行工具， 可以记录、查看、修改任务， 详情请查看:todo --help`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := todoCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

/*
在init函数中添加命令的参数标记等条件
*/
func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.todo.yaml)")

	// 添加子命令
	// fmt.Print(addCmd)
	var addCmd = GenAddCmd()
	var listCmd = GenListCmd()
	todoCmd.AddCommand(addCmd)
	todoCmd.AddCommand(listCmd)

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// todoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
