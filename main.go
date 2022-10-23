package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	log.Print("Begin remove duplicate tabs...")
	lines, err := readLines("./tabs.txt")
	if nil != err {
		log.Fatal(err)
	}

	noResultNum := 0
	if noResultNum == len(lines) {
		log.Print("There is no URL found, finish the program")
		return
	}

	log.Printf("Found %d results, begin removing duplication and sorting...", len(lines))

	//TODO improve speed by sorting the URL before removing process (removing process needs an improvement not just compare equirity on a whole value)
	lines = removeDuplicateLine(lines)

	sort.Strings(lines)
	err = writeLines(lines, "./result_tabs.txt")
	if nil != err {
		log.Fatal(err)
	}
	log.Printf("Finish removing and sorting with %d result", len(lines))
}

func removeDuplicateLine(lines []string) []string {
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if lines[i] == lines[j] {
				lines = remove(lines, j)
			}
		}
	}

	return lines
}

func remove(slice []string, index int) []string {
	if index > len(slice) || index < 0 {
		log.Printf("Unable to remove index %d from a slice of %d size", index, len(slice))
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if nil != err {
		return nil, err
	}

	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			lines = append(lines, scanner.Text())
		}
	}

	return lines, scanner.Err()
}

const REQUIRED_PARAM_ERROR = "required parameter is needed"

func writeLines(lines []string, path string) error {
	if nil == lines {
		return errors.New(REQUIRED_PARAM_ERROR)
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(writer, line)
	}

	return writer.Flush()
}
