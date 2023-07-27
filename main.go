package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 600
const HEIGHT = 400
const BOTTOM = HEIGHT - FONT_SIZE
const FONT_SIZE = 40

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

	rl.InitWindow(WIDTH, HEIGHT, "VCL")
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	pause := false

	m := rl.LoadMusicStream(filePath)
	rl.SetMusicVolume(m, vol)

	fName := strings.Split(filePath, "/")
	title := fmt.Sprintf("%s", fName[len(fName)-1])
	rl.SetWindowTitle(title)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.UpdateMusicStream(m)

		v := fmt.Sprintf("VOL: %.1f", vol)
		rl.DrawText(v, WIDTH-WIDTH+10, HEIGHT/100, FONT_SIZE/2, rl.LightGray)

		if !rl.IsMusicStreamPlaying(m) {
			rl.DrawCircle(WIDTH/2, HEIGHT/100+10, (WIDTH-HEIGHT)/20, rl.Red)
		}
		if rl.IsMusicStreamPlaying(m) {
			rl.DrawCircle(WIDTH/2, HEIGHT/100+10, (WIDTH-HEIGHT)/20, rl.Green)
		}
		if !rl.IsMusicReady(m) {
			rl.DrawText("FILE COULD NOT BE LOADED", WIDTH/4, HEIGHT/2, FONT_SIZE/2, rl.Red)
		}

		if rl.IsKeyPressed(rl.KeySpace) || rl.IsKeyPressed(rl.KeyP) {
			pause = !pause
			if !pause {
				rl.PauseMusicStream(m)
			} else {
				rl.PlayMusicStream(m)
			}
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
