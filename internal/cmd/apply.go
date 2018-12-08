package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/mellena1/photo-resizer/pkg/photos"

	"github.com/mellena1/photo-resizer/internal/helpers"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

var width int
var height int
var planFileToApply string
var fileSuffix string
var helpFlag bool

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Resize the images",
	Long:  "Resizes the images given. One of plan (-f) or fileglob (-g) must be set.",
	Run:   apply,
}

func init() {
	applyCmd.PersistentFlags().IntVarP(&width, "width", "w", -1, "Width to set the images to")
	applyCmd.PersistentFlags().IntVarP(&height, "height", "h", -1, "Height to set the images to")
	applyCmd.PersistentFlags().StringVarP(&planFileToApply, "plan", "f", "", "Plan file to use")
	applyCmd.PersistentFlags().StringVar(&fileSuffix, "suffix", "", "Suffix to give to the images")
	applyCmd.PersistentFlags().BoolVar(&helpFlag, "help", false, "Help default flag")
	rootCmd.AddCommand(applyCmd)
}

func apply(cmd *cobra.Command, args []string) {
	if width == -1 || height == -1 {
		log.Errorln("Must set height (-h) and width (-w)")
		cmd.Help()
		os.Exit(1)
	}

	// Get the list of images
	images := []*photos.FileImage{}
	var err error
	if planFileToApply == "" {
		// No plan
		if fileglob == "" {
			log.Errorln("Must set either plan (-f) or fileglob (-g)")
			cmd.Help()
			os.Exit(1)
		}
		images, err = photos.FindImagesFromGlob(fileglob)
		helpers.PanicIfErr(err)
	} else {
		// Plan
		filebytes, err := ioutil.ReadFile(planFileToApply)
		helpers.PanicIfErr(err)
		unmarshalled := []*photos.FileImage{}
		json.Unmarshal(filebytes, &unmarshalled)
		for _, img := range unmarshalled {
			if err := img.LoadImage(); err == nil {
				images = append(images, img)
			} else {
				log.Warnf("Unable to load image %s from plan: %v", img.Filename, err)
			}
		}
	}

	// Resize the images
	for _, img := range images {
		var newName string
		if fileSuffix == "" {
			newName = img.Filename
		} else {
			ext := filepath.Ext(img.Filename)
			newName = img.Filename[:len(img.Filename)-len(ext)] + "-" + fileSuffix + ext
		}
		photos.Resize(img, newName, width, height)
	}
	fmt.Printf("Resized %d images.\n", len(images))
}
