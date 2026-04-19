package internal

import (
	"sync"
	"time"
	"github.com/rs/zerolog/log"
)

var (
	lockA sync.Mutex
	lockB sync.Mutex
)

func CauseDeadlock() {
	// Горутина 1: захоплює A, потім чекає на B
	go func() {
		lockA.Lock()
		log.Info().Msg("Горутина 1: захопила A, чекає на B...")
		time.Sleep(1 * time.Second) // Даємо час другій горутині почати
		lockB.Lock() 
		
		log.Info().Msg("Горутина 1: виконала роботу")
		lockB.Unlock()
		lockA.Unlock()
	}()

	// Горутина 2: захоплює B, потім чекає на A
	go func() {
		lockB.Lock()
		log.Info().Msg("Горутина 2: захопила B, чекає на A...")
		time.Sleep(1 * time.Second)
		lockA.Lock() 
		
		log.Info().Msg("Горутина 2: виконала роботу")
		lockA.Unlock()
		lockB.Unlock()
	}()
}