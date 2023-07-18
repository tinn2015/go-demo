package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

type TodoItem struct {
	Id     string `json:"id"`
	Date   string `json:"date"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}

type TodoList []TodoItem

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func GenAddCmd() *cobra.Command {
	var content string
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "添加一条任务",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("run add...")

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
				fmt.Println("error", err)
				return
			}
			fmt.Println("文件内容", string(f))

			var todoList TodoList
			jsonErr := json.Unmarshal(f, &todoList)
			if jsonErr != nil {
				fmt.Println(jsonErr)
				return
			}
			fmt.Println("转换后文件内容", todoList)
			// 1. 往切片添加item
			// 2. 更新.json文件
			var timeNow = time.Now()
			newTask := TodoItem{
				Id:     strings.ReplaceAll(uuid.New().String(), "-", "")[:10],
				Task:   content,
				Date:   timeNow.Format("2006-01-02 15:04:05"),
				IsDone: false,
			}
			todoList = append(todoList, newTask)
			fmt.Println("添加后的list", todoList)
			todoListStr, MarshalErr := json.Marshal(todoList)
			if MarshalErr != nil {
				fmt.Println("todoList 序列化失败")
				return
			}
			file2, _ := os.OpenFile("todo.json", os.O_CREATE|os.O_RDWR, 0644)
			defer file2.Close()
			file2.Write([]byte(todoListStr))
		},
		Args: func(cmd *cobra.Command, args []string) error {
			if content == "" {
				return errors.New("请输入有效的任务")
			}
			return nil
		},
	}
	addCmd.Flags().StringVarP(&content, "content", "c", "", "添加需要记录的任务")
	return addCmd
}
