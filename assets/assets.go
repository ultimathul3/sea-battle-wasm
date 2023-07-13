package assets

import (
	"embed"
)

var (
	//go:embed images/backgrounds
	BackgroundImagesDir     embed.FS
	BackgroundImagesDirPath = "images/backgrounds"
)

var (
	//go:embed fonts/main.otf
	MainOTF []byte
)
