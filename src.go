package main

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "simple resize tool"
	app.Action = run
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func run(c *cli.Context) error {
	n := c.NArg()
	if n == 0 {
		return errors.New("arguments is not specified")
	}
	// exifのrotate考慮したいですね
	// 引数で長辺の長さを指定したいが
	for _, v := range c.Args() {
		var img, res image.Image
		f, _ := os.Open(v)
		conf, format, err := image.DecodeConfig(f)
		splits := strings.Split(filepath.Base(v), ".")
		outf := fmt.Sprintf("%s-resize.jpg", splits[0])
		if err != nil {
			log.Print(fmt.Sprintf("%s is not image file", v))
			continue
		}
		f, _ = os.Open(v);
		if format == "jpg" {
			img, err = jpeg.Decode(f)
		} else if format == "gif" {
			img, err = gif.Decode(f)
		} else if format == "png" {
			img, err = png.Decode(f)
		}
		if err != nil {
			return errors.New(fmt.Sprintf("loaded error : %s", v))
		}
		f.Close();
		if conf.Width > conf.Height {
			res = resize.Resize(280, 0, img, resize.Bicubic)
		} else {
			res = resize.Resize(0, 280, img, resize.Bicubic)
		}
		out, err := os.Create(outf)
		if err != nil {
			return errors.New("")
		}
		jpeg.Encode(out, res, nil)
		out.Close()
	}
	return nil
}
