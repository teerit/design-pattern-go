package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// syncOnce init() -- thread safety
// laziness
// init() — we could, but it's not lazy

func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}

	return result, nil
}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readData("capitals.txt")
		db := singletonDatabase{}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Population of Seoul: ", pop)
}
