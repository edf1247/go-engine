package camera

import (
    rl "github.com/gen2brain/raylib-go/raylib"
)

const (
    SPEED = 5.0
    KEY_W = 87
    KEY_S = 83
    KEY_A = 65
    KEY_D = 68
)

func InitCamera() rl.Camera3D{
    cameraPos := rl.Vector3{0, 10, 150}
    cameraTarget := rl.Vector3{0, 0, 0}
    rotation := rl.Vector3{0, 1, 0}
    fov := float32(90.0)
    camera := rl.NewCamera3D(cameraPos, cameraTarget, rotation, fov, rl.CameraPerspective)
    return camera
}

func MoveCamera(mainCam *rl.Camera3D) {
    delta := SPEED * rl.GetFrameTime()

    if (rl.IsKeyDown(KEY_W)) {
        rl.CameraMoveForward(mainCam, delta, 0)
    }
    if (rl.IsKeyDown(KEY_S)) {
        rl.CameraMoveForward(mainCam, -delta, 0)
    }
    if (rl.IsKeyDown(KEY_A)) {
        rl.CameraMoveRight(mainCam, -delta, 0)
    }
    if (rl.IsKeyDown(KEY_D)) {
        rl.CameraMoveRight(mainCam, delta, 0)
    }
}
