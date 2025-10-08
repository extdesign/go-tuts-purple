package bins

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"time"

	"a.tasks.11/files"
	"a.tasks.11/output"
	"a.tasks.11/storage"
)

type Bin struct {
	ID        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func BinsExample() {
	List := []string{
		"name1",
		"name2",
		"name3",
	}

	BinList := GenerateBinList(&List)
	fmt.Println(BinList)
}

func GenerateBinList(list *[]string) []Bin {
	var res []Bin

	for idx, name := range *list {
		bin, err := newBin(name, false)

		bin.ID = fmt.Sprint(idx) + "-" + bin.ID

		if err != nil {
			output.PrintRed(fmt.Sprintf("\nНе удалось создать Bin с именем: '%s'. Ошибка: %s\n\n", name, err))
			continue
		}

		res = append(res, bin)
	}

	storage, err := storage.NewStorage(files.NewJsonDb("storage.json"))

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось инициализировать Storage. Ошибка: %s\n\n", err))
	}

	resJSON, err := json.MarshalIndent(res, "", "  ")

	if err != nil {
		output.PrintRed(fmt.Sprintf("\nНе удалось записать Json строку. Ошибка: %s\n\n", err))
	}

	storage.Data = resJSON
	storage.Save()

	return res
}

func newBin(name string, private bool) (Bin, error) {
	if len(name) == 0 {
		return Bin{}, errors.New("name is empty")
	}

	res := Bin{
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
	res.generateId()

	return res, nil
}

func (bin *Bin) generateId() {
	symbols := []rune("abcdefABCDEF")
	res := make([]rune, 10)
	for i := range 10 {
		res[i] = symbols[rand.IntN(len(symbols))]
	}
	bin.ID = string(res)
}
