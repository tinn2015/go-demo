package todo

import (
	"fmt"
	"os"

	"github.com/gookit/color"

	"github.com/spf13/cobra"
)

func GenClearCmd() *cobra.Command {
	var clearCmd = &cobra.Command{
		Use:   "clear",
		Short: "清空所有任务",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run clear...")

			// 执行add 操作
			_, err := os.ReadFile("todo.json")
			if err != nil {
				// 判断文件是否存在
				if os.IsNotExist(err) {
					file, _ := os.OpenFile("todo.json", os.O_CREATE|os.O_RDWR, 0644)
					defer file.Close()
					file.Write([]byte("[]"))
				}
				fmt.Println("error", err)
				return
			} else {
				file, _ := os.OpenFile("todo.json", os.O_CREATE|os.O_TRUNC, 0644)
				defer file.Close()
				file.Write([]byte("[]"))
			}
			fmt.Println(color.Yellow.Render("清空任务列表成功..."))
		},
	}
	return clearCmd
}
