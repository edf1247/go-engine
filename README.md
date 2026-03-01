## Solar System Simulation

## Usage

1. Make sure you have Go installed
2. Run `go run ./cmd`
3. To build, run `go build -o sim ./cmd`

### Brief Background

Originally, I was planning on making a game engine in go using raylib, but I realized that would involve implementing a lot of features that I just did not find interesting(mainly lighting). I decided to pivot to creating a physics simulation of some kind, and eventually settled on simulating the solar system after I realized that I didn't have an understanding of how planets actually orbited. Of course, I'd seen the universal gravitation equation in physics classes, but I didn't understand how that created actual orbits. Hence this project.

### Implementation

The project is a bit messy, so I'll probably do some work refactoring and creating better class structures. Current class structures:
- Body:
    - These are the Celestial Bodies
    - Mesh:
        - stores the mesh details for the spheres
    - RigidBody:
        - stores information relevant to the physics simulation
- SceneManager:
    - Acts as an interface for the objects in the scene
    - sceneObjects:
        - An array storing pointers to the Bodies

UniversalGravitation:
- This function iterates over every object in the SceneManager.sceneObjects array
- For each object, iterates over the rest of the objects:
    - Calculates the net force due to gravity on the current object, then calculates acceleration due to that force
    - Updating object positions happens after all gravitational calculations are performed

## Credits

This implementation is inspired by Sebastion Lague's Solar System [video](https://www.youtube.com/watch?v=7axImc1sxa0&t=50s&pp=ygUcc2ViYXN0aWFuIGxhZ3VlIHNvbGFyIHN5c3RlbQ%3D%3D). The main differences are that I'm working in Go, without the abstractions of the Unity API. 
