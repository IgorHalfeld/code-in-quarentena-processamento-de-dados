package main

import (
	"code-in-quarentena-golang-processamento-dados/database"
	"code-in-quarentena-golang-processamento-dados/domain"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup
var db *database.DatabaseImpl = database.New(database.Config{})

func extract(file *os.File) {
	reader := csv.NewReader(file)

	input := make(chan []string)
	output := make(chan domain.Issue)

	go transform(input, output)
	go load(output)
	wg.Add(2)

	index := 0
	for {
		row, err := reader.Read()
		if err == io.EOF {
			wg.Done()
			wg.Done()
			break
		}

		if index != 0 {
			input <- row
		}
		index++
		fmt.Println("Processing index:", index)
	}
}

func transform(input chan []string, output chan domain.Issue) {
	for {
		select {
		case r, ok := <-input:
			if !ok {
				close(input)
				close(output)
				wg.Done()
				break
			}

			wg.Add(1)
			go func(row []string) {
				url := row[0]

				cleanString := strings.ReplaceAll(url, "\"", "")
				ss := strings.Split(cleanString, "/")

				owner := ss[3]
				repo := ss[4]
				id := ss[6]

				issue := domain.Issue{
					ID:    id,
					Owner: owner,
					Repo:  repo,
					Url:   url,
					Body:  row[1],
				}
				output <- issue
				wg.Done()
			}(r)

		default:
		}
	}
}

func load(ch chan domain.Issue) {
	for {
		select {
		case i, ok := <-ch:
			if !ok {
				wg.Done()
				break
			}

			wg.Add(1)
			go func(issue domain.Issue) {
				_, err := db.Create("issues", issue)
				if err != nil {
					log.Println("Could not insert issue", err)
				}
				wg.Done()
			}(i)
		default:
		}
	}
}

func main() {
	file, err := os.Open("./github_issues.csv")
	if err != nil {
		log.Fatalln(err)
	}

	extract(file)
	defer file.Close()
}
