package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"os"
	"time"
)

func ExecSpinner(cod int, color, prefix, suffix string) *spinner.Spinner {
	s := spinner.New(spinner.CharSets[cod], 100*time.Millisecond)
	s.Prefix = prefix
	s.Color(color)
	s.Suffix = suffix
	s.FinalMSG = "\nconcluded!"
	s.Start()
	return s
}

func main() {
	root := &cobra.Command{
		Use:   "root",
		Short: "Msg short",
		Long:  `Message long`,
		Run: func(cmd *cobra.Command, args []string) {
			defer ExecSpinner(11, "blue", "before: ", " after...").Stop()
			printLong("Root...")
		},
	}

	cmd := &cobra.Command{
		Use:   "cmd",
		Short: "Msg short",
		Long:  `Message long`,
		Run: func(cmd *cobra.Command, args []string) {
			defer ExecSpinner(35, "green", "before: ", " after...").Stop()
			printLong("Add...")
		},
	}

	root.AddCommand(cmd)

	if err := root.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printLong(c string) {
	fmt.Println(c)
	time.Sleep(4 * time.Second)
}
