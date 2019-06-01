package img

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/kataras/iris/core/errors"
	"github.com/rs/zerolog/log"
	"image"
	"image/jpeg"
	"image/png"
	"strings"
)

var (
	errImgData = errors.New("empty image")
)

// ReadImage in base64 html format with prefix: data:image/png;base64,
func ReadImage(im string) (image.Image, error) {
	if len(im) < 20 {
		return nil, errImgData
	}

	coI := strings.Index(im, ",")
	if coI == -1 {
		return nil, errImgData
	}

	rawImage := string(im)[coI+1:]

	// Encoded Image DataUrl //
	unbased, _ := base64.StdEncoding.DecodeString(string(rawImage))
	res := bytes.NewReader(unbased)
	f := strings.TrimSuffix(im[5:coI], ";base64")
	log.Info().Int("size", len(im)).Str("format", f).Msg("read image")

	switch f {
	case "image/png":
		return png.Decode(res)

	case "image/jpeg":
		return jpeg.Decode(res)
	}

	return nil, fmt.Errorf("not supported image format: %s", f)
}

// ConvertToJPEG with defined quality
func JPEGwithBase64(i image.Image, qt int) (string, error) {
	in := bytes.NewBufferString("")
	if err := jpeg.Encode(in, i, &jpeg.Options{Quality: qt}); err != nil {
		return "", err
	}

	res := base64.RawStdEncoding.EncodeToString(in.Bytes())
	res = "data:image/jpeg;base64," + res
	log.Info().Int("size", len(res)).Msg("image to base64 with compression")
	return res, nil
}
