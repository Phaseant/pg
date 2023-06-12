/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/Phaseant/pg/cmd/constants"
	"github.com/spf13/cobra"
)

var length int
var numbers, symbols, uppercase, lowercase, ambiguous bool

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Generate a new password",
	Long:  `Generate a new password. Flags example: -l 16 --nums=false -syms=false --up=false --low=false -amb=false`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Your new password is: ", generatePassword())
	},
}

func init() {
	newCmd.Flags().IntVarP(&length, "len", "l", 16, "Length of the password")
	newCmd.Flags().BoolVar(&numbers, "nums", true, "Include numbers (0,1,2,3...)")
	newCmd.Flags().BoolVar(&symbols, "syms", true, "Include symbols (*,&,$,#...)")
	newCmd.Flags().BoolVar(&uppercase, "up", true, "Include uppercase letters")
	newCmd.Flags().BoolVar(&lowercase, "low", true, "Include lowercase letters ")
	newCmd.Flags().BoolVar(&ambiguous, "amb", false, "Include ambiguous symbols ({},(),`,~...)")

	rootCmd.AddCommand(newCmd)
}

func generatePassword() string {
	dict := generateDictionary()
	result := strings.Builder{}
	for i := 0; i < length; i++ {
		result.WriteByte(dict[rand.Intn(len(dict))])
	}
	return result.String()
}

// generateDictionary generates a dictionary based on the user's options
func generateDictionary() string {
	result_dict := strings.Builder{}
	if numbers {
		result_dict.WriteString(constants.Numbers)
	}

	if symbols {
		result_dict.WriteString(constants.Symbols)
	}

	if uppercase {
		result_dict.WriteString(constants.Uppercase)
	}

	if lowercase {
		result_dict.WriteString(constants.Lowercase)
	}

	if ambiguous {
		result_dict.WriteString(constants.Ambiguous)
	}

	return result_dict.String()
}
