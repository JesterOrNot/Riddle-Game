package main

import (
	"fmt"
	"strings"
	"errors"
	"github.com/manifoldco/promptui"
)


func main() {
	riddles := map[string]string {
		"What Goes up but never comes down": "Your Age",
	}
	validate := func(input string) error {
		if strings.ToLower(input) != strings.ToLower(riddles["What Goes up but never comes down"]) {
			return errors.New("")
		}
		return nil
	}
	prompt := promptui.Prompt{
		Label:    "What Goes up but never comes down",
		Validate: validate,
	}
	prompt.Run()
}