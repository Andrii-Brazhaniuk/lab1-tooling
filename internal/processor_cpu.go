package internal

import (
	"regexp"
)

// Повільна версія: компілює регулярку на кожній ітерації
func ProcessDataSlow(data []string) int {
	count := 0
	for _, s := range data {
		// ПОМИЛКА: Compile всередині циклу — це дуже дорого
		re := regexp.MustCompile(`^[a-z]+$`)
		if re.MatchString(s) {
			count++
		}
	}
	return count
}

// Оптимізована версія: компілює один раз
var reOptimized = regexp.MustCompile(`^[a-z]+$`)

func ProcessDataFast(data []string) int {
	count := 0
	for _, s := range data {
		if reOptimized.MatchString(s) {
			count++
		}
	}
	return count
}
