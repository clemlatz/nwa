// Copyright 2023 BINARY Members
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

const (
	name    = "nwa"
	version = "dev"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   name,
	Short: "A More Powerful License Header Management Tool",
	Long: `
███╗   ██╗██╗    ██╗ █████╗ 
████╗  ██║██║    ██║██╔══██╗
██╔██╗ ██║██║ █╗ ██║███████║
██║╚██╗██║██║███╗██║██╔══██║
██║ ╚████║╚███╔███╔╝██║  ██║
╚═╝  ╚═══╝ ╚══╝╚══╝ ╚═╝  ╚═╝
`,
	Version: version,
}

// Execute executes the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("{{ .Version }}")
	rootCmd.AddGroup(&cobra.Group{
		ID:    common,
		Title: "Common Mode Commands:",
	}, &cobra.Group{
		ID:    config,
		Title: "Config Mode Commands:",
	})
}

// TODO: nwa check??
const (
	common = "common"
	config = "config"
)

var (
	MuteF    bool
	HolderF  string
	YearF    string
	LicenseF string
	TmplF    string
	SkipF    []string
)

func setupCommonCmd(common *cobra.Command) {
	rootCmd.AddCommand(common)

	common.Flags().BoolVarP(&MuteF, "mute", "m", defaultConfig.Nwa.Mute, "mute mode")
	common.Flags().StringVarP(&HolderF, "copyright", "c", defaultConfig.Nwa.Holder, "copyright holder")
	common.Flags().StringVarP(&YearF, "year", "y", defaultConfig.Nwa.Year, "copyright year")
	common.Flags().StringVarP(&LicenseF, "license", "l", defaultConfig.Nwa.License, "license type")
	common.Flags().StringVarP(&TmplF, "tmpl", "t", defaultConfig.Nwa.Tmpl, "template file path")
	common.Flags().StringSliceVarP(&SkipF, "skip", "s", defaultConfig.Nwa.Skip, "skip file")

	common.MarkFlagsMutuallyExclusive("copyright", "tmpl")
	common.MarkFlagsMutuallyExclusive("year", "tmpl")
	common.MarkFlagsMutuallyExclusive("license", "tmpl")
	common.MarkFlagsMutuallyExclusive("skip", "tmpl")
}

func setupConfigCmd(config *cobra.Command) {
	rootCmd.AddCommand(config)
}