package main

import (
	"github.com/Andrii-Brazhaniuk/lab1-tooling/internal"
	"github.com/rs/zerolog/log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Info().Msg("Діагностичний сервер pprof запущено на http://localhost:6060")
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			log.Error().Err(err).Msg("pprof сервер не зміг запуститися")
		}
	}()

	// Дані для обробки
	testData := []string{"apple", "123", "banana", "Cherry", "go", "golang", "100", "test"}

	// Запускаємо нескінченний цикл з "повільною" функцією в горутині
	go func() {
		log.Info().Msg("Запущено навантаження на CPU (Slow Mode)...")
		for {
			internal.ProcessDataSlow(testData)
		}
	}()

	select {}
}
