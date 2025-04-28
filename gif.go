package geom

import (
	"image"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"log"
	"os"
)

func CreateGif(images []image.Image, filename string) error {
	imgs := make([]*image.Paletted, 0)
	delays := make([]int, 0)
	for _, img := range images {
		bounds := img.Bounds()
		dst := image.NewPaletted(bounds, palette.WebSafe)
		draw.Draw(dst, bounds, img, bounds.Min, draw.Src)
		imgs = append(imgs, dst)
		delays = append(delays, 5)
	}

	f, _ := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0600)
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal("failed to close file")
		}
	}()
	return gif.EncodeAll(f, &gif.GIF{
		Image: imgs,
		Delay: delays,
	})
}
