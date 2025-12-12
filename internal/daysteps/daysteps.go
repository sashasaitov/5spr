package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("некорректный формат данных: ожидается 2 части")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}
	ds.Steps = steps
	if steps <= 0 {
		return errors.New("шагов должно быть больше нуля")
	}

	durationStr := strings.TrimSpace(parts[1])
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return err
	}
	ds.Duration = duration
	if duration <= 0 {
		return errors.New("продолжительность должна быть положительной")

	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Height, ds.Weight, ds.Duration)
	if err != nil {
		return "", err
	}

	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	)

	return info, nil
}
