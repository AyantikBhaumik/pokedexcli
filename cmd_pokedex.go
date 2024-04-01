package main

import "fmt"

func callbackPokedex(cfg *config, param string) error {
	fmt.Println("Your Pokedex: ")
	for k := range cfg.pokedex {
		fmt.Printf(" - %s\n", k)
	}
	fmt.Println()

	return nil
}