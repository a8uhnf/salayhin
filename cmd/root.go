package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

func RootCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "salayhin",
		Short: "cli for salayhin",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("salayhin sucks")
		},
	}
}
