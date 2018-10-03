// Copyright © 2018 Miguel Chan <vvchan@outlook.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/Miguel-Chan/selpg-go/selpg"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "selpg-go",
	Short: "A utility tool for selecting page",
	Long: fmt.Sprintf(`selpg is a command line utility tool that helps you select certain range of pages from your input to print.
You can specify your input with the last command line argument or by default input from stdin.

USAGE: %v -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]`, os.Args[0]),

	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			inputFile = ""
		} else {
			inputFile = args[0]
		}
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		if useFormFeed && lineNum != 72 {
			fmt.Printf("%v: Invalid flags: -lNumber and -f are exclusive.\n", os.Args[0])
			os.Exit(1)
		}
		if startPage <= 0 {
			fmt.Printf("%v: Invalid start page %v\n", os.Args[0], startPage)
			os.Exit(1)
		}
		if endPage < startPage {
			fmt.Printf("%v: Invalid end page %v\n", os.Args[0], endPage)
		}
		sp := selpg.NewSelpg(startPage, endPage, lineNum, destination, inputFile, useFormFeed)
		sp.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var startPage, endPage, lineNum int
var useFormFeed bool
var destination, inputFile string
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().StringVar(&cfgFile, "config", "", "config file (defalut in $HOME/.cobra.yaml)")
	rootCmd.Flags().IntVarP(&startPage, "start_page", "s", 0,  "The first page to be selected")
	rootCmd.Flags().IntVarP(&endPage, "end_page", "e", 0,  "The last page to be selected")
	rootCmd.Flags().IntVarP(&lineNum, "line_number", "l", 72,  "The number of lines in every page")
	rootCmd.Flags().BoolVarP(&useFormFeed, "form_feed", "f", false, "Use \\f as the seperator of each page")
	rootCmd.Flags().StringVarP(&destination, "destination", "d", "", "Specify printer for output instead of stdout")

	rootCmd.MarkFlagRequired("start_page")
	rootCmd.MarkFlagRequired("end_page")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".temp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".temp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
