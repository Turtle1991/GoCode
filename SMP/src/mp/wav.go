package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
}

func (w *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music", source)

	w.progress = 0

	for w.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		w.progress += 10
	}

	fmt.Println("\nFinished playing", source)
}
