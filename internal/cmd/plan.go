package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mellena1/photo-resizer/pkg/photos"

	"github.com/mellena1/photo-resizer/internal/helpers"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var planFile string

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Plan out a resize of the images",
	Run:   plan,
}

func init() {
	planCmd.PersistentFlags().StringVarP(&planFile, "output", "o", "", "Write the plan to a file")
	rootCmd.AddCommand(planCmd)
}

func plan(cmd *cobra.Command, args []string) {
	if fileglob == "" {
		log.Errorln("Must set fileglob (-g)")
		cmd.Help()
		os.Exit(1)
	}
	images, err := photos.FindImagesFromGlob(fileglob)
	helpers.PanicIfErr(err)

	output, err := json.MarshalIndent(&images, "", "  ")
	helpers.PanicIfErr(err)

	if planFile != "" {
		ioutil.WriteFile(planFile, output, 0644)
	} else {
		fmt.Println(string(output))
	}
}
