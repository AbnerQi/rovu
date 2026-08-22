package cmd

import (
	"fmt"
	"strings"
)

const (
	KB = 1024
	MB = 1024 * KB
	GB = 1024 * MB
)

func formatSize(size int64) string {
	switch {
	case size < KB:
		return fmt.Sprintf("%d bytes", size)
	case size < MB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	case size < GB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	default:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	}
}

func barChart(count, total, bMax int) string {
	length := int(float64(count) / float64(total) * float64(bMax))

	if length == 0 {
		length = 1
	}

	return strings.Repeat("█", length)
}
