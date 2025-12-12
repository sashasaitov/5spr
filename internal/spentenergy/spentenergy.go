package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
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

	meanSpeed := MeanSpeed(steps, height, duration)
	mins := duration.Minutes()
	calories := (weight * meanSpeed * mins) / minInH
	caloricOutput := calories * walkingCaloriesCoefficient

	return caloricOutput, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
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

	meanSpeed := MeanSpeed(steps, height, duration)
	mins := duration.Minutes()
	calorie := (weight * meanSpeed * mins) / minInH

	return calorie, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	dist := Distance(steps, height)
	hours := duration.Hours()
	return dist / hours
}

func Distance(steps int, height float64) float64 {
	stepLen := height * stepLengthCoefficient
	distM := float64(steps) * stepLen
	return distM / float64(mInKm)
}
