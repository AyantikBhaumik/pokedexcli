package main

import (
	"fmt"
	"math/rand"
)

func callbackCatch(c *config, param string) error {
	res, err := c.pokeapiClient.FindPokemon(param)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", param)

	randomNumber := rand.Intn(5) + 1

	if randomNumber==1 {
		fmt.Printf("%s was caught!\n", param)
		fmt.Println()
		c.pokedex[param] = res
	} else {
		fmt.Printf("%s escaped!\n", param)
		fmt.Println()
	}

	return nil
}