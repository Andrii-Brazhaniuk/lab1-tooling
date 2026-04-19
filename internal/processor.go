package internal

import (
	"time"
)

// Глобальна змінна, яка "з'їдає" пам'ять
var metadataCache []string

func StartProcessing() {
	// Імітація обробки запитів
	go func() {
		// Зменшуємо інтервал до 10мс і додаємо важкі дані
		ticker := time.NewTicker(10 * time.Millisecond)
		for range ticker.C {
			// Створюємо довгий рядок (приблизно 100 КБ за раз)
			data := make([]byte, 102400)
			metadataCache = append(metadataCache, string(data))
		}
	}()
}
