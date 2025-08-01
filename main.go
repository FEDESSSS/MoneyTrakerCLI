package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/k0kubun/pp/v3"
)

type Expense struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
	Amount      string    `json:"amount"`
}

const fileJson = "moneyList.json"

var maxId int
var traker []Expense
var monthWaste []Expense

func main() {
	if len(os.Args) < 2 {
		CommandList()
		return
	}

	if err := runTreker(); err != nil {
		log.Fatal(err)
	}
}

func CommandList() {
	pp.Println("Список команд")
	pp.Println("=================================")
	pp.Println("list - просмотр всех трат")
	pp.Println("csv- сформировать таблицу")
	pp.Println("add (--description ?) (--amount ?)- добавление новой затраты")
	pp.Println("delete (--id ?) - удаление затраты с выбранным id")
	pp.Println("summary - просмотр общей суммы затрат")
	pp.Println("summary (--month ?)- просмотр общей суммы затрат за определенный месяц")
}

func runTreker() error {
	args := os.Args[1:]
	cmd := args[0]

	switch cmd {
	case "add":
		if len(args) > 5 {
			return errors.New("неправильная команда")
		}
		if args[1] == "--description" && args[3] == "--amount" {
			AddWaste(args[2], args[4])
		}
	case "list":
		if len(args) < 1 {
			return errors.New("неправильная команда")
		}
		ListWaste()

	case "csv":
		if len(args) < 1 {
			return errors.New("неправильная команда")
		}
		CsvFormated()

	case "summary":
		if len(args) < 1 {
			return errors.New("неправильная команда")
		}
		if len(args) == 3 && args[1] == "--month" {
			return SummaryWasteMonth(args[2])
		}
		SummaryWaste()
	case "delete":
		if len(args) < 2 {
			return errors.New("неправильная команда")
		}
		DeleteWaste(args[1])

	default:
		return errors.New("неправильная команда 12")
	}
	return nil
}

func JsonData() error {
	data, err := os.ReadFile(fileJson)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения файла: %v", err)
	}

	if len(data) > 0 {
		json.Unmarshal(data, &traker)
	}
	return nil
}

func AddWaste(description string, amount string) error {
	maxId++

	JsonData()

	maxId = 0
	for _, i := range traker {
		if i.Id > maxId {
			maxId = i.Id
		}
	}

	newWaste := Expense{
		Id:          maxId + 1,
		Date:        time.Now(),
		Description: description,
		Amount:      amount,
	}

	traker := append(traker, newWaste)

	jsonData, err := json.MarshalIndent(traker, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	if err := os.WriteFile(fileJson, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}

	fmt.Println("Добавлена затрата:")
	pp.Println("===============================")
	pp.Println(newWaste)
	return nil
}

func ListWaste() {
	JsonData()
	fmt.Println("Список затрат:")
	pp.Println("===============================")
	pp.Println(traker)
}

func SummaryWaste() {
	var Sum int
	JsonData()
	for _, i := range traker {
		number, _ := strconv.Atoi(i.Amount)
		Sum += number
	}
	pp.Println("===============================")
	fmt.Printf("Общая сумма затрат: %d$", Sum)
}

func DeleteWaste(id string) error {
	JsonData()
	number, _ := strconv.Atoi(id)
	for i, j := range traker {
		if j.Id == number {
			k := i + 1
			traker = append(traker[:i], traker[k:]...)
		}
	}

	jsonData, err := json.MarshalIndent(traker, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации: %v", err)
	}

	if err := os.WriteFile(fileJson, jsonData, 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %v", err)
	}
	pp.Println("===============================")
	fmt.Printf("Трата удалена")
	pp.Println(traker)
	return nil
}

func SummaryWasteMonth(month string) error {
	JsonData()
	var sum int
	number, _ := strconv.Atoi(month)
	for j, i := range traker {
		if int(i.Date.Month()) == number {
			monthWaste = append(monthWaste, traker[j])
			num, _ := strconv.Atoi(i.Amount)
			sum += num
		}
	}

	if len(monthWaste) == 0 {
		return fmt.Errorf("записей в выбранном месяце нет")
	}

	fmt.Println("===============================")
	fmt.Printf("Список затрат в %d месяце:\n", number)
	for _, expense := range monthWaste {
		pp.Println(expense)
	}
	fmt.Printf("Всего потрачено: %d$", sum)

	return nil
}

func CsvFormated() error {
	JsonData()
	file, err := os.Create("csvWaste.csv")
	if err != nil {
		return fmt.Errorf("не удалось создать файл")
	}

	defer file.Close()

	w := csv.NewWriter(file)

	header := []string{"Id", "Description", "Date", "Amount"}

	if err := w.Write(header); err != nil {
		return fmt.Errorf("ошибка")
	}

	var Waste [][]string
	for _, i := range traker {
		waste := []string{
			strconv.Itoa(i.Id),
			i.Date.Format("2006-01-01"),
			i.Description,
			i.Amount,
		}
		Waste = append(Waste, waste)
	}

	if err := w.WriteAll(Waste); err != nil {
		return fmt.Errorf("ошибка")
	}

	defer w.Flush()

	return nil
}
