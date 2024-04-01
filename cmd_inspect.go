package main

import "fmt"

func callbackInspect(cfg *config, param string) error {
	dat, ok := cfg.pokedex[param]
	if !ok {
		fmt.Println("you have not caught that pokemon!!")
		fmt.Println()
		return nil
	}

	fmt.Printf("Name: %s\n", param)
	fmt.Printf("Height: %v\n", dat.Height)
	fmt.Printf("Weight: %v\n", dat.Weight)
	fmt.Println("Stats:")
	for _, stat := range dat.Stats {
		fmt.Printf(" -%s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, stat := range dat.Types {
		fmt.Printf(" - %s\n", stat.Type.Name)
	}
	fmt.Println()
	return nil
}