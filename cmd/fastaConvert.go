/*
Copyright © 2022 Daniel Rivas <danielrivasmd@gmail.com>

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
	"log"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// convertCmd represents the convert command
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: `Convert ` + chalk.Yellow.Color("fasta") + ` files.`,
	Long: chalk.Green.Color("Daniel Rivas <danielrivasmd@gmail.com>") + `

Convert ` + chalk.Yellow.Color("fastq") + ` files to ` + chalk.Yellow.Color("fasta") + ` file format.
`,

	Example: ``,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		// TODO: finish logic
		log.Fatal("convert called")
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	fastaCmd.AddCommand(convertCmd)

	// flags

}

////////////////////////////////////////////////////////////////////////////////////////////////////
