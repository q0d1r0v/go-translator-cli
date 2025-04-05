// cmd/translate.go

package cmd

import (
	"fmt"

	"github.com/q0d1r0v/go-translator-cli/translator"

	"github.com/spf13/cobra"
)

// translateCmd represents the translate command
var translateCmd = &cobra.Command{
	Use:   "translate",
	Short: "Translate text from one language to another",
	Run: func(cmd *cobra.Command, args []string) {
		from, _ := cmd.Flags().GetString("from")
		to, _ := cmd.Flags().GetString("to")
		text, _ := cmd.Flags().GetString("text")

		translatedText, err := translator.Translate(text, from, to)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Translated text:", translatedText)
		}
	},
}

func init() {
	rootCmd.AddCommand(translateCmd)

	// Define flags for translate command
	translateCmd.Flags().StringP("from", "f", "en", "Source language")
	translateCmd.Flags().StringP("to", "t", "ru", "Target language")
	translateCmd.Flags().StringP("text", "x", "", "Text to translate")
}
