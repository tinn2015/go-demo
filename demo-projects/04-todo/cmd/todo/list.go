package todo

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gookit/color"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// type TodoItem struct {
// 	Id     string `json:"id"`
// 	Date   string `json:"date"`
// 	Task   string `json:"task"`
// 	IsDone bool   `json:"isDone"`
// }

// type TodoList []TodoItem

// type Person struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Email string `json:"email"`
// }

func GenListCmd() *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "查看当前未完成的任务",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run list...")

			// var person Person
			// data := []byte(`{"name":"John Doe","age":30,"email":"johndoe@example.com"}`)
			// err := json.Unmarshal(data, &person)
			// if err != nil {
			// 	return
			// }
			// fmt.Println("demo", person)

			// 执行add 操作
			f, err := os.ReadFile("todo.json")
			if err != nil {
				// 判断文件是否存在
				if os.IsNotExist(err) {
					file, _ := os.OpenFile("todo.json", os.O_CREATE|os.O_RDWR, 0644)
					defer file.Close()
					file.Write([]byte("[]"))
				}
				fmt.Println("还没有添加过任务！")
				return
			}

			var todoList TodoList
			jsonErr := json.Unmarshal(f, &todoList)
			if jsonErr != nil {
				fmt.Println(jsonErr)
				return
			}

			// 创建 tablewriter
			table := tablewriter.NewWriter(os.Stdout)

			// 设置表头
			table.SetHeader([]string{"DATE", "ID", "TASK"})

			// 添加数据行
			for _, v := range todoList {
				columnData := []string{color.Green.Render(v.Date), color.Green.Render(v.Id), color.Green.Render(v.Task)}
				table.Append(columnData)
			}

			// 设置表格格式
			table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
			table.SetColumnSeparator("|")
			table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
			table.SetAlignment(tablewriter.ALIGN_LEFT)
			table.SetAutoWrapText(false)

			// 输出表格
			table.Render()
		},
	}
	return listCmd
}
