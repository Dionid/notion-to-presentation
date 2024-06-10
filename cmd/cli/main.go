package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Dionid/notion-to-presentation/features"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "n2p",
		Short: "Notion to Presentation",
	}

	var generateCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate a presentation from a Notion page",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Fatalf("Error: %s", "No url provided")
			}

			targetUrl := args[len(args)-1]

			err := features.FormFullHtmlPageFromNotion(
				targetUrl,
			)
			if err != nil {
				log.Fatalf("Error: %s", err)
			}
		},
	}

	rootCmd.AddCommand(generateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
