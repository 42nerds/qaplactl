package cmd
/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import (
	"text/template"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/42nerds/qaplactl/templates"
)

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	DisplayName string                `json:"displayName"`
	IconSrc     string                `json:"iconSrc"`
	MenuItems   []ApplicationMenuItem `json:"menuItems,omitempty"`
}

// ApplicationMenuItem ...
type ApplicationMenuItem struct {
	Text  string                `json:"text"`
	Items []ApplicationMenuItem `json:"items,omitempty"`
	Href  string                `json:"href,omitempty"`
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mainFile, err := os.Create("main.go.generated")
    if err != nil {
        fmt.Println(err)
		}

		applicationFile, err := os.Create("application.yaml")
    if err != nil {
        fmt.Println(err)
		}
		
		mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
		err = mainTemplate.Execute(mainFile, nil)
		if err != nil {
			fmt.Println(err)
		}
		
		a := ApplicationSpec{DisplayName: args[0], IconSrc: "https://icons.com/headicon"}

		applicationTemplate := template.Must(template.New("application").Parse(string(tpl.ApplicationTemplate())))
		err = applicationTemplate.Execute(applicationFile, a)
		if err != nil {
			fmt.Println(err)
		}
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
