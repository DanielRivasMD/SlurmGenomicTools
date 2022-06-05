/*
Copyright © 2021 Daniel Rivas <danielrivasmd@gmail.com>

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
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/biogo/biogo/alphabet"
	"github.com/biogo/biogo/io/seqio"
	"github.com/biogo/biogo/io/seqio/fasta"
	"github.com/biogo/biogo/seq/linear"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: documentation assembly sequence
// sequenceCmd represents the sequence command
var sequenceCmd = &cobra.Command{
	Use:   "sequence",
	Short: "Retrieve sequence from assembly",
	Long:  `Retrieve sequence from assembly from coordinates.`,

	Example: `` + chalk.Cyan.Color("bender") + ` assembly sequence
--assembly aweSap01.fa
--species Awesome_sapiens
--inDir pathToAssembly/
--scaffold Scaffold_1
--start	102400
--end	124600
--hood	25000
`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {

		// scaffold
		syncytin.scaffoldIdent = scaffoldID

		// positions
		syncytin.positionIdent.minMax(startCoor, endCoor)

		// declare file output
		if outFile == "" {
			outFile = coordinateOut(assembly)
			outFile = outFile + ".fasta"
		}

		// execute logic
		collectCoordinates(inDir + "/" + assembly)

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	assemblyCmd.AddCommand(sequenceCmd)

	// flags

}

////////////////////////////////////////////////////////////////////////////////////////////////////

// read file & collect sequences
func collectCoordinates(readFile string) {

	// open an input file, exit on error
	contentFile, err := ioutil.ReadFile(readFile)
	if err != nil {
		log.Fatal("Error opending input file :", err)
	}

	// check whether file exists to avoid appending
	if fileExist(outFile) {
		os.Remove(outFile)
	}

	// mount data string
	dataFasta := strings.NewReader(string(contentFile))

	// fasta.Reader requires a known type template to fill
	// with FASTA data. Here we use *linear.Seq.
	template := linear.NewSeq("", nil, alphabet.DNAredundant)
	readerFasta := fasta.NewReader(dataFasta, template)

	// make a seqio.Scanner to simplify iterating over a
	// stream of data.
	scanFasta := seqio.NewScanner(readerFasta)

	// iterate through each sequence in a multifasta and
	// examine the ID, description and sequence data.
	for scanFasta.Next() {
		// get the current sequence and type assert to *linear.Seq.
		// while this is unnecessary here, it can be useful to have
		// the concrete type.
		sequence := scanFasta.Seq().(*linear.Seq)

		// find scaffold
		if sequence.ID == syncytin.scaffoldIdent {

			// define extraction coordinates
			startI := int64(syncytin.positionIdent.startPos - 1 - hood)
			endI := int64(syncytin.positionIdent.endPos + hood)

			// verify limits
			if startI < 0 {
				startI = 0
			}

			if endI > int64(len(sequence.Seq)) {
				endI = int64(len(sequence.Seq) - 1)
			}

			// determine strand
			strand := ""
			if startCoor == syncytin.positionIdent.startPos {
				strand = "+"
			} else if startCoor == syncytin.positionIdent.endPos {
				strand = "-"
			} else {
				log.Fatal("Strand could not be determined")
			}

			// define record id
			id := syncytin.scaffoldIdent + "_" +
				strconv.FormatFloat(syncytin.positionIdent.startPos, 'f', 0, 64) + "_" +
				strconv.FormatFloat(syncytin.positionIdent.endPos, 'f', 0, 64) + " " +
				species + " " +
				syncytin.scaffoldIdent + " " +
				strconv.FormatInt(startI, 10) + " " +
				strconv.FormatInt(endI, 10) + " " +
				strand

			// find coordinates
			targatSeq := linear.NewSeq(id, sequence.Seq[startI:endI], alphabet.DNA)

			// write candidate
			writeFasta(outFile, targatSeq)
		}
	}

	if err := scanFasta.Error(); err != nil {
		log.Fatal(err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// write positions
func writeFasta(outFile string, sequence *linear.Seq) {

	// declare io
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	// declare writer
	w := fasta.NewWriter(f, 10000)

	// writing
	_, err = w.Write(sequence)

	if err != nil {
		panic(err)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////