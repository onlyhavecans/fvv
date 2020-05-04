package cmd

import (
	"fmt"
	"log"
	"os"
	"text/template"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const templateGlob = "*.tmpl"

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "fvv",
	Args:  cobra.ExactArgs(1),
	Short: "Quickly render go templates to stdout",
	Long: `Quick Usage: fvv <template definition name>

Fusozay Var Var is a CLI application for quickly rendering out text templates.
I often write outfit and character descriptions that reuses a lot of elements.
This allows me to DRY up my descriptions and still quickly get results.

Template requirements:
- All of the templates must be valid golang templates.
- They must not require any variables to be passed in.
- All templates must have the '.tmpl' extension. All other files ignored.
- The template you want to render must be "named"
https://golang.org/pkg/text/template/#hdr-Nested_template_definitions
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { renderTemplate(args[0]) },
}

func renderTemplate(definition string) {
	tmpl, err := template.ParseGlob(templateGlob)
	fatalIfErr("ParsingGlob", err)

	err = tmpl.ExecuteTemplate(os.Stdout, definition, "no data needed")
	fatalIfErr("ExecutingTemplate", err)
}

func fatalIfErr(helpMessage string, err error) {
	if err != nil {
		l := log.New(os.Stderr, "", 0)
		l.Fatal(helpMessage, ": ", err)
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fvv.toml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
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

		// Search config in home directory with name ".fvv" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".fvv")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
