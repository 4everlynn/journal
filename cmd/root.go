package cmd

import (
	"diswares.com.journal/config"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string
var cfg config.JournalConfig

func GetConfig() config.JournalConfig {
	return cfg
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "journal",
	Short: "一个小而美，简单又不失性能的工作报告命令行工作(CLI)(基于代码版本管理仓库)",
	Long: `一个小而美，简单又不失性能的工作报告生成器(基于代码版本管理仓库)
使用例子
	./journal stylish --day 生成工作日报
	./journal stylish --week 生成工作周报
	./journal stylish --month 生成工作月报
`,
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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .journal.yaml in run-directory)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
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

		viper.SetConfigName(".journal")
		// Search config in home directory with name ".journal" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		// find configuration files in the current directory
		viper.AddConfigPath(".")
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		// fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Println("err when decode config")
	}
}
