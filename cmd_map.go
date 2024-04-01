package main

import (
	"fmt"
)

func callbackMap(cfg *config, param string) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err!=nil {
		return err
	}
	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf(" - %s\n", area.Name)
	}
	fmt.Println()
	cfg.nextLocationAreaURL=res.Next
	cfg.prevLocationAreaURL=res.Previous
	return nil
}