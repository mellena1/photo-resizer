package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var fileglob string

var rootCmd = &cobra.Command{
	Use:   "photo-resizer",
	Short: "Resizes images",
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&fileglob, "fileglob", "g", "", "File glob to search with")
}

// Execute runs the rootcmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
