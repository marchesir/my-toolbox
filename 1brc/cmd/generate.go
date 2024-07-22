/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// const enum of city codes.
const (
	LAX = iota
	DXB
	JFK
	NRT
	SYD
	CDG
	BOM
	YYZ
	HKG
	GRU
	FRA
	MEX
	AMS
	CPT
	MUC
	ICN
	SFO
	ZRH
	DEL
	JNB
)

// Define a type for CityCode.
type CityCode [3]byte

// Map of CityCode.
var cities = map[int]CityCode{
	LAX: {'L', 'A', 'X'},
	DXB: {'D', 'X', 'B'},
	JFK: {'J', 'F', 'K'},
	NRT: {'N', 'R', 'T'},
	SYD: {'S', 'Y', 'D'},
	CDG: {'C', 'D', 'G'},
	BOM: {'B', 'O', 'M'},
	YYZ: {'Y', 'Y', 'Z'},
	HKG: {'h', 'K', 'G'},
	GRU: {'G', 'R', 'U'},
	FRA: {'F', 'R', 'A'},
	MEX: {'M', 'E', 'X'},
	AMS: {'A', 'M', 'S'},
	CPT: {'C', 'P', 'T'},
	MUC: {'M', 'U', 'C'},
	ICN: {'I', 'C', 'N'},
	ZRH: {'Z', 'R', 'H'},
	DEL: {'D', 'E', 'L'},
	JNB: {'J', 'N', 'B'},
}

// Encapsulate weather a station reading.
type measurement struct {
	CityCode
	temperature float32
}

// 1) loop for upto 1 billion;
// 2) generate random float between -50.0-50.0;
// 3) get random citycode;
// 4) write out dat to file;

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate new data file that can be used by the run command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generate:TODO")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
