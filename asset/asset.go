package asset

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
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
)

func LoadImageFromBytes(pic *ebiten.Image, b []byte) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		panic(err)
	}
	pic = ebiten.NewImageFromImage(img)
}

func LoadStaticImages() {
	LoadImageFromBytes(BasketcarsImage, basketcars_png)
	LoadImageFromBytes(BasketfoodImage, basketfood_png)
	LoadImageFromBytes(BasketplushImage, basketplush_png)
	LoadImageFromBytes(CarImage1, car1_png)
	LoadImageFromBytes(CarImage2, car2_png)
	LoadImageFromBytes(CarImage3, car3_png)
	LoadImageFromBytes(CarImage4, car4_png)
	LoadImageFromBytes(FoodImage1, food1_png)
	LoadImageFromBytes(FoodImage2, food2_png)
	LoadImageFromBytes(FoodImage3, food3_png)
	LoadImageFromBytes(FoodImage4, food4_png)
	LoadImageFromBytes(PlushImage1, plush1_png)
	LoadImageFromBytes(PlushImage2, plush2_png)
	LoadImageFromBytes(PlushImage3, plush3_png)
	LoadImageFromBytes(PlushImage4, plush4_png)
	LoadImageFromBytes(PatienceImage1, patience1_png)
	LoadImageFromBytes(PatienceImage2, patience2_png)
	LoadImageFromBytes(PatienceImage3, patience3_png)
	LoadImageFromBytes(PatienceImage4, patience4_png)
	LoadImageFromBytes(PatienceImage5, patience5_png)
	LoadImageFromBytes(PatienceImage6, patience6_png)
	LoadImageFromBytes(StaminaImage1, stamina1_png)
	LoadImageFromBytes(StaminaImage2, stamina2_png)
	LoadImageFromBytes(StaminaImage3, stamina3_png)
	LoadImageFromBytes(StaminaImage4, stamina4_png)
	LoadImageFromBytes(StaminaImage5, stamina5_png)
	LoadImageFromBytes(StaminaImage6, stamina6_png)
	LoadImageFromBytes(KidLeftImage, kidleft_png)
	LoadImageFromBytes(KidRightImage, kidright_png)
	LoadImageFromBytes(KidStillImage, kidstill_png)
	LoadImageFromBytes(KidUpImage, kidup_png)
}
