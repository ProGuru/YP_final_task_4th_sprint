// Пакет daysteps отвечает за учёт активности в течение дня. Он собирает переданную информацию в виде строк,
// парсит их и выводит информацию о количестве шагов, пройденной дистанции и потраченных калориях.
package daysteps

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах.
	stepLength = 0.65
	// Количество метров в одном километре.
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		log.Println("parsing problem")
		return 0, 0, errors.New("split string problem")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Println("number conversation problem")
		return 0, 0, fmt.Errorf("conversation to integer problem: %v", err)
	}
	if steps <= 0 {
		log.Println("sign or zero problem")
		return 0, 0, fmt.Errorf("negative or zero steps: %v", err)
	}

	strollDuration, err := time.ParseDuration(parts[1])
	if err != nil {
		log.Println("parsing problem")
		return 0, 0, fmt.Errorf("parsing time problem: %v", err)
	}

	if strollDuration <= 0 {
		log.Println("time problem")
		return 0, 0, fmt.Errorf("parsing time problem: %v", err)
	}
	return steps, strollDuration, nil
}

// DayActionInfo парсит строку с данными с помощью parsePackage(),
// вычисляет дистанцию в километрах, количество потраченных калорий и возвращает строку.
func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию
	steps, strollDuration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps <= 0 {
		return ""
	}

	distanceInM := float64(steps) * stepLength
	distanceInKm := distanceInM / mInKm
	burntCalories, _ := spentcalories.WalkingSpentCalories(steps, weight, height, strollDuration)
	return fmt.Sprintf(
		"Количество шагов: %d.\n"+
			"Дистанция составила %.2f км.\n"+
			"Вы сожгли %.2f ккал.\n",
		steps, distanceInKm, burntCalories)
}
