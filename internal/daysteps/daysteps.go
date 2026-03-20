package daysteps

import (
	"errors"
	"fmt"
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
	// переменная для возврата нулевого значения в случае ошибки
	duration := time.Duration(0)
	// получение слайс со стоковыми значениями
	parseList := strings.Split(data, ",")
	// проверка слайса на количество элементов
	if len(parseList) != 2 {
		return 0, duration, errors.New("неверный формат")
	}
	// получение количество шагов
	steps, err := strconv.Atoi(parseList[0])
	// проверка корректности Atoi, обработка ошибки и некорректного зисла шагов
	if err != nil || steps <= 0 {
		return 0, duration, errors.New("неверное количество шагов")
	}
	// вычисление переменной, которую вернет функция
	durat, err := time.ParseDuration(parseList[1])
	// обработка ошибки или некорректного значения
	if err != nil || durat <= 0 {
		return 0, duration, errors.New("неверная продолжительность")
	}
	return steps, durat, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// получение значений из строки при помощи parsePackage()
	steps, duration, err := parsePackage(data)
	// обработка возможной ошибки
	if err != nil {
		log.Println(err)
		return ""
	}
	// обработка возможной ошибки
	if steps <= 0 {
		log.Println(err)
		return ""
	}
	// вычисление дистанции в метрах
	distance := float64(steps) * stepLength
	// вычисление дистанции в километрах
	way := distance / float64(mInKm)
	// расчет потраченных калорий при помощи функции WalkingSpentCalories()
	calories, err := spentcalories.WalkingSpentCalories(steps, weight, height, duration)
	// обработка возможной ошибки
	if err != nil {
		return ""
	}
	// рендеринг возвратной строки
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, way, calories)
	return result
}
