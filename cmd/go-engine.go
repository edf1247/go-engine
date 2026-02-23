package main

import (
    camera "github.com/edf1247/go-engine/internal/camera"
    rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 800
const HEIGHT = 450

var (
    centerPos = rl.Vector3{0, 0, 0}
    initialAcceleration = rl.Vector3{0, 0, 0}
)

func main() {
  rl.InitWindow(WIDTH, HEIGHT, "Go Engine")
  defer rl.CloseWindow()

  var start = false
  var mainCamera rl.Camera3D
  var sm SceneManager
  var sceneObjects []*Planet
  var counter = 0

  rl.SetTargetFPS(60.0)


  for !rl.WindowShouldClose() {
    rl.BeginDrawing()
    rl.ClearBackground(rl.Black)
    rl.DrawFPS(10.0, 10.0)

    if !start {
        mainCamera = camera.InitCamera()

        radius := float32(100.0)
        rings := int32(15)
        slices := int32(15)
        mass := float32(1.0)
        
        sun := Planet{centerPos, radius, rings, slices, mass, rl.White, rl.Maroon, initialAcceleration}
        earth := Planet{rl.Vector3{100.0, 0.0, 100.0}, float32(.9157), rings, slices, float32(0.000003), rl.Red, rl.White, initialAcceleration}
        
        sceneObjects = append(sceneObjects, &sun)
        sceneObjects = append(sceneObjects, &earth)

        sm.objects = &sceneObjects

        start = true
    }
    rl.UpdateCamera(&mainCamera, rl.CameraFree)
    
    rl.BeginMode3D(mainCamera)
    
    //camera.MoveCamera(&mainCamera)
    
    sm.Render()
    
    if (counter % 10 == 0) {
        UniversalGravitation(sm.objects)
    }

    rl.DrawGrid(10.0, 1.0)
    rl.EndMode3D()

    rl.EndDrawing()
    counter += 1
  }
}
