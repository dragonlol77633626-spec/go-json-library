package cli

import (
	"fmt"
)

func Run(lib interface{}, dbPath string) {
	for {
		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			fmt.Println("Сохранение данных и выход.")
			if err := storage.SaveLibraryToJSON(dbPath, lib); err != nil {
				fmt.Println("Произошла ошибка сохранения списка книг:", err)
			}
			return
		}
	}
}
