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
	MainOTF        []byte
	LargeFontSize  = 72.0
	LargeFontDPI   = 144.0
	MediumFontSize = 48.0
	MediumFontDPI  = 96.0
)

var (
	//go:embed sounds/button_tick.ogg
	ButtonTickSound  []byte
	SoundsSampleRate = 44100
)

var (
	//go:embed images/ships/1.png
	SingleDeckShipImage []byte
	//go:embed images/ships/2.png
	DoubleDeckShipImage []byte
	//go:embed images/ships/3.png
	ThreeDeckShipImage []byte
	//go:embed images/ships/4.png
	FourDeckShipImage []byte
)

var (
	//go:embed images/field/field.png
	FieldImage []byte
	//go:embed images/select.png
	SelectImage []byte
)
