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
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	qapla "github.com/42nerds/qapla-base/api/v1alpha1"
	tpl "github.com/42nerds/qaplactl/templates"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// mainFile, err := os.Create("main.go.generated")
		// if err != nil {
		//     fmt.Println(err)
		// }

		if _, err := os.Stat(args[0]); os.IsNotExist(err) {
			os.Mkdir(args[0], os.FileMode(0775))
		} else {
			log.Fatalf("Target Directory '%s' exists allready and is not empty", args[0])
		}

		if _, err := os.Stat(args[0] + "/manifests"); os.IsNotExist(err) {
			os.Mkdir(args[0]+"/manifests", os.FileMode(0775))
		}

		applicationFile, err := os.Create(args[0] + "/manifests/application.yaml")
		if err != nil {
			fmt.Println(err)
		}

		// mainTemplate := template.Must(template.New("main").Parse(string(tpl.MainTemplate())))
		// err = mainTemplate.Execute(mainFile, nil)
		// if err != nil {
		// 	fmt.Println(err)
		// }

		if len(args) == 0 {
			args[0] = "Example Application"
		}

		reg, err := regexp.Compile("[^a-zA-Z0-9-]+")
		if err != nil {
			log.Fatal(err)
		}
		escapedName := strings.ReplaceAll(args[0], " ", "-")
		escapedName = reg.ReplaceAllString(escapedName, "")
		escapedName = strings.ToLower(escapedName)

		a := qapla.Application{
			ObjectMeta: v1.ObjectMeta{
				Name: escapedName,
			},
			Spec: qapla.ApplicationSpec{
				DisplayName: args[0],
				IconSrc:     "https://example.com/Your/Application/Icon",
			},
		}

		applicationTemplate := template.Must(template.New("application").Parse(tpl.ApplicationTemplate))
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
