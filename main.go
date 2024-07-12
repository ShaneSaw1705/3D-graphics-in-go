package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(800, 450, "3D test scene")
	defer rl.CloseWindow()

	camera := rl.Camera{
		Position: rl.NewVector3(0, 10, 10),
		Target:   rl.NewVector3(0, 0, 0),
		Up:       rl.NewVector3(0, 1, 0),
		Fovy:     45.0,
	}
	rl.DisableCursor()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {

		rl.UpdateCamera(&camera, rl.CameraFree)

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		rl.DrawPlane(rl.NewVector3(0, 0, 0), rl.NewVector2(100, 100), rl.DarkGreen)

		rl.EndMode3D()

		rl.DrawText("3D test", 10, 10, 20, rl.LightGray)
		rl.EndDrawing()
	}
}
