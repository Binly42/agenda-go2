package model

import (
	"bufio"
	"encoding/json"
	"entity"
	"log"
	"os"
	"sync"
)

var wg sync.WaitGroup
var fin = os.Stdin

// Load : load all resources for agenda.
func Load() {
	loadConfig()
	loadRegisteredUserList()
}

// Save : Save all data for agenda.
func Save() {
	saveRegisteredUserList()
	saveConfig()
}

func loadConfig() {
	fcfg, err := os.Open(AgendaConfigPath())
	if err != nil {
		log.Fatalf("Load config fail, for config path: %v\n", AgendaConfigPath())
	}
	decoder := json.NewDecoder(fcfg)

	entity.LoadConfig(decoder)
}
func saveConfig() error {
	fcfg, err := os.Create(AgendaConfigPath())
	if err != nil {
		log.Fatalf("Save config fail, for config path: %v\n", AgendaConfigPath())
	}
	encoder := json.NewEncoder(fcfg)

	return entity.SaveConfig(encoder)
}

func loadRegisteredUserList() {
	fin, err := os.Open(UserDataRegisteredPath())
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(fin)

	entity.LoadUsersAllRegistered(decoder)
}
func saveRegisteredUserList() error {
	fout, err := os.Create(UserDataRegisteredPath())
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(fout)

	if err := entity.SaveUsersAllRegistered(encoder); err != nil {
		log.Printf(err.Error()) // TODO: hadnle ?
		return err
	}
	return nil
}

// .....
// ref to before

func readInput() (<-chan string, error) {
	channel := make(chan string)
	scanner := bufio.NewScanner(fin)
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	wg.Add(1)
	go func() {
		for scanner.Scan() {
			channel <- scanner.Text() + "\n"
		}
		defer wg.Done()
		close(channel)
	}()

	return channel, nil
}
