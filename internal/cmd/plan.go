package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/mellena1/photo-resizer/pkg/photos"

	"github.com/mellena1/photo-resizer/internal/helpers"
	"github.com/spf13/cobra"
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
		panic("Must set fileglob (-g)")
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
