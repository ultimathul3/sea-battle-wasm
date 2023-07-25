package assets

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Assets struct {
	BackgroundImages []*ebiten.Image
	LargeFont        font.Face
	MediumFont       font.Face

	ButtonTickPlayer *audio.Player
	HitPlayer        *audio.Player
	MissPlayer       *audio.Player
	ThemePlayer      *audio.Player

	SingleDeckShipImage *ebiten.Image
	DoubleDeckShipImage *ebiten.Image
	ThreeDeckShipImage  *ebiten.Image
	FourDeckShipImage   *ebiten.Image

	SingleDeckShipPickImage *ebiten.Image
	DoubleDeckShipPickImage *ebiten.Image
	ThreeDeckShipPickImage  *ebiten.Image
	FourDeckShipPickImage   *ebiten.Image
	PickFrameImage          *ebiten.Image

	FieldImage   *ebiten.Image
	SelectImage  *ebiten.Image
	ArrowImage   *ebiten.Image
	CurtainImage *ebiten.Image

	MissImage *ebiten.Image
	HitImage  *ebiten.Image
}

func New() *Assets {
	largeFont, mediumFont := loadFonts()

	assets := &Assets{
		BackgroundImages: loadBackgroundImages(),
		LargeFont:        largeFont,
		MediumFont:       mediumFont,

		SingleDeckShipImage: imageFromBytes(SingleDeckShipImage),
		DoubleDeckShipImage: imageFromBytes(DoubleDeckShipImage),
		ThreeDeckShipImage:  imageFromBytes(ThreeDeckShipImage),
		FourDeckShipImage:   imageFromBytes(FourDeckShipImage),

		SingleDeckShipPickImage: imageFromBytes(SingleDeckShipPickImage),
		DoubleDeckShipPickImage: imageFromBytes(DoubleDeckShipPickImage),
		ThreeDeckShipPickImage:  imageFromBytes(ThreeDeckShipPickImage),
		FourDeckShipPickImage:   imageFromBytes(FourDeckShipPickImage),
		PickFrameImage:          imageFromBytes(PickFrameImage),

		FieldImage:   imageFromBytes(FieldImage),
		SelectImage:  imageFromBytes(SelectImage),
		ArrowImage:   imageFromBytes(ArrowImage),
		CurtainImage: imageFromBytes(CurtainImage),

		MissImage: imageFromBytes(MissImage),
		HitImage:  imageFromBytes(HitImage),
	}

	assets.ButtonTickPlayer, assets.HitPlayer, assets.MissPlayer, assets.ThemePlayer = loadSounds()

	return assets
}

func loadSounds() (*audio.Player, *audio.Player, *audio.Player, *audio.Player) {
	context := audio.NewContext(SoundsSampleRate)

	buttonTickStream, err := vorbis.DecodeWithSampleRate(SoundsSampleRate, bytes.NewReader(ButtonTickSound))
	if err != nil {
		log.Fatal(err)
	}

	buttonTickPlayer, err := context.NewPlayer(buttonTickStream)
	if err != nil {
		log.Fatal(err)
	}

	hitStream, err := vorbis.DecodeWithSampleRate(SoundsSampleRate, bytes.NewReader(HitSound))
	if err != nil {
		log.Fatal(err)
	}

	hitPlayer, err := context.NewPlayer(hitStream)
	if err != nil {
		log.Fatal(err)
	}

	missStream, err := vorbis.DecodeWithSampleRate(SoundsSampleRate, bytes.NewReader(MissSound))
	if err != nil {
		log.Fatal(err)
	}

	missPlayer, err := context.NewPlayer(missStream)
	if err != nil {
		log.Fatal(err)
	}

	themeStream, err := vorbis.DecodeWithSampleRate(SoundsSampleRate, bytes.NewReader(ThemeSound))
	if err != nil {
		log.Fatal(err)
	}

	themePlayer, err := context.NewPlayer(themeStream)
	if err != nil {
		log.Fatal(err)
	}

	buttonTickPlayer.SetVolume(0.25)
	hitPlayer.SetVolume(0.09)
	missPlayer.SetVolume(0.01)
	themePlayer.SetVolume(0.01)

	return buttonTickPlayer, hitPlayer, missPlayer, themePlayer
}

func loadFonts() (font.Face, font.Face) {
	otf, err := opentype.Parse(MainOTF)
	if err != nil {
		log.Fatal(err)
	}

	largeFont, err := opentype.NewFace(otf, &opentype.FaceOptions{
		Size:    LargeFontSize,
		DPI:     LargeFontDPI,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	mediumFont, err := opentype.NewFace(otf, &opentype.FaceOptions{
		Size:    MediumFontSize,
		DPI:     MediumFontDPI,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	return largeFont, mediumFont
}

func loadBackgroundImages() []*ebiten.Image {
	var images []*ebiten.Image

	dir, err := BackgroundImagesDir.ReadDir(BackgroundImagesDirPath)
	if err != nil {
		log.Fatal(err)
	}

	for i := range dir {
		if i == 1 && os.Getenv("DEVELOPMENT") == "1" {
			break
		}

		data, err := BackgroundImagesDir.ReadFile(
			fmt.Sprintf("%s/%d.png", BackgroundImagesDirPath, i),
		)
		if err != nil {
			log.Fatal(err)
		}

		images = append(images, imageFromBytes(data))
	}

	return images
}

func imageFromBytes(imageBytes []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(imageBytes))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
