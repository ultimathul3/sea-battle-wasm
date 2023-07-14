package assets

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

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
}

func New() *Assets {
	largeFont, mediumFont := loadFonts()

	return &Assets{
		BackgroundImages: loadBackgroundImages(),
		LargeFont:        largeFont,
		MediumFont:       mediumFont,
		ButtonTickPlayer: loadSounds(),
	}
}

func loadSounds() *audio.Player {
	sampleRate := 44100

	context := audio.NewContext(sampleRate)

	buttonTickStream, err := vorbis.DecodeWithSampleRate(sampleRate, bytes.NewReader(ButtonTickSound))
	if err != nil {
		log.Fatal(err)
	}

	buttonTickPlayer, err := context.NewPlayer(buttonTickStream)
	if err != nil {
		log.Fatal(err)
	}

	return buttonTickPlayer
}

func loadFonts() (font.Face, font.Face) {
	otf, err := opentype.Parse(MainOTF)
	if err != nil {
		log.Fatal(err)
	}

	largeFont, err := opentype.NewFace(otf, &opentype.FaceOptions{
		Size:    72,
		DPI:     144,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	mediumFont, err := opentype.NewFace(otf, &opentype.FaceOptions{
		Size:    48,
		DPI:     96,
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
