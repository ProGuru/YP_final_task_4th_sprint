// Пакет spentcalories обрабатывает переданную информацию и рассчитывает потраченные калории
// в зависимости от вида активности — бега или ходьбы. И тоже возвращает информацию обо всех тренировках.
package spentcalories

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return 0, "", 0, errors.New("split string problem")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, "", 0, fmt.Errorf("number conversation problem: %v", err)
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("negative or zero steps: %v", err)
	}

	activityDuration, err := time.ParseDuration(parts[2])
	if err != nil {
		return 0, "", 0, fmt.Errorf("parsing time problem: %v", err)
	}
	if activityDuration <= 0 {
		return 0, "", 0, fmt.Errorf("negative or zero duration of activity: %v", err)
	}

	return steps, parts[1], activityDuration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	stepLength := height * stepLengthCoefficient
	distanceInM := float64(steps) * stepLength
	distanceInKm := distanceInM / mInKm

	return distanceInKm
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
	if duration <= 0 {
		return 0
	}

	manDistance := distance(steps, height)
	return manDistance / duration.Hours()
}

// TrainingInfo выводит полный отчёт о тренировке.
func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	steps, active, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
		return "", fmt.Errorf("parsing problem: %v", err)
	}
	if steps <= 0 {
		return "", fmt.Errorf("negative or zero steps: %v", err)
	}
	if active == "" {
		return "", fmt.Errorf("empty activity name: %v", err)
	}
	if duration <= 0 {
		return "", fmt.Errorf("negative or zero duration of activity: %v", err)
	}

	var (
		trainingDistance float64
		speed            float64
		calories         float64
	)

	switch active {
	case "Бег":
		trainingDistance = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, _ = RunningSpentCalories(steps, weight, height, duration)
	case "Ходьба":
		trainingDistance = distance(steps, height)
		speed = meanSpeed(steps, height, duration)
		calories, _ = WalkingSpentCalories(steps, weight, height, duration)
	default:
		return "", errors.New("unknown training name")
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %.2f ч.\n"+
			"Дистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\n"+
			"Сожгли калорий: %.2f\n",
		active, duration.Hours(), trainingDistance, speed, calories), nil
}

// RunningSpentCalories - рассчитывает потраченные калории после бега.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, errors.New("one of input parameters less or equal 0")
	}

	speed := meanSpeed(steps, height, duration)
	durationInMin := duration.Minutes()
	return (weight * speed * durationInMin) / minInH, nil
}

// WalkingSpentCalories - рассчитывает потраченные калории после ходьбы.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	return calories * walkingCaloriesCoefficient, err
}
