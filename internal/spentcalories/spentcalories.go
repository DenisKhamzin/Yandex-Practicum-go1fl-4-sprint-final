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
	// нулевое значение для возврата в случае ошибок
	duration := time.Duration(0)
	// получение слайса со строками из начальной строки
	parseSlice := strings.Split(data, ",")
	// проверка слайса на количество элементов
	if len(parseSlice) != 3 {
		return 0, "", duration, errors.New("")
	}
	// получение количества шагов при помощи Atoi
	steps, err := strconv.Atoi(parseSlice[0])
	// обработка возможной ошибки Atoi и корректного количества шагов
	if err != nil || steps <= 0 {
		return 0, "", duration, errors.New("количество шагов должно быть положительным")
	}
	// определение для вида активности отдельной переменной
	activity := parseSlice[1]
	// получение значения для возврата функцией
	durat, err := time.ParseDuration(parseSlice[2])
	// обработка ошибки или некорректного значения
	if err != nil || durat <= 0 {
		return 0, "", duration, errors.New("продолжительность должна быть положительной")
	}
	return steps, activity, durat, nil
}

func distance(steps int, height float64) float64 {
	// вычисление длины шага
	stepSize := height * stepLengthCoefficient
	// вычисление длины пути
	pathOfSuffer := stepSize * float64(steps)
	return pathOfSuffer / float64(mInKm)
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// проверка входящено значение на положительность
	if duration <= 0 {
		return 0
	}
	// расчет длины пути
	wayOfSuffer := distance(steps, height)
	// расчет и возврат средней скорочти
	return wayOfSuffer / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// получение значений из строки при помощи parseTraining
	steps, activity, duration, err := parseTraining(data)
	// логирование возможной ошибки
	if err != nil {
		log.Println(err)
	}
	// получение дистанции в километрах
	pathOfSuffer := float64(steps) * height * float64(stepLengthCoefficient) / mInKm
	// расчет калорий при разных вариантах активности
	switch activity {
	case "Ходьба":
		// получение количества калорий при помощи WalkingSpentCalories()
		walking, err := WalkingSpentCalories(steps, weight, height, duration)
		// обработка возможной ошибки
		if err != nil {
			return "", err
		}
		// расчет скорости
		aver := pathOfSuffer / duration.Hours()
		// склеивание возвратной строки из двух значений для лучшей читаемости
		result := fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f ч.\n", duration.Hours())
		result2 := fmt.Sprintf("Дистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", pathOfSuffer, aver, walking)
		return result + result2, nil
	case "Бег":
		// получение количества калорий при помощи RunnungSpentCalories()
		running, err := RunningSpentCalories(steps, weight, height, duration)
		// обработка возможной ошибки
		if err != nil {
			return "", err
		}
		// расчет скорости
		aver := pathOfSuffer / duration.Hours()
		// склеивание возвратной строки их двух значений для лучшей читаемости
		result := fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f ч.\n", duration.Hours())
		result2 := fmt.Sprintf("Дистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", pathOfSuffer, aver, running)
		return result + result2, nil
	default:
		// возврат ошибки в случае неизвестного типа тренировок
		return "", errors.New("неизвестный тип тренировки")
	}
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверка на положительность всех аргументов
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
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
	// вычисление и возврат количества потраченных при беге калорий
	averageSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	res := (weight * averageSpeed * minutes) / float64(minInH)
	return res, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// проверка на положительность всех аргументов
	if steps <= 0 {
		return 0, errors.New("количество шагов должно быть положительным")
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
	// вычисление и возврат количества потраченных при ходьбе
	averageSpeed := meanSpeed(steps, height, duration)
	minutes := duration.Minutes()
	res := weight * averageSpeed * minutes * walkingCaloriesCoefficient / float64(minInH)
	return res, nil
}
