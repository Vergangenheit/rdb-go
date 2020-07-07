package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tommy351/rdb-go"
)

// nolint: gochecknoglobals
var (
	outputFormat string

	rootCmd = &cobra.Command{
		Use:  "rdb [path]",
		Args: cobra.MaximumNArgs(1),
		Example: formatExamples(map[string]string{
			"Parse a RDB dump file.": "rdb path/to/dump.rdb",
			"Read RDB from stdin.":   "cat file | rdb",
		}),
		RunE: func(cmd *cobra.Command, args []string) error {
			var printer Printer

			writer := bufio.NewWriter(os.Stdout)
			defer writer.Flush()

			switch outputFormat {
			case "json":
				printer = NewJSONPrinter(writer)
			default:
				// nolint: goerr113
				return fmt.Errorf("unsupported format %q", outputFormat)
			}

			var reader io.Reader

			if len(args) > 0 {
				file, err := os.Open(args[0])

				if err != nil {
					return err
				}

				defer file.Close()

				reader = file
			} else {
				reader = bufio.NewReader(os.Stdin)
			}

			return printParserData(reader, printer)
		},
	}
)

func formatExamples(examples map[string]string) string {
	lines := make([]string, 0, len(examples))
	indent := "  "

	for title, content := range examples {
		lines = append(lines, indent+"# "+title+"\n"+indent+content)
	}

	return strings.Join(lines, "\n\n")
}

func main() {
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "json", "output format")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printParserData(reader io.Reader, printer Printer) error {
	parser := rdb.NewParser(reader)

	if err := printer.Start(); err != nil {
		return err
	}

	for {
		data, err := parser.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		switch v := data.(type) {
		case *rdb.StringData:
			err = printer.String(v)
		case *rdb.ListHead:
			err = printer.ListHead(v)
		case *rdb.ListEntry:
			err = printer.ListEntry(v)
		case *rdb.ListData:
			err = printer.ListData(v)
		case *rdb.SetHead:
			err = printer.SetHead(v)
		case *rdb.SetEntry:
			err = printer.SetEntry(v)
		case *rdb.SetData:
			err = printer.SetData(v)
		case *rdb.SortedSetHead:
			err = printer.SortedSetHead(v)
		case *rdb.SortedSetEntry:
			err = printer.SortedSetEntry(v)
		case *rdb.SortedSetData:
			err = printer.SortedSetData(v)
		case *rdb.HashHead:
			err = printer.HashHead(v)
		case *rdb.HashEntry:
			err = printer.HashEntry(v)
		case *rdb.HashData:
			err = printer.HashData(v)
		}

		if err != nil {
			return err
		}
	}

	return printer.End()
}
