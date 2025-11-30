package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, errors.New("ошибка в процессе деления строки")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при конвертации строки в число: %v", err)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("количество шагов меньше или равно 0: %v", err)
	}

	strollTime, err := time.Parse("15h04m", parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при преобразовании строки в длительность прогулки: %v", err)
	}
	strollDuration := time.Duration(strollTime.Hour() + strollTime.Minute())

	return steps, strollDuration, nil
}

// DayActionInfo должна парсить строку с данными с помощью parsePackage(),
// вычислять дистанцию в километрах и количество потраченных калорий и возвращать строку
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

	distanceInM := steps * stepLength
	distanceInKm := distanceInM / mInKm
	burntCalories := spentcalories.WalkingSpentCalories(steps, weight, height, strollDuration)
	return fmt.Sprintf("Количество шагов: %d./nДистанция составила %.2f км./nВы сожгли %.2f ккал.", steps, distanceInKm, burntCalories)
}
