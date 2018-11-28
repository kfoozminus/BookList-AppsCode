package cmd

import (
	"fmt"

	"github.com/kfoozminus/BookList-AppsCode/booklist"
	"github.com/spf13/cobra"
)

var port string
var RootCmd = &cobra.Command{
	Use:   "booklist",
	Short: "Booklist is RESTful API Server",
	Long:  "Booklist RESTful API was built during the training period at *, to learn Go and How to build a RESTful API using http package of golang. It is used to add, update, delete, show a booklist. It needs to authorization to modify the list.",
	Run: func(cmd *cobra.Command, args []string) {
		//booklist.main("8080")
		fmt.Println("starting the server at", port)
		booklist.Main(port)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().StringVarP(&port, "port", "p", "8080", "enter the server port")
}
