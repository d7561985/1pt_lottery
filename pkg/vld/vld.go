package vld

import (
	img "github.com/d7561985/1pt_lottery/pkg/image"
	"github.com/gookit/validate"
)

func init() {
	validate.AddValidators(map[string]interface{}{
		"img2": isImage,
	})
}

func isImage(in interface{}) bool {
	image, ok := in.(string)
	if !ok {
		return false
	}

	if _, err := img.ReadImage(image); err != nil {
		return false
	}

	return true
}
