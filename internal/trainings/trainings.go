package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentenergy"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	if t.Steps <= 0 {
		return errors.New("шагов должно быть больше нуля")
	}
	if t.Duration <= 0 {
		return errors.New("продолжительность должна быть положительной")
	}
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("некорректный формат данных: ожидается 3 части")
	}
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	t.Steps = steps

	t.TrainingType = strings.TrimSpace(parts[1])

	durationStr := strings.TrimSpace(parts[2])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	tType := strings.ToLower(strings.TrimSpace(t.TrainingType))

	var distanceKm, speedKmh, calories float64
	var err error

	distanceKm = spentenergy.Distance(t.Steps, t.Height)

	speedKmh = spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch tType {
	case "бег", "run":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Height, t.Weight, t.Duration)
		if err != nil {
			calories = 0
		}

	case "ходьба", "walk":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Height, t.Weight, t.Duration)
		if err != nil {
			calories = 0
		}
	default:
		return "", errors.New("неизвестный тип тренировки")
	}
	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		tType,
		t.Duration.Hours(),
		distanceKm,
		speedKmh,
		calories,
	)

	return result, nil
}
