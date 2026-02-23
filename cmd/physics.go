package main

import (
    rl "github.com/gen2brain/raylib-go/raylib"
    "math"
)

type Planet struct {
    centerPos rl.Vector3
    radius float32
    rings int32
    slices int32
    mass float32
    color rl.Color
    wireColor rl.Color
    accelerationVector rl.Vector3
}

const (
    G = math.Pi * math.Pi * 4
    SIM_SPEED = 100
)

func (planet* Planet) Render() {
    rl.DrawSphereEx(planet.centerPos, planet.radius, planet.rings, planet.slices, planet.color)
    rl.DrawSphereWires(planet.centerPos, planet.radius, planet.rings, planet.slices, planet.wireColor)
}

func UniversalGravitation(objects *[]*Planet) {
    for i := 0; i < len(*objects); i++ {
        netForce := rl.Vector3{0, 0, 0}
        m1 := (*objects)[i]
        for j := 0; j < len(*objects); j++ {
            if i == j {
                continue
            }
            
            m2 := (*objects)[j]
            
            dx := float64(m2.centerPos.X - m1.centerPos.X)
            dz := float64(m2.centerPos.Z - m1.centerPos.Z)
            
            distSq := (dx * dx) + (dz * dz)
            dist := math.Sqrt((dx * dx) + (dz * dz))
            
            force := float64(G * m1.mass * m2.mass) / distSq
            
            netForce.X += float32((force * dx) / dist)
            netForce.Z += float32((force * dz) / dist)

        }
        m1.accelerationVector.X += netForce.X / m1.mass
        m1.accelerationVector.Z += netForce.Z / m1.mass
        

        m1.centerPos.X += SIM_SPEED * m1.accelerationVector.X * rl.GetFrameTime()
        m1.centerPos.Z += SIM_SPEED * m1.accelerationVector.Z * rl.GetFrameTime()

    }
}
