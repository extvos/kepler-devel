package main

import (
	"fmt"
	"github.com/extvos/kepler/service"
	"github.com/spf13/cobra"
	"log"
)

// newCmd represents the version command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run a demo application powered by kepler.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("To be done...")
		listenAddr, _ := cmd.Flags().GetString("listen")
		log.Fatalln(service.Listen(listenAddr))
	},
}

func init() {
	serveCmd.Flags().StringP("listen", "L", "127.0.0.1:8080", "Demo application listen address and port.")
	serveCmd.Flags().StringP("config", "C", "", "Configuration filename.")
	// rootCmd.AddCommand(newCmd)
}
