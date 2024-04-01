package main

import (
	"fmt"
)

func callbackMapb(cfg *config, param string) error {
	if(cfg.prevLocationAreaURL==nil) {
		fmt.Printf("You are on the first page\n")
		fmt.Println()
		return nil
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)
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