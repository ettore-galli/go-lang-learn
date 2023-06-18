package csvutils

import (
	"encoding/csv"
	"io"
	"os"
)

type OutMap map[string]string

func ReadCsvMap(csvFileName string, firstValueIsOk bool) (OutMap, error) {

	csvFile, ferr := os.Open(csvFileName)
	if ferr != nil {
		return nil, ferr
	}

	defer csvFile.Close()

	return PerformReadCsvMap(csvFile, firstValueIsOk)

}

func PerformReadCsvMap(reader io.Reader, firstValueIsOk bool) (OutMap, error) {
	var m OutMap = OutMap{}

	csvReader := csv.NewReader(reader)

	for {
		record, err := csvReader.Read()
		if err != nil {
			break
		}

		if firstValueIsOk {
			_, okmap := m[record[0]]
			if !okmap {
				m[record[0]] = record[1] // TODO: Eliminare ridondanza
			}
		} else {
			m[record[0]] = record[1] // TODO: Eliminare ridondanza
		}

	}

	return m, nil
}
