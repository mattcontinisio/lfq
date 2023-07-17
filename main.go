// Copyright 2023 Matthew Continisio
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/mattcontinisio/lfq/lfq"
)

// Flags

var input string          // Input format (logfmt or json)
var output string         // Output format (value, logfmt, or json)
var noColor bool          // Disable colorized output
var forceQuote bool       // Force quoting of all values (for value and logfmt output)
var disableQuote bool     // Disable quoting of all values (for value and logfmt output)
var quoteEmptyFields bool // Wrap empty values in quotes (for value and logfmt output)

// Root command (i.e. the only command)
var rootCmd = &cobra.Command{
	Short: "command-line logfmt processor",
	Long: `lfq is a tool for processing logfmt inputs, filtering keys, and producing results as values, logfmt, or JSON on
standard output.`,
	Use:     "lfq [flags] <keys>",
	Example: "lfq -o value time,level",
	Run: func(cmd *cobra.Command, args []string) {
		// Exit if data is not being piped to stdin
		stat, _ := os.Stdin.Stat()
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			// data is being piped to stdin
		} else {
			// stdin is from a terminal
			_ = cmd.Help()
			os.Exit(0)
		}

		// Disable colors
		// The environment variable NO_COLOR can also disable colors
		if noColor {
			color.NoColor = true
		}

		// Input
		var reader lfq.Reader
		if input == "logfmt" {
			reader = lfq.NewLogfmtReader()
		} else if input == "json" {
			reader = lfq.NewJsonReader()
		} else {
			_ = cmd.Help()
			os.Exit(1)
		}

		// Output
		var writer lfq.Writer
		if output == "logfmt" {
			writer = lfq.NewLogfmtWriter(lfq.LogfmtWriter{
				ForceQuote:       forceQuote,
				DisableQuote:     disableQuote,
				QuoteEmptyFields: quoteEmptyFields,
			})
		} else if output == "json" {
			writer = lfq.NewJsonWriter()
		} else if output == "value" {
			writer = lfq.NewValueWriter(lfq.LogfmtWriter{
				ForceQuote:       forceQuote,
				DisableQuote:     disableQuote,
				QuoteEmptyFields: quoteEmptyFields,
			})
		} else {
			_ = cmd.Help()
			os.Exit(1)
		}

		// Processors (just basic filtering for now)
		processors := []lfq.Processor{}
		keys := []string{}
		if len(args) > 0 {
			keys = strings.Split(args[0], ",")
		}
		processors = append(processors, &lfq.FilterProcessor{Keys: keys})

		// Scan
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			m, err := reader.Read(scanner.Bytes())
			if err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}

			for _, p := range processors {
				m = p.Process(m)
			}

			err = writer.Write(m)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error:", err)
				os.Exit(1)
			}
		}

		err := scanner.Err()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "logfmt", "input format (logfmt or json)")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "logfmt", "output format (value, logfmt, or json)")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable colorized output")
	rootCmd.PersistentFlags().BoolVar(&forceQuote, "force-quote", false, "force quoting of all values (for value and logfmt output)")
	rootCmd.PersistentFlags().BoolVar(&disableQuote, "disable-quote", false, "disable quoting of all values (for value and logfmt output)")
	rootCmd.PersistentFlags().BoolVar(&quoteEmptyFields, "quote-empty-fields", false, "wrap empty values in quotes (for value and logfmt output)")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
