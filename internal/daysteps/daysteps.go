package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	duration := time.Duration(0)
	parseList := strings.Split(data, ",")
	if len(parseList) != 2 {
		return 0, duration, errors.New("")
	}
	steps, err := strconv.Atoi(parseList[0])
	if err != nil || steps <= 0 {
		return 0, duration, err
	}
	walkTime := strings.Split(strings.Split(parseList[1], "m")[0], "h")
	hours, err := strconv.Atoi(walkTime[0])
	if err != nil {
		return 0, duration, err
	}
	minutes, err := strconv.Atoi(walkTime[1])
	if err != nil {
		return 0, duration, err
	}
	duration = time.Duration(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute)
	return steps, duration, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if steps == 0 {
		return ""
	}
	distance := float64(steps) * stepLength
	way := distance / float64(mInKm)
	calories, err := WalkingSpentCalories(steps, weight, height, duration)
	//// форматировать строку ответа
	return ""
}
