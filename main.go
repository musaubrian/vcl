package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 600
const HEIGHT = 450
const BOTTOM = HEIGHT - FONT_SIZE
const FONT_SIZE = 40

//	var pos = rl.Vector2{
//		X: float32(WIDTH / 2),
//		Y: float32(HEIGHT / 2),
//	}
var size = rl.Vector2{
	X: 50,
	Y: 50,
}

var pos = rl.Vector2{
	X: float32(WIDTH / 2),
	Y: float32(HEIGHT / 2),
}
var speed = rl.Vector2{
	X: 3.0,
	Y: 3.0 / 2,
}

func main() {
	var vol float32 = 0.3

	var filePath string
	flag.StringVar(&filePath, "f", "", "PATH TO MUSIC FILE")
	flag.Parse()

	// If using build.sh
	if strings.Contains(filePath, "#") {
		filePath = strings.ReplaceAll(filePath, "#", "_")
	} else {
		if strings.Contains(filePath, "_") {
			filePath = strings.ReplaceAll(filePath, "_", " ")
		}
	}

	if len(filePath) < 1 {
		log.Fatal("YOU NEED TO PARSE IN PATH TO THE MUSIC FILE\n\n vcl -f ~/path/to/music")
	}
	// radius := 20

	rl.InitWindow(WIDTH, HEIGHT, "VCL")
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	pause := true

	m := rl.LoadMusicStream(filePath)
	rl.SetMusicVolume(m, vol)

	fName := strings.Split(filePath, "/")
	title := fmt.Sprintf("%s", fName[len(fName)-1])
	rl.SetWindowTitle(title)

	for !rl.WindowShouldClose() {
		v := fmt.Sprintf("VOL: %.1f", vol)
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.UpdateMusicStream(m)

		rec := rl.Rectangle{
			X:      pos.X,
			Y:      pos.Y,
			Width:  size.X,
			Height: size.Y,
		}

		rl.DrawText(v, 10, HEIGHT/100, FONT_SIZE/2, rl.LightGray)
		if !pause {
			pos.X += speed.X
			pos.Y += speed.Y

			if pos.X >= (float32(rl.GetRenderWidth())-float32(size.X)) || pos.X <= float32(size.X)/2 {
				speed.X *= -1
			}
			if pos.Y >= float32(rl.GetRenderHeight())-float32(size.Y) || pos.Y <= float32(FONT_SIZE) {
				speed.Y *= -1
			}
		}

		if rl.IsMusicStreamPlaying(m) {
			rl.DrawCircle(WIDTH/2, HEIGHT/100+10, (WIDTH-HEIGHT)/20, rl.Green)
			rl.DrawRectangleRounded(rec, 0.15, 2, rl.Blue)
		} else {
			rl.DrawCircle(WIDTH/2, HEIGHT/100+10, (WIDTH-HEIGHT)/20, rl.Red)
			rl.DrawRectangleRounded(rec, 0.15, 2, rl.Blue)
		}
		if !rl.IsMusicReady(m) {
			rl.DrawText("FILE COULD NOT BE LOADED", WIDTH/4, HEIGHT/2, FONT_SIZE/2, rl.Red)
		}

		if rl.IsKeyPressed(rl.KeySpace) {
			pause = !pause
		}

		if pause {
			rl.PauseMusicStream(m)
		} else {
			rl.PlayMusicStream(m)
		}
		if rl.IsKeyPressed(rl.KeyUp) {
			vol = vol + 0.1
			if vol <= 1.0 {
				rl.SetMusicVolume(m, vol)
			}
			if vol >= 1.0 {
				vol = 1.0
				rl.SetMusicVolume(m, vol)
			}
		}
		if rl.IsKeyPressed(rl.KeyDown) {
			vol = vol - 0.10
			if vol <= 0 {
				vol = 0.0
			}
			rl.SetMusicVolume(m, vol)
		}
		rl.EndDrawing()
	}
}
