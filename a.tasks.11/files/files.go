package files

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"a.tasks.11/output"
)

func ReadFile(filename string) ([]byte, error) {
	if !isFileJson(filename) {
		output.PrintRed("\nФайл должен быть с расширением json\n\n")
		return nil, errors.New("файл должен быть с расширением json")
	}

	data, err := os.ReadFile(filename)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось прочитать данные из файла. Ошибка: %s\n\n", err.Error()))
		return nil, err
	}

	return data, nil
}

func WriteFile(content []byte, filename string) error {
	if !isFileJson(filename) {
		output.PrintRed("\nФайл должен быть с расширением json\n\n")
		return errors.New("файл должен быть с расширением json")
	}

	file, err := os.Create(filename)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось создать файл. Ошибка: %s\n\n", err.Error()))
		return err
	}

	defer file.Close()

	_, err = file.Write(content)

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось записать информаци в файл. Ошибка: %s\n\n", err.Error()))
		return err
	}

	output.PrintYellow("\nДанные успешно записаны в файл\n\n")

	return nil
}

func isFileJson(fp string) bool {
	return filepath.Ext(fp) != "json"
}
