package utils

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func GetBucketFileTimestamp(file string) (int64, error) {
	reg := regexp.MustCompile(`mongodb-snapshot-(?P<Time>\d+)\.(gz|log)`)
	match := reg.FindStringSubmatch(file)
	if len(match) != 3 {
		return 0, fmt.Errorf("File does not match pattern in folder: " + file)
	}

	timestamp, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("File has invalid timestamp in folder: " + file)
	}

	return timestamp, nil
}

func GetHumanFileSize(filename string) string {
	stat, err := os.Stat(filename)
	if err != nil {
		return "UNKNOWN"
	}
	return GetHumanBytes(stat.Size())
}
