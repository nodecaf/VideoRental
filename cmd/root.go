package cmd

import (
	"log"
	"net/http"

	"github.com/nodecaf/VideoRental/pkg/apiserver"
	"github.com/nodecaf/VideoRental/pkg/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bbuster",
	Short: "Starts bbusters client or api server",
	Long:  "To start the client or api server, used keyword client/apiserver",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var serverCmd = &cobra.Command{
	Use:     "apiserver",
	Aliases: []string{"server"},
	Short:   "Starts API server",
	Long:    "Really starts the API server on port 8080",
	Run: func(cmd *cobra.Command, args []string) {

		r := apiserver.CreateRoutes()
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Fatalln(err)
		}
	},
}
var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Starts client server",
	Long:  "Really starts the client server on port 50080",
	Run: func(cmd *cobra.Command, args []string) {

		r := client.CreatePages(apis)
		err := http.ListenAndServe(":50080", r)
		if err != nil {
			log.Fatalln(err)
		}
	},
}

func Execute() {
	rootCmd.Execute()
}

var apis string

func init() {
	clientCmd.Flags().StringVarP(&apis, "APIserver", "s", "localhost", "Add apiservers address")
	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(serverCmd)

}
