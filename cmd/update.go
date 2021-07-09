package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/4everlynn/journal/config"
	"github.com/4everlynn/journal/global"
	"github.com/4everlynn/journal/support"
	"github.com/gookit/color"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const GithubApi = "https://api.github.com/repos/4everlynn/journal/releases/latest"

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "get the latest journal version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get(GithubApi)
		if err != nil {
			fmt.Println("Fatal error ", err.Error())
			os.Exit(0)
		}
		var git = support.GitRelease{}

		all, err := ioutil.ReadAll(response.Body)
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Fatal error ", err.Error())
			}
		}(response.Body)

		err = json.Unmarshal(all, &git)

		if err != nil {
			fmt.Println(err)
		}
		// start parse
		handleGit(git)
	},
}

func handleGit(git support.GitRelease) {
	version := strings.Split(git.TagName, ".")
	master, _ := strconv.Atoi(version[0])
	mirror, _ := strconv.Atoi(version[1])
	patch, _ := strconv.Atoi(version[2])
	versionConfig := config.Version{
		Master: master,
		Mirror: mirror,
		Patch:  patch,
	}

	if JournalVersion.Master < versionConfig.Master ||
		JournalVersion.Mirror < versionConfig.Mirror ||
		JournalVersion.Patch < versionConfig.Patch {
		fmt.Println(color.FgLightYellow.Render("     ____.                                 .__   \n    |    | ____  __ _________  ____ _____  |  |  \n    |    |/  _ \\|  |  \\_  __ \\/    \\\\__  \\ |  |  \n/\\__|    (  <_> )  |  /|  | \\/   |  \\/ __ \\|  |__\n\\________|\\____/|____/ |__|  |___|  (____  /____/\n                                  \\/     \\/      "))
		fmt.Printf("%s\n%s\n",
			color.FgCyan.Render("Found new version ",
				buildVersion(),
				" => ",
				git.TagName,
			),
			color.FgLightGreen.Render("Topic is: ", git.Name),
		)
		download(git)
	} else {
		fmt.Println(color.FgLightMagenta.Render("UP TO DATE"))
	}

}

func download(git support.GitRelease) {
	assets := git.Assets
	if len(assets) > 0 {
		asset := getDownloadUrl(assets)
		if len(asset.Url) > 0 {
			response, _ := http.Get(asset.Url)
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					fmt.Println(err.Error())
				}
			}(response.Body)

			path := env()
			f, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					fmt.Println(err.Error())
				}
			}(f)

			bar := progressbar.DefaultBytes(
				response.ContentLength,
				"Downloading...",
			)
			_, err := io.Copy(io.MultiWriter(f, bar), response.Body)

			if err != nil {
				return
			}

			fmt.Printf("%s\n",
				color.Blue.Render("Done 🍬"),
			)
		}

	}
}

func env() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}

func getDownloadUrl(assets []support.GitReleaseAsset) support.GitReleaseAsset {
	for index := range assets {
		asset := assets[index]
		asset.Name = strings.ReplaceAll(asset.Name, "osx", global.OSX)
		if strings.Contains(asset.Name, global.GetRuntime()) {
			return asset
		}
	}
	return assets[0]
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
