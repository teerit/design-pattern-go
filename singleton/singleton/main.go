package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

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
	// Use sync.Once to ensure that initializeDatabase is called only once
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

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	} // depend directly on Db instance
	return result
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	DummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.DummyData) == 0 {
		d.DummyData = map[string]int{
			"alpha": 1,
		}
	}
	return d.DummyData[name]
}

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Population of Seoul: ", pop)

	cities := []string{"Seoul", "Mexico City"}
	tp := GetTotalPopulation(cities)
	ok := tp == (17500000 + 17400000)
	fmt.Println(ok)

	names := []string{"alpha"}
	tp = GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 1)

	for i := 0; i < 20; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
