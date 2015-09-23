/*
This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.
This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.
You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
// CONTRIBUTORS AND COPYRIGHT HOLDERS (c) 2015:
// Dag Robøle (dag D0T robole AT gmail D0T com)

package main

import (
	"bufio"
	"encoding/json"
	"os"
)

// Structure representing a sample writer
type sampleWriterJson struct {
	JsonFile string
	fd       *os.File
	fw       *bufio.Writer
	sep      string
}

// Create a new sample writer
func NewSampleWriterJson(jsonFile string) (SampleWriter, error) {

	// Initialize a sample writer
	sw := new(sampleWriterJson)
	sw.JsonFile = jsonFile
	sw.sep = ""

	var err error
	sw.fd, err = os.Create(sw.JsonFile)
	if err != nil {
		return nil, err
	}

	sw.fw = bufio.NewWriter(sw.fd)
	sw.fw.WriteString("[\n")

	return sw, nil
}

// Write a sample to the json file
func (sw *sampleWriterJson) Write(s *Sample) error {

	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	sw.fw.WriteString(sw.sep + string(b))

	if len(sw.sep) == 0 {
		sw.sep = ","
	}

	return nil
}

// Finish the json file
func (sw *sampleWriterJson) Close() error {

	sw.fw.WriteString("\n]")
	sw.fw.Flush()
	sw.fd.Close()

	return nil
}
