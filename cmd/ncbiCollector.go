/*
Copyright © 2020 Daniel Rivas <danielrivasmd@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"bytes"
	"log"

	"os/exec"

	"github.com/DanielRivasMD/Bender/aux"
	"github.com/labstack/gommon/color"
	"github.com/spf13/cobra"
)

// ncbiCollectorCmd represents the ncbiCollector command
var ncbiCollectorCmd = &cobra.Command{
	Use:   "ncbiCollector",
	Short: "Collect sequences from NCBI",
	Long: `ncbiCollector uses ncbi-acc-download to collect
sequences from NCBI and place them locally
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {

		// lineBreaks
		aux.LineBreaks()

		// buffers
		var stdout bytes.Buffer
		var stderr bytes.Buffer

		// shell call
		commd := "/Users/drivas/bin/goTools/sh/ncbiCollector.sh"
		// commd := "/home/drivas/bin/goTools/sh/ncbiCollector.sh"

		shCmd := exec.Command(commd)

		// run
		shCmd.Stdout = &stdout
		shCmd.Stderr = &stderr
		err := shCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// stdout
		color.Println(color.Cyan(stdout.String(), color.B))

		// stderr
		if stderr.String() != "" {
			color.Println(color.Red(stderr.String(), color.B))
		}

		// lineBreaks
		aux.LineBreaks()

	},

	////////////////////////////////////////////////////////////////////////////////////////////////////
}

func init() {
	rootCmd.AddCommand(ncbiCollectorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ncbiCollectorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ncbiCollectorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}