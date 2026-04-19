package internal
import (
	"time"
	"github.com/rs/zerolog/log" 
)
// Глобальна змінна, яка "з'їдає" пам'ять 
var metadataCache []string
func StartProcessing() {
	// Імітація обробки запитів
	go func() {
		ticker := time.NewTicker(10 * time.Millisecond)
		defer ticker.Stop() 
		for range ticker.C {
			// Створюємо важкі дані 
			data := make([]byte, 102400)
			metadataCache = append(metadataCache, string(data))

			// ЕТАП 2: Структуроване логування
			// Логуємо кожні 100 оброблених елементів
			if len(metadataCache)%100 == 0 {
				log.Info().
					Int("total_processed", len(metadataCache)).
					Str("status", "active").
					Msg("Image metadata processing heart-beat")
			}
		}
	}()
}