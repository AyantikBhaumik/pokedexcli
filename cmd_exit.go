package main

import (
	"fmt"
	"os"
)

func callbackExit(c *config, param string) error {
	fmt.Println("Thank You. Visit Again")
	fmt.Println("============================================")
	os.Exit(0)
	return nil
}