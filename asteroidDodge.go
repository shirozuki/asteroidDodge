package main

import (
	"image"
	_ "image/png"
	"math/rand"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func loadPicture(path string) (pixel.Picture, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

func mapPngFiles() map[int]string {
	asteroidImgsName := map[int]string{
		1:  "gfx/asteroids/1.png",
		2:  "gfx/asteroids/2.png",
		3:  "gfx/asteroids/3.png",
		4:  "gfx/asteroids/4.png",
		5:  "gfx/asteroids/5.png",
		6:  "gfx/asteroids/6.png",
		7:  "gfx/asteroids/7.png",
		8:  "gfx/asteroids/8.png",
		9:  "gfx/asteroids/9.png",
		10: "gfx/asteroids/10.png",
		11: "gfx/asteroids/11.png",
		12: "gfx/asteroids/12.png",
		13: "gfx/asteroids/13.png",
		14: "gfx/asteroids/14.png",
		15: "gfx/asteroids/15.png",
		16: "gfx/asteroids/16.png",
		17: "gfx/asteroids/17.png",
		18: "gfx/asteroids/18.png",
		19: "gfx/asteroids/19.png",
		20: "gfx/asteroids/20.png",
	}
	return asteroidImgsName
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Asteroids!",
		Bounds: pixel.R(0, 0, 800, 600),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	win.SetSmooth(true)
	if err != nil {
		panic(err)
	}

	win.Clear(colornames.Black)

	// Zmienne sterujące przebiegiem gry
	var currLvl, fpsCount, timer, numOfAsts int
	//var xAstSpd, yAstSpd float64
	//var xPlySpd, yPlySpd float64
	//var chngAstSpdVec pixel.Vec
	//var chngAstXPosVec, changeAstYPosVec pixel.Vec
	var plyPos pixel.Vec
	var astPosArr [20]pixel.Vec
	//var chngPlyXPosVec, changePlyYPosVec pixel.Vec
	//var astImgsMap map[int]string // = mapPngFiles()

	astImgsMap := mapPngFiles()
	//println(asteroidImgsMap[1])

	// NIEPOTRZEBNE?
	//var plyMat pixel.Matrix
	//var astMatArr [20]pixel.Matrix

	var plyPic pixel.Picture
	var plySpr *pixel.Sprite
	var astPicArr [20]pixel.Picture
	var astSprArr [20]*pixel.Sprite
	//asteroidsMatrix := [20]pixel.IM

	plyPic, err = loadPicture("gfx/plyShipCenter.png")
	if err != nil {
		panic(err)
	}

	plySpr = pixel.NewSprite(plyPic, plyPic.Bounds())

	for i := 0; i < 20; i++ {
		picNum := random(1, 20)
		astPic, err := loadPicture(astImgsMap[picNum])
		if err != nil {
			panic(err)
		}
		astPicArr[i] = astPic
		astSprArr[i] = pixel.NewSprite(astPicArr[i], astPicArr[i].Bounds())
	}

	plyPos.X = 400
	plyPos.Y = 200

	for i := 0; i < 5; i++ {
		astPosArr[i].X = float64(random(100, 500))
		astPosArr[i].Y = float64(random(100, 500))
	}

	for !win.Closed() {
		win.Clear(colornames.Black)

		var plyMat pixel.Matrix
		var astMatArr [20]pixel.Matrix

		plyMat = pixel.IM
		for i := 0; i < 20; i++ {
			astMatArr[i] = pixel.IM
		}

		plyMat = plyMat.Moved(plyPos)
		for i := 0; i < 5; i++ {
			astMatArr[i] = astMatArr[i].Moved(astPosArr[i])
		}

		plySpr.Draw(win, plyMat)
		for i := 0; i < 5; i++ {
			astSprArr[i].Draw(win, astMatArr[i])
		}

		win.Update()

		
		fpsCount++
		
		if fpsCount < 3600 {
			fmt.Println("Level: 1")
		}

		if fpsCount = 3600 {
			fmt.Println("Level: 2")
			level++
		}

		if fpsCount = 7200 {
			fmt.Println("Level: 3")
			level++
		}

		if fpsCount = 10800 {
			fmt.Println("Level: 4")
			level++
		}

		if fpsCount = 14400 {
			fmt.Println("Level: 5")
			level++
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}
