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
	duration := time.Duration(0)
	parseSlice := strings.Split(data, ",")
	if len(parseSlice) != 3 {
		return 0, "", duration, errors.New("")
	}
	steps, err := strconv.Atoi(parseSlice[0])
	if err != nil || steps <= 0 {
		return 0, "", duration, errors.New("количество шагов должно быть положительным")
	}
	activity := parseSlice[1]

	durat, err := time.ParseDuration(parseSlice[2])
	if err != nil || durat <= 0 {
		return 0, "", duration, errors.New("продолжительность должна быть положительной")
	}
	//actTime := strings.Split(strings.Split(stringDur, "m")[0], "h")
	//hours, err := strconv.Atoi(actTime[0])
	//if err != nil {
	//	return 0, "", duration, err
	//}
	//minutes, err := strconv.Atoi(actTime[1])
	//if err != nil {
	//	return 0, "", duration, err
	//}
	//duration = time.Duration(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute)
	return steps, activity, durat, nil
}

func distance(steps int, height float64) float64 {
	stepSize := height * stepLengthCoefficient
	pathOfSuffer := stepSize * float64(steps)
	return pathOfSuffer / float64(mInKm)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	wayOfSuffer := distance(steps, height)
	return wayOfSuffer / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	steps, activity, duration, err := parseTraining(data)
	if err != nil {
		log.Println(err)
	}
	pathOfSuffer := float64(steps) * lenStep / mInKm

	switch activity {
	case "Ходьба":
		walking, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		aver := pathOfSuffer / duration.Hours()
		result := fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\n", duration.Hours())
		result2 := fmt.Sprintf("Дистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", pathOfSuffer, aver, walking)
		return result + result2, nil
	case "Бег":
		running, err := WalkingSpentCalories(steps, weight, height, duration)
		if err != nil {
			return "", err
		}
		aver := pathOfSuffer / duration.Hours()
		result := fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\n", duration.Hours())
		result2 := fmt.Sprintf("Дистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", pathOfSuffer, aver, running)
		return result + result2, nil
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть полоительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	res := (weight * averageSpeed * minutes) / float64(minInH)
	return res, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть полоительным")
	}
	if weight <= 0 {
		return 0, errors.New("вес должен быть положительным")
	}
	if height <= 0 {
		return 0, errors.New("рост должен быть положительным")
	}
	if duration <= 0 {
		return 0, errors.New("продолжительность должна быть положительной")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	res := weight * averageSpeed * minutes * walkingCaloriesCoefficient / float64(minInH)
	return res, nil
}
