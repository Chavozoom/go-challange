package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Person struct {
	Name  string
	Age   int
	Score int
}

func readCSV(filePath string) ([]Person, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var people []Person
	for _, record := range records[1:] { // Skip header
		age, _ := strconv.Atoi(record[1])
		score, _ := strconv.Atoi(record[2])
		person := Person{
			Name:  record[0],
			Age:   age,
			Score: score,
		}
		people = append(people, person)
	}

	return people, nil
}

func writeCSV(filePath string, people []Person) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Nome", "Idade", "Pontuação"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write data
	for _, person := range people {
		row := []string{person.Name, strconv.Itoa(person.Age), strconv.Itoa(person.Score)}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func sortByAge(people []Person) {
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
}

func sortByName(people []Person) {
	sort.Slice(people, func(i, j int) bool {
		return strings.Compare(strings.ToLower(people[i].Name), strings.ToLower(people[j].Name)) < 0
	})
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: go run main.go <arquivo-origem.csv> <arquivo-destino.csv>")
		os.Exit(1)
	}

	inputFilePath := os.Args[1]
	outputFilePath := os.Args[2]

	people, err := readCSV(inputFilePath)
	if err != nil {
		fmt.Printf("Erro ao ler o arquivo CSV: %v\n", err)
		os.Exit(1)
	}

	// Ordena por nome
	sortByName(people)
	if err := writeCSV(outputFilePath+"_ordenado_por_nome.csv", people); err != nil {
		fmt.Printf("Erro ao escrever o arquivo CSV ordenado por nome: %v\n", err)
		os.Exit(1)
	}

	// Ordena por idade
	sortByAge(people)
	if err := writeCSV(outputFilePath+"_ordenado_por_idade.csv", people); err != nil {
		fmt.Printf("Erro ao escrever o arquivo CSV ordenado por idade: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Processamento concluído.")
}
