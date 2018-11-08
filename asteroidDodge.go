﻿package main

import (
	"fmt"
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
	var currLvl, fpsCount, numOfAsts, astMinSpd, astMaxSpd int
	//var xAstSpd, yAstSpd float64
	//var xPlySpd, yPlySpd float64
	//var chngAstSpdVec pixel.Vec
	//var chngAstXPosVec, changeAstYPosVec pixel.Vec
	var plyPos pixel.Vec
	var astPosArr [20]pixel.Vec
	var astSpdArr [20]float64
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

	for i := 0; i < 20; i++ {
		astPosArr[i].X = float64(random(20, 780))
		astPosArr[i].Y = float64(random(560, 580))
	}

	for i := 0; i < 20; i++ {
		astSpdArr[i] = float64(random(15, 30)) / 10
		fmt.Println(astSpdArr[i])
	}

	for !win.Closed() {
		win.Clear(colornames.Black)

		if currLvl == 0 {
			numOfAsts = 4
			astMinSpd = 15
			astMaxSpd = 30
		} else if currLvl == 1 {
			numOfAsts = 5
			astMinSpd = 20
			astMaxSpd = 45
		} else if currLvl == 2 {
			numOfAsts = 6
			astMinSpd = 35
			astMaxSpd = 60
		} else if currLvl == 3 {
			numOfAsts = 8
			astMinSpd = 40
			astMaxSpd = 75
		} else if currLvl == 4 {
			numOfAsts = 10
			astMinSpd = 50
			astMaxSpd = 90
		}

		if win.Pressed(pixelgl.KeyRight) {
			if plyPos.X < 780 {
				plyPos.X += 3
			}
		}
		if win.Pressed(pixelgl.KeyLeft) {
			if plyPos.X > 20 {
				plyPos.X -= 3
			}
		}
		if win.Pressed(pixelgl.KeyUp) {
			if plyPos.Y < 580 {
				plyPos.Y += 3
			}
		}
		if win.Pressed(pixelgl.KeyDown) {
			if plyPos.Y > 20 {
				plyPos.Y -= 3
			}
		}

		var plyMat pixel.Matrix
		var astMatArr [20]pixel.Matrix

		for i := 0; i < numOfAsts; i++ {
			astPosArr[i].Y -= astSpdArr[i]
		}

		plyMat = pixel.IM
		for i := 0; i < 20; i++ {
			astMatArr[i] = pixel.IM
		}

		plyMat = plyMat.Moved(plyPos)
		for i := 0; i < numOfAsts; i++ {
			astMatArr[i] = astMatArr[i].Moved(astPosArr[i])
		}

		plySpr.Draw(win, plyMat)
		for i := 0; i < numOfAsts; i++ {
			astSprArr[i].Draw(win, astMatArr[i])
		}

		win.Update()

		// SPRÓBOWAĆ ZMIENIĆ SPRITE?
		for i := 0; i < numOfAsts; i++ {
			if astPosArr[i].Y < 0 {
				astPosArr[i].X = float64(random(20, 780))
				astPosArr[i].Y = float64(random(650, 800))
				//astSpdArr[i] = float64(random(astMaxSpd-15, astMaxSpd)) / 10
				astSpdArr[i] = float64(random(astMinSpd, astMaxSpd)) / 10
			}
		}

		//fmt.Println(astPosArr[1])
		//fmt.Println("frame: ", fpsCount)
		fpsCount++

		if fpsCount == 0 {
			fmt.Println("Level: 1")
		}

		if fpsCount == 3600 {
			fmt.Println("Level: 2")
			currLvl++
		}

		if fpsCount == 7200 {
			fmt.Println("Level: 3")
			currLvl++
		}

		if fpsCount == 10800 {
			fmt.Println("Level: 4")
			currLvl++
		}

		if fpsCount == 14400 {
			fmt.Println("Level: 5")
			currLvl++
		}
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}
