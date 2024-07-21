package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lgc",
	Short: "golgc is a logic compiler",
	Long:  `golgc is a logic compiler`,
	Run: func(cmd *cobra.Command, args []string) {

		// Validate is an lgc file
		if len(args) != 1 {
			fmt.Println("Usage: lgc <file.lgc>")
			os.Exit(1)
		}

		if args[0][len(args[0])-4:] != ".lgc" {
			fmt.Println("Must be an lgc file")
			os.Exit(1)
		}

		// Open file
		file, err := os.Open(args[0])
		if err != nil {
			panic(err)
		}
		defer file.Close()

		// Read file
		stat, err := file.Stat()
		if err != nil {
			panic(err)
		}
		bs := make([]byte, stat.Size())
		_, err = file.Read(bs)
		if err != nil {
			panic(err)
		}
		CompileString(string(bs))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
