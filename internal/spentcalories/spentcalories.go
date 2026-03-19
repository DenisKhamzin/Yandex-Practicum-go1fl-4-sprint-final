package spentcalories

import (
	"strconv"
	"time"
	"strings"
	"errors"
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
	duration := time.Duration(0)
	parseSlice := strings.Split(data, ",")
	if len(parseSlice) != 3 {
		return 0, "", duration, errors.New("")
	}
	steps, err := strconv.Atoi(parseSlice[0])
	if err != nil {
			return 0, "", duration, err
	}
	activity := parseSlice[1]
	stringDur := parseSlice[2]
	actTime := strings.Split(strings.Split(stringDur, "m")[0], "h")
	hours, err := strconv.Atoi(actTime[0])
	if err != nil {
		return 0, "", duration, err
	}
	minutes, err := strconv.Atoi(actTime[1])
	if err != nil {
		return 0, "", duration, err
	}
	duration = time.Duration(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute)
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
}
