package lab

import "fmt"

// ----------- Подсистема: AudioSystem -----------

type AudioSystem struct {
	isOn   bool
	volume int
}

func NewAudioSystem() *AudioSystem {
	return &AudioSystem{isOn: false, volume: 5}
}

func (a *AudioSystem) TurnOn() {
	if !a.isOn {
		a.isOn = true
		fmt.Println("Audio system is turned on.")
	}
}

func (a *AudioSystem) SetVolume(level int) {
	if !a.isOn {
		fmt.Println("Audio system is off, cannot set volume.")
		return
	}
	if level < 0 {
		level = 0
	}
	if level > 10 {
		level = 10
	}
	a.volume = level
	fmt.Printf("Audio volume is set to %d.\n", level)
}

func (a *AudioSystem) TurnOff() {
	if a.isOn {
		a.isOn = false
		fmt.Println("Audio system is turned off.")
	}
}

// ----------- Подсистема: VideoProjector -----------

type VideoProjector struct {
	isOn       bool
	resolution string
}

func NewVideoProjector() *VideoProjector {
	return &VideoProjector{isOn: false, resolution: "HD"}
}

func (v *VideoProjector) TurnOn() {
	if !v.isOn {
		v.isOn = true
		fmt.Println("Video projector is turned on.")
	}
}

func (v *VideoProjector) SetResolution(resolution string) {
	if !v.isOn {
		fmt.Println("Video projector is off, cannot set resolution.")
		return
	}
	v.resolution = resolution
	fmt.Printf("Video resolution is set to %s.\n", resolution)
}

func (v *VideoProjector) TurnOff() {
	if v.isOn {
		v.isOn = false
		fmt.Println("Video projector is turned off.")
	}
}

// ----------- Подсистема: LightingSystem -----------

type LightingSystem struct {
	isOn       bool
	brightness int
}

func NewLightingSystem() *LightingSystem {
	return &LightingSystem{isOn: false, brightness: 10}
}

func (l *LightingSystem) TurnOn() {
	if !l.isOn {
		l.isOn = true
		fmt.Println("Lights are turned on.")
	}
}

func (l *LightingSystem) SetBrightness(level int) {
	if !l.isOn {
		fmt.Println("Lights are off, cannot set brightness.")
		return
	}
	if level < 0 {
		level = 0
	}
	if level > 10 {
		level = 10
	}
	l.brightness = level
	fmt.Printf("Lights brightness is set to %d.\n", level)
}

func (l *LightingSystem) TurnOff() {
	if l.isOn {
		l.isOn = false
		fmt.Println("Lights are turned off.")
	}
}

// ----------- Фасад: HomeTheaterFacade -----------

type HomeTheaterFacade struct {
	audio   *AudioSystem
	video   *VideoProjector
	lighting *LightingSystem
}

// Конструктор фасада: создаёт и собирает все подсистемы.
func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		audio:   NewAudioSystem(),
		video:   NewVideoProjector(),
		lighting: NewLightingSystem(),
	}
}

// StartMovie – сценарий запуска фильма.
func (h *HomeTheaterFacade) StartMovie() {
	fmt.Println("Preparing to start the movie...")
	h.lighting.TurnOn()
	h.lighting.SetBrightness(5)
	h.audio.TurnOn()
	h.audio.SetVolume(8)
	h.video.TurnOn()
	h.video.SetResolution("HD")
	fmt.Println("Movie started.")
}

// EndMovie – сценарий завершения фильма.
func (h *HomeTheaterFacade) EndMovie() {
	fmt.Println("Shutting down movie...")
	h.video.TurnOff()
	h.audio.TurnOff()
	h.lighting.TurnOff()
	fmt.Println("Movie ended.")
}
