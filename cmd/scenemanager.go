package main

type SceneManager struct {
    objects *[]*Body
}

func (sm* SceneManager) Render() {
    for _, object := range *sm.objects {
        object.Render()
    }
}
