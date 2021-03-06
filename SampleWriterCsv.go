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
// Copyright: (c) 2015 Norwegian Radiation Protection Authority
// Contributors: Dag Robøle (dag D0T robole AT gmail D0T com)

package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

// SampleWriterCsv Structure representing a sample writer
type SampleWriterCsv struct {
	CsvFile       string
	UseScientific bool
	fd            *os.File
	fw            *csv.Writer
}

// NewSampleWriterCsv Create a new CSV sample writer
func NewSampleWriterCsv(csvFile string, useScientific bool) (SampleWriter, error) {

	// Initialize a sample writer
	sw := new(SampleWriterCsv)
	sw.CsvFile = csvFile
	sw.UseScientific = useScientific

	var err error
	sw.fd, err = os.Create(sw.CsvFile)
	if err != nil {
		return nil, err
	}

	sw.fw = csv.NewWriter(sw.fd)
	sw.fw.Write([]string{"Date", "Latitude", "Longitude", "Altitude", "Value", "Unit"})

	return sw, nil
}

// Write Write a sample to the csv file
func (sw *SampleWriterCsv) Write(s *Sample) error {

	// Set the number format
	mod := byte('f')
	if sw.UseScientific {
		mod = byte('E')
	}

	lat := strconv.FormatFloat(s.Latitude, 'f', 8, 64)
	lon := strconv.FormatFloat(s.Longitude, 'f', 8, 64)
	alt := strconv.FormatFloat(s.Altitude, 'f', 8, 64)
	val := strconv.FormatFloat(s.Value, mod, 8, 64)

	sw.fw.Write([]string{s.Date.String(), lat, lon, alt, val, s.Unit})

	return nil
}

// Close Finish the CSV file
func (sw *SampleWriterCsv) Close() error {

	sw.fw.Flush()
	sw.fd.Close()

	return nil
}
