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
)

var db *database.DatabaseImpl = database.New(database.Config{})

func extract(file *os.File) {
	reader := csv.NewReader(file)
	reader.ReuseRecord = true

	index := 1
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}

		if index != 1 {
			transform(row)
		}

		index++

		if index == 1000 {
			break
		}
	}
}

func transform(row []string) {

	url := row[0]

	ss := strings.Split(strings.ReplaceAll(url, "\"", ""), "/")

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

	load(issue)
}

func load(issue domain.Issue) {
	_, err := db.Create("issues", issue)
	if err != nil {
		log.Println("Could not insert issue", err)
	}

	fmt.Println("Got issue", issue.ID)
}

func main() {
	file, _ := os.Open("./github_issues.csv")
	defer file.Close()

	extract(file)
}
