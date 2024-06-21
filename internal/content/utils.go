package content

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

func parseCoordinates(args []string) ([]float64, error) {
	coords := make([]float64, 9)
	for i := 0; i < 9; i++ {
		coord, err := strconv.ParseFloat(args[i], 64)
		if err != nil {
			return nil, fmt.Errorf("invalid input %s: %v", args[i], err)
		}
		coords[i] = coord
	}
	return coords, nil
}

func generateMetaDataContent(filePath, cid string, coords []float64) []byte {
	var metaData MetaData
	metaData.CID = cid
	fileName := filepath.Base(filePath)
	metaData.File = fileName
	metaData.X = coords[0]
	metaData.Y = coords[1]
	metaData.Z = coords[2]
	metaData.RX = coords[3]
	metaData.RY = coords[4]
	metaData.RZ = coords[5]
	metaData.SX = coords[6]
	metaData.SY = coords[7]
	metaData.SZ = coords[8]
	metaBytes, err := yaml.Marshal(metaData)
	if err != nil {
		fmt.Println("Error marshalling metadata: ", err)
		return nil
	}
	return metaBytes
}

func encryptAndCompressMetaData(metaDataContent []byte) ([]byte, error) {
	compressedMetaData, err := compressData(metaDataContent)
	if err != nil {
		return nil, err
	}

	return encryptData(compressedMetaData)
}

func extractFileDetailsFromMetaData(metaDataContent string) (string, string, error) {
	var fileCid, fileName string
	lines := strings.Split(metaDataContent, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "file: ") {
			fileName = strings.TrimPrefix(line, "file: ")
		} else if strings.HasPrefix(line, "cid: ") {
			fileCid = strings.TrimPrefix(line, "cid: ")
		}
	}

	if fileCid == "" {
		return "", "", fmt.Errorf("file CID not found in metadata")
	}
	return fileCid, fileName, nil
}
