package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for i, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("Ошибка парсинга строки %d: %v", i+1, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("Ошибка получения информации для строки %d: %v", i+1, err)
			continue
		}
		fmt.Println(info)
		if i < len(dataset)-1 {
			fmt.Println()
		}
	}
}
