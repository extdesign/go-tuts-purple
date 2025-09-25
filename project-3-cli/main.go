package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func main() {
	List := []string{
		"name1",
		"name2",
		"name3",
	}

	BinList := generateBinList(&List)
	fmt.Println(BinList)
}

func generateBinList(list *[]string) []Bin {
	var res []Bin

	for idx, name := range *list {
		bin, err := newBin(name, false)

		bin.id = fmt.Sprint(idx) + "-" + bin.id

		if err != nil {
			fmt.Printf("Не удалось создать Bin с именем: '%s'. Ошибка: %s\n", name, err)
			continue
		}

		res = append(res, bin)
	}

	return res
}

func newBin(name string, private bool) (Bin, error) {
	if len(name) == 0 {
		return Bin{}, errors.New("name is empty")
	}

	res := Bin{
		private:   private,
		createdAt: time.Now(),
		name:      name,
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
	bin.id = string(res)
}
