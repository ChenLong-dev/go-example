/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var foo2 *string
var print2 string

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server called")
		fmt.Println("server called", "foo:", *foo)
		fmt.Println("server called", "foo2:", *foo2)
		fmt.Println("server called", "print:", print)
		fmt.Println("server called", "print2:", print2)
		fmt.Println("server called", "show:", show)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	foo2 = serverCmd.Flags().String("foo2", "fo2", "help for foo2")
	serverCmd.PersistentFlags().StringVar(&print2, "print2", "defaultPrint2", "print2")
	serverCmd.Flags().BoolP("serve", "s", false, "Help message for serve")
}
