package spentcalories

import (
	"errors"
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
	switch activity {
		case "Ходьба":
			
		case "Бег":

		default:
			return "", errors.New("неизвестный вид тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps == 0 || weight == 0 || height == 0 || duration == 0 {
		return 0, errors.New("")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	res := (weight * averageSpeed * minutes) / float64(minInH)
	return res, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps == 0 || weight == 0 || height == 0 || duration == 0 {
		return 0, errors.New("")
	}
	averageSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	res := weight * averageSpeed * minutes * walkingCaloriesCoefficient / float64(minInH)
	return res, nil
}
