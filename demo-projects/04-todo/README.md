# 说明
使用 cobra 库实现一个todo功能的 cli工具

cobra 使用说明：
https://github.com/spf13/cobra-cli/blob/main/README.md
https://github.com/spf13/cobra/blob/main/site/content/user_guide.md

知乎说明
https://zhuanlan.zhihu.com/p/627848739

# 使用
## 初始化项目
```
go mode init <moduleName>
```
## 安装依赖
```
go install github.com/spf13/cobra-cli@latest
```
## 初始化cobra项目

```
cobra-cli init
```

# 描述一个命令
## 添加一个子命令
```go
// 这里的add 属于子命令
eg:
    todo add


var versionCmd = &cobra.Command{
    Use:   "version",
    Short: "Print the version number of Hugo",
    Long:  `All software has versions. This is Hugo's`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
    },
}

func init() {
    rootCmd.AddCommand(versionCmd)
}
```
## 标志
