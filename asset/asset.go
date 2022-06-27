package asset

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

var (
	BasketcarsImage  *ebiten.Image
	BasketfoodImage  *ebiten.Image
	BasketplushImage *ebiten.Image
	CarImage1        *ebiten.Image
	CarImage2        *ebiten.Image
	CarImage3        *ebiten.Image
	CarImage4        *ebiten.Image
	FoodImage1       *ebiten.Image
	FoodImage2       *ebiten.Image
	FoodImage3       *ebiten.Image
	FoodImage4       *ebiten.Image
	PlushImage1      *ebiten.Image
	PlushImage2      *ebiten.Image
	PlushImage3      *ebiten.Image
	PlushImage4      *ebiten.Image
	PatienceImage1   *ebiten.Image
	PatienceImage2   *ebiten.Image
	PatienceImage3   *ebiten.Image
	PatienceImage4   *ebiten.Image
	PatienceImage5   *ebiten.Image
	PatienceImage6   *ebiten.Image
	StaminaImage1    *ebiten.Image
	StaminaImage2    *ebiten.Image
	StaminaImage3    *ebiten.Image
	StaminaImage4    *ebiten.Image
	StaminaImage5    *ebiten.Image
	StaminaImage6    *ebiten.Image
	KidLeftImage     *ebiten.Image
	KidRightImage    *ebiten.Image
	KidStillImage    *ebiten.Image
	KidUpImage       *ebiten.Image
	WonImage         *ebiten.Image
	LossImage        *ebiten.Image
)

var (
	AudioContext *audio.Context
	PickWav      *audio.Player
	RightBasket  *audio.Player
	WrongBasket  *audio.Player
)

const (
	SampleRate = 48000
)

func LoadAudio() {
	AudioContext = audio.NewContext(SampleRate)

	temp, err := wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(pickup_wav))
	if err != nil {
		panic(err)
	}
	PickWav, err = AudioContext.NewPlayer(temp)
	if err != nil {
		panic(err)
	}
	temp, err = wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(wrongbasket_wav))
	if err != nil {
		panic(err)
	}
	WrongBasket, err = AudioContext.NewPlayer(temp)
	if err != nil {
		panic(err)
	}
	temp, err = wav.DecodeWithSampleRate(SampleRate, bytes.NewReader(rightbasket_wav))
	if err != nil {
		panic(err)
	}
	RightBasket, err = AudioContext.NewPlayer(temp)
	if err != nil {
		panic(err)
	}
}

func LoadImageFromBytes(b []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

func LoadStaticImages() {
	BasketcarsImage = LoadImageFromBytes(basketcars_png)
	BasketfoodImage = LoadImageFromBytes(basketfood_png)
	BasketplushImage = LoadImageFromBytes(basketplush_png)
	CarImage1 = LoadImageFromBytes(car1_png)
	CarImage2 = LoadImageFromBytes(car2_png)
	CarImage3 = LoadImageFromBytes(car3_png)
	CarImage4 = LoadImageFromBytes(car4_png)
	FoodImage1 = LoadImageFromBytes(food1_png)
	FoodImage2 = LoadImageFromBytes(food2_png)
	FoodImage3 = LoadImageFromBytes(food3_png)
	FoodImage4 = LoadImageFromBytes(food4_png)
	PlushImage1 = LoadImageFromBytes(plush1_png)
	PlushImage2 = LoadImageFromBytes(plush2_png)
	PlushImage3 = LoadImageFromBytes(plush3_png)
	PlushImage4 = LoadImageFromBytes(plush4_png)
	PatienceImage1 = LoadImageFromBytes(patience1_png)
	PatienceImage2 = LoadImageFromBytes(patience2_png)
	PatienceImage3 = LoadImageFromBytes(patience3_png)
	PatienceImage4 = LoadImageFromBytes(patience4_png)
	PatienceImage5 = LoadImageFromBytes(patience5_png)
	PatienceImage6 = LoadImageFromBytes(patience6_png)
	StaminaImage1 = LoadImageFromBytes(stamina1_png)
	StaminaImage2 = LoadImageFromBytes(stamina2_png)
	StaminaImage3 = LoadImageFromBytes(stamina3_png)
	StaminaImage4 = LoadImageFromBytes(stamina4_png)
	StaminaImage5 = LoadImageFromBytes(stamina5_png)
	StaminaImage6 = LoadImageFromBytes(stamina6_png)
	KidLeftImage = LoadImageFromBytes(kidleft_png)
	KidRightImage = LoadImageFromBytes(kidright_png)
	KidStillImage = LoadImageFromBytes(kidstill_png)
	KidUpImage = LoadImageFromBytes(kidup_png)
	WonImage = LoadImageFromBytes(won_png)
	LossImage = LoadImageFromBytes(loss_png)
}
