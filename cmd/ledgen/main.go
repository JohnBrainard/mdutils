package main

import (
	"flag"
	"os"

	ctrlutils "github.com/JohnBrainard/droputils"
	"github.com/JohnBrainard/droputils/config"
)

type Options struct {
	Path string
}

var OPTS Options

func main() {
	flag.StringVar(&OPTS.Path, "path", "", "path to configuration yaml")
	flag.Parse()

	reader, err := os.Open(OPTS.Path)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := reader.Close(); err != nil {
			panic(err)
		}
	}()

	cfg, err := config.Read(reader)
	if err != nil {
		panic(err)
	}

	keyMap := ctrlutils.New(cfg)
	renderer := ctrlutils.NewKeymapRenderer(&keyMap, &cfg)
	if err = renderer.Render(os.Stdout); err != nil {
		panic(err)
	}
}
