package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const jsonFile = "people.json"

// The mostSuspicious returns a slice of full names of persons with the maximum number of crimes.
// If there are no persons who committed the crimes, returns nil.
func mostSuspicious(susp []string, db map[string]Man) (names []string) {
	var maxCrimes int
	for _, name := range susp {
		if s, ok := db[name]; ok && s.Crimes > maxCrimes {
			maxCrimes = s.Crimes
		}
	}

	// The task do not say whether a person with zero crimes should be considered suspicious.
	// if maxCrimes == 0 {
	// 	return
	// }

	for _, name := range susp {
		if s, ok := db[name]; ok && s.Crimes == maxCrimes {
			names = append(names, name)
		}
	}

	return
}

func display(criminals []string) {
	if len(criminals) > 0 {
		fmt.Printf("Наибольшее кол-во совершённых преступлений у: %s\n", strings.Join(criminals, ", "))
	} else {
		fmt.Println("В базе данных нет информации по запрошенным подозреваемым")
	}
}

func main() {
	file, err := os.Open(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var peopleSlice []Man
	err = json.Unmarshal(bytes, &peopleSlice)
	if err != nil {
		log.Fatal(err)
	}

	var people = make(map[string]Man)
	for _, p := range peopleSlice {
		people[p.Name+" "+p.LastName] = p
	}

	var suspects = []string{"Jessica Miller", "David Brown", "Daniel Garcia", "Michael Johnson"}
	criminals := mostSuspicious(suspects, people)

	display(criminals)
}
