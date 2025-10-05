package files

import (
	"fmt"
	"os"

	"a.tasks.11/output"
)

func ReadFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось прочитать данные из файла. Ошибка: %s\n\n", err.Error()))
		return nil, err
	}

	return data, nil
}

func WriteFile(content []byte, filename string) {
	file, err := os.Create(filename)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось создать файл. Ошибка: %s\n\n", err.Error()))
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось записать информаци в файл. Ошибка: %s\n\n", err.Error()))
		return
	}

	output.PrintYellow("\nДанные успешно записаны в файл\n\n")
}
