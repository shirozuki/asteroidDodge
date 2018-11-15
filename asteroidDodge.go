package main

import (
	"fmt"
	"image"
	_ "image/png"
	"math/rand"
	"os"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
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
	var currLvl, fpsCount, numOfAsts, astMinSpd, astMaxSpd, gunTimeout, points int
	var misLaunch, gameOver bool
	var plyPos pixel.Vec
	var misPos pixel.Vec
	var astPosArr [20]pixel.Vec
	var astSpdArr [20]float64

	loadAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	loadTxt := text.New(pixel.V(670, 10), loadAtlas)
	loadTxt.Color = colornames.Orange
	fmt.Fprintln(loadTxt, "Dzialo : ")
	ptsAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	ptsTxt := text.New(pixel.V(10, 10), ptsAtlas)
	ptsTxt.Color = colornames.Yellow
	fmt.Fprintln(ptsTxt, "Punkty : ", points)

	astImgsMap := mapPngFiles()

	var titlePic pixel.Picture
	var titleSpr *pixel.Sprite
	var endPic pixel.Picture
	var endSpr *pixel.Sprite

	var plyPic pixel.Picture
	var plySpr *pixel.Sprite
	var misPic pixel.Picture
	var misSpr *pixel.Sprite
	var astPicArr [20]pixel.Picture
	var astSprArr [20]*pixel.Sprite

	titlePic, err = loadPicture("gfx/titleScreen.png")
	if err != nil {
		panic(err)
	}
	titleSpr = pixel.NewSprite(titlePic, titlePic.Bounds())

	endPic, err = loadPicture("gfx/endScreen.png")
	if err != nil {
		panic(err)
	}
	endSpr = pixel.NewSprite(endPic, endPic.Bounds())

	plyPic, err = loadPicture("gfx/plyShipCenter.png")
	if err != nil {
		panic(err)
	}
	plySpr = pixel.NewSprite(plyPic, plyPic.Bounds())

	misPic, err = loadPicture("gfx/missile2.png")
	if err != nil {
		panic(err)
	}
	misSpr = pixel.NewSprite(misPic, misPic.Bounds())

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

	misPos.X = 1000
	misPos.Y = 1000
	misLaunch = false

	for i := 0; i < 20; i++ {
		astPosArr[i].X = float64(random(20, 780))
		astPosArr[i].Y = float64(random(650, 800))
	}

	for i := 0; i < 20; i++ {
		astSpdArr[i] = float64(random(15, 30)) / 10 // 15, 30
	}

	// INTRO
	var start bool
	start = false

	for start == false {
		titleSpr.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
		win.Update()
		if win.Pressed(pixelgl.KeyEnter) {
			start = true
		}
	}

	for !win.Closed() {
		for gameOver == false {
			win.Clear(colornames.Black)

			if currLvl == 0 {
				numOfAsts = 3 //1
				astMinSpd = 8
				astMaxSpd = 14
			} else if currLvl == 1 {
				numOfAsts = 5  //5
				astMinSpd = 20 //20
				astMaxSpd = 45 //45
			} else if currLvl == 2 {
				numOfAsts = 6  //6
				astMinSpd = 35 //35
				astMaxSpd = 60 //60
			} else if currLvl == 3 {
				numOfAsts = 8  //8
				astMinSpd = 40 //40
				astMaxSpd = 75 //75
			} else if currLvl == 4 {
				numOfAsts = 10 //10
				astMinSpd = 50 //50
				astMaxSpd = 90 //90
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
			if win.Pressed(pixelgl.KeySpace) {
				if gunTimeout <= 0 {
					misLaunch = true
					misPos.X = plyPos.X / 2
					misPos.Y = plyPos.Y / 2
					gunTimeout = 200
				}
			}

			if misPos.Y > 1000 {
				misPos.X = 1000
				misPos.Y = 1000
				misLaunch = false
			}

			var plyMat pixel.Matrix
			var misMat pixel.Matrix
			var astMatArr [20]pixel.Matrix

			for i := 0; i < numOfAsts; i++ {
				astPosArr[i].Y -= astSpdArr[i]
			}

			plyMat = pixel.IM
			misMat = pixel.IM
			for i := 0; i < 20; i++ {
				astMatArr[i] = pixel.IM
			}

			plyMat = plyMat.Moved(plyPos)
			misMat = misMat.Moved(misPos)
			for i := 0; i < numOfAsts; i++ {
				astMatArr[i] = astMatArr[i].Moved(astPosArr[i])
			}

			if misLaunch == true {
				misPos.Y += 8
				misMat = misMat.Moved(misPos)
			}

			plySpr.Draw(win, plyMat)
			misSpr.Draw(win, misMat)
			for i := 0; i < numOfAsts; i++ {
				astSprArr[i].Draw(win, astMatArr[i])
			}

			// DETEKCJA KOLIZJI
			for i := 0; i < numOfAsts; i++ {
				if (((astPosArr[i].X - plyPos.X) < 40) && ((astPosArr[i].X - plyPos.X) > -40)) &&
					(((astPosArr[i].Y - plyPos.Y) < 40) && ((astPosArr[i].Y - plyPos.Y) > -40)) {
					time.Sleep(15000)
					win.Clear(colornames.Red)
					win.Update()
					time.Sleep(15000)
					gameOver = true
				}
			}

			for i := 0; i < numOfAsts; i++ {
				if (((astPosArr[i].X - misPos.X*2) < 40) && ((astPosArr[i].X - misPos.X*2) > -40)) &&
					(((astPosArr[i].Y - misPos.Y*2) < 40) && ((astPosArr[i].Y - misPos.Y*2) > -40)) {
					win.Clear(colornames.Yellow)
					points += 100
					ptsTxt.Clear()
					fmt.Fprintln(ptsTxt, "Punkty : ", points)
					win.Update()
					misLaunch = false
					misPos.X = 1000
					misPos.Y = 1000
					astPosArr[i].X = float64(random(20, 780))
					astPosArr[i].Y = float64(random(650, 800))
					astSpdArr[i] = float64(random(astMinSpd, astMaxSpd)) / 10
				}
			}
			if gunTimeout <= 0 {
				loadTxt.Clear()
				fmt.Fprintln(loadTxt, "Dzialo gotowe")
			} else {
				loadTxt.Clear()
				fmt.Fprintln(loadTxt, "Zaladunek : ", gunTimeout)
			}

			ptsTxt.Draw(win, pixel.IM)
			loadTxt.Draw(win, pixel.IM)
			win.Update()

			for i := 0; i < numOfAsts; i++ {
				if astPosArr[i].Y < -20 {
					astPosArr[i].X = float64(random(20, 780))
					astPosArr[i].Y = float64(random(650, 800))
					astSpdArr[i] = float64(random(astMinSpd, astMaxSpd)) / 10
				}
			}
			fpsCount++
			gunTimeout--

			if fpsCount == 1 {
				fmt.Println("Level: 1")
			}

			if fpsCount == 3600/2 { //3600
				fmt.Println("Level: 2")
				points += 2000
				ptsTxt.Clear()
				fmt.Fprintln(ptsTxt, "Punkty : ", points)
				currLvl++
			}

			if fpsCount == 7200/2 { //7200
				fmt.Println("Level: 3")
				points += 3000
				ptsTxt.Clear()
				fmt.Fprintln(ptsTxt, "Punkty : ", points)
				currLvl++
			}

			if fpsCount == 10800/2 { // 10800
				fmt.Println("Level: 4")
				points += 4000
				ptsTxt.Clear()
				fmt.Fprintln(ptsTxt, "Punkty : ", points)
				currLvl++
			}

			if fpsCount == 14400/2 { // 14400
				fmt.Println("Level: 5")
				points += 5000
				ptsTxt.Clear()
				fmt.Fprintln(ptsTxt, "Punkty : ", points)
				currLvl++
			}
		}

		win.Clear(colornames.Black)
		win.Update()

		gameOverAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
		gameOverText := text.New(pixel.V(280, 350), gameOverAtlas)
		gameOverText.Color = colornames.Yellow

		for {
			gameOverText.Clear()

			fmt.Fprintln(gameOverText, "KONIEC GRY!\nPUNKTY: ", points)
			endSpr.Draw(win, pixel.IM.Moved(win.Bounds().Center()))
			gameOverText.Draw(win, pixel.IM.Scaled(gameOverText.Orig, 3))
			win.Update()

			if win.Pressed(pixelgl.KeyEnter) {
				return
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pixelgl.Run(run)
}
