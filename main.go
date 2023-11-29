package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var csvPath string
var limit int
var shuffle bool

type problem struct {
	question string
	answer   string
}

func readFile(path string) (*os.File, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error during file open: %q", err)
	}

	return file, nil
}

func initFlags() {
	flag.StringVar(&csvPath, "csv", "problems.csv", "a csv file in the format of 'question,answer' (default 'problems.cvs')")
	flag.IntVar(&limit, "limit", 30, "the time limit for the quiz in secods (default 30)")
	flag.BoolVar(&shuffle, "shuffle", false, "shuffles questions in the csv (default false)")
	flag.Parse()
}

func parseCsv(csvFile *os.File) ([][]string, error) {
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	data, err := reader.ReadAll()

	if err != nil {
		return nil, fmt.Errorf("error during reading csv content: %q", err)
	}
	defer csvFile.Close()

	return data, nil
}

func start(problems []problem) {
	var total int
	answer := ""

	if shuffle {
		rand.Shuffle(len(problems), func(i, j int) { problems[i], problems[j] = problems[j], problems[i] })
	}

	fmt.Println(problems)

	for i, problem := range problems {
		fmt.Printf("Problem %d#: %s = ", i, problem.question)
		_, err := fmt.Scanf("%s\n", &answer)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}

		if answer == problem.answer {
			total++
		}
	}
	fmt.Printf("Your total score: %v\n", total)
}

func main() {

	initFlags()

	file, err := readFile(csvPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	data, err := parseCsv(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	problems := make([]problem, len(data))
	for i, val := range data {
		problems[i] = problem{
			question: val[0],
			answer:   strings.TrimSpace(val[1]),
		}
	}

	start(problems)
}
