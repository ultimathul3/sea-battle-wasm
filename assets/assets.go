package assets

import (
	"embed"
)

var (
	//go:embed images/background
	BackgroundImagesDir     embed.FS
	BackgroundImagesDirPath = "images/background"
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
	ButtonTickSound []byte
	//go:embed sounds/hit.ogg
	HitSound []byte
	//go:embed sounds/miss.ogg
	MissSound []byte
	//go:embed sounds/theme.ogg
	ThemeSound []byte

	SoundsSampleRate = 44100
)

var (
	//go:embed images/field/ships/1.png
	SingleDeckShipImage []byte
	//go:embed images/field/ships/2.png
	DoubleDeckShipImage []byte
	//go:embed images/field/ships/3.png
	ThreeDeckShipImage []byte
	//go:embed images/field/ships/4.png
	FourDeckShipImage []byte
)

var (
	//go:embed images/field/pick/1.png
	SingleDeckShipPickImage []byte
	//go:embed images/field/pick/2.png
	DoubleDeckShipPickImage []byte
	//go:embed images/field/pick/3.png
	ThreeDeckShipPickImage []byte
	//go:embed images/field/pick/4.png
	FourDeckShipPickImage []byte
	//go:embed images/field/pick/frame.png
	PickFrameImage []byte
)

var (
	//go:embed images/field/field.png
	FieldImage []byte
	//go:embed images/field/select.png
	SelectImage []byte
	//go:embed images/field/arrow.png
	ArrowImage []byte
	//go:embed images/field/curtain.png
	CurtainImage []byte
)

var (
	//go:embed images/field/miss.png
	MissImage []byte
	//go:embed images/field/hit.png
	HitImage []byte
)

var (
	//go:embed images/field/explosion
	ExplosionImagesDir     embed.FS
	ExplosionImagesDirPath = "images/field/explosion"
)
