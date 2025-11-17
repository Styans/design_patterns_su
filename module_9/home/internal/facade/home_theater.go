package facade

import "fmt"

// --- Подсистемы ---

type TV struct {
	channel int
}

func (t *TV) On()  { fmt.Println("TV: on") }
func (t *TV) Off() { fmt.Println("TV: off") }

func (t *TV) SetChannel(ch int) {
	t.channel = ch
	fmt.Printf("TV: channel set to %d\n", ch)
}

type AudioSystem struct {
	volume int
}

func (a *AudioSystem) On()  { fmt.Println("Audio: on") }
func (a *AudioSystem) Off() { fmt.Println("Audio: off") }

func (a *AudioSystem) SetVolume(v int) {
	if v < 0 {
		v = 0
	}
	if v > 100 {
		v = 100
	}
	a.volume = v
	fmt.Printf("Audio: volume set to %d\n", v)
}

type DVDPlayer struct {
	movie string
}

func (d *DVDPlayer) Play(movie string) {
	d.movie = movie
	fmt.Printf("DVD: playing %q\n", movie)
}

func (d *DVDPlayer) Pause() {
	if d.movie == "" {
		return
	}
	fmt.Printf("DVD: paused %q\n", d.movie)
}

func (d *DVDPlayer) Stop() {
	if d.movie == "" {
		return
	}
	fmt.Printf("DVD: stopped %q\n", d.movie)
	d.movie = ""
}

type GameConsole struct {
	game string
}

func (g *GameConsole) On() {
	fmt.Println("Console: on")
}

func (g *GameConsole) StartGame(game string) {
	g.game = game
	fmt.Printf("Console: starting game %q\n", game)
}

// --- Фасад ---

type HomeTheaterFacade struct {
	tv      *TV
	audio   *AudioSystem
	dvd     *DVDPlayer
	console *GameConsole
}

func NewHomeTheaterFacade() *HomeTheaterFacade {
	return &HomeTheaterFacade{
		tv:      &TV{},
		audio:   &AudioSystem{},
		dvd:     &DVDPlayer{},
		console: &GameConsole{},
	}
}

// Включить систему и начать фильм
func (h *HomeTheaterFacade) WatchMovie(movie string) {
	fmt.Println("=== Scenario: Watch movie ===")
	h.tv.On()
	h.tv.SetChannel(1)
	h.audio.On()
	h.audio.SetVolume(50)
	h.dvd.Play(movie)
}

// Остановить фильм и выключить всё
func (h *HomeTheaterFacade) StopMovie() {
	fmt.Println("=== Scenario: Stop movie ===")
	h.dvd.Stop()
	h.audio.Off()
	h.tv.Off()
}

// Запустить игровую консоль
func (h *HomeTheaterFacade) PlayGame(game string) {
	fmt.Println("=== Scenario: Play game ===")
	h.tv.On()
	h.tv.SetChannel(2)
	h.audio.On()
	h.audio.SetVolume(60)
	h.console.On()
	h.console.StartGame(game)
}

// Сценарий: слушать музыку (TV как источник, звук через аудио)
func (h *HomeTheaterFacade) ListenMusic() {
	fmt.Println("=== Scenario: Listen music ===")
	h.tv.On()
	h.tv.SetChannel(3)
	h.audio.On()
	h.audio.SetVolume(40)
	fmt.Println("Music: playing via TV + AudioSystem")
}

// Регулировка громкости через фасад
func (h *HomeTheaterFacade) SetVolume(level int) {
	h.audio.SetVolume(level)
}
