package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

// 1. 加载全局变量
var rootCmd = &cobra.Command{}

// 3. 执行Execute函数
func Execute() error {
	return rootCmd.Execute()
}

// 2. init函数
func init() {
	log.Println("root.go的init")
	rootCmd.AddCommand(wordCmd)
	rootCmd.AddCommand(timeCmd)
}
