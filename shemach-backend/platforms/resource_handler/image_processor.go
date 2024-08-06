package resource_handler

import (
	"bytes"
	"errors"
	"os"
	"os/exec"

	"github.com/samuael/shemach/shemach-backend/pkg/constants/state"
	"github.com/samuael/shemach/shemach-backend/platforms/helper"
)

func GetBlurredImage(imagename string) (resource string, er error) {
	assetDirectory := os.Getenv("ASSETS_DIRECTORY")

	randomString := state.BLURRED_POST_IMAGES_RELATIVE_PATH + helper.GenerateRandomString(10, helper.CHARACTERS) + "." + "jpg"
	oldimage := assetDirectory + imagename

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	//convert -strip -interlace Plane -gaussian-blur 0.1 -quality 20% aMHH90.jpg result.jpg  // for Original Image
	// convert -strip -interlace Plane -quality 10% -filter Gaussian -blur 0x9  aMHH90.jpg result.jpg

	cmd := exec.Command("./../app/convert", "-strip", "-interlace", "Plane", "-quality", "1%", "-filter", "Gaussian", "-blur", "0x9", oldimage, assetDirectory+randomString)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		println(stderr.String())
		println(err.Error())
		return "", err
	}
	if stderr.String() != "" {
		return "", errors.New(stderr.String())
	}
	return randomString, nil
}
