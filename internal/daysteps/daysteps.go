package daysteps

import (
	"errors"
	"fmt"

	//"internal/spentcalories"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/DenisKhamzin/Yandex-Practicum-go1fl-4-sprint-final/internal/spentcalories"
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
		return 0, duration, errors.New("неверный формат")
	}
	steps, err := strconv.Atoi(parseList[0])
	if err != nil || steps <= 0 {
		return 0, duration, errors.New("неверное количество шагов")
	}
	//walkTime := strings.Split(strings.Split(parseList[1], "m")[0], "h")
	//hours, err := strconv.Atoi(walkTime[0])
	//if err != nil {
	//	return 0, duration, err
	//}
	//minutes, err := strconv.Atoi(walkTime[1])
	//if err != nil {
	//	return 0, duration, err
	//}
	//duration = time.Duration(time.Duration(hours)*time.Hour + time.Duration(minutes)*time.Minute)
	durat, err := time.ParseDuration(parseList[1])
	if err != nil || durat <= 0 {
		return 0, duration, errors.New("неверная продолжительность")
	}
	return steps, durat, nil
}

func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}
	if steps <= 0 {
		log.Println(err)
		return ""
	}
	distance := float64(steps) * stepLength
	way := distance / float64(mInKm)
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	if err != nil {
		return ""
	}
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, way, calories)
	return result
}
