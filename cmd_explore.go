package main

import "fmt"

func callbackExplore(cfg *config, param string) error {
	res, err := cfg.pokeapiClient.ExploreLocationAreas(param)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", param)
	fmt.Println("Found Pokemon:")
	for _, element := range res.PokemonEncounters {
		fmt.Printf(" - %s\n", element.Pokemon.Name)
	}
	fmt.Println()
	return nil
}