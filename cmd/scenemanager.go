package main

type SceneManager struct {
    objects *[]*Planet
}

func (sm* SceneManager) Render() {
    for _, object := range *sm.objects {
        object.Render()
    }
}
