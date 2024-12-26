package day4

import (
	"strings"
)

func Day4Part2(data string) any {
	height := strings.Count(data, "\n")
	width := strings.Index(data, "\n")

	dataAsBytes := []byte(strings.ReplaceAll(data, "\n", ""))
	founds := 0

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			pos := y*width + x

			if dataAsBytes[pos] == 'A' {
				tl := pos - width - 1
				br := pos + width + 1
				left := (dataAsBytes[tl] + dataAsBytes[br]) == 160

				if left {
					tr := pos - width + 1
					bl := pos + width - 1
					right := (dataAsBytes[tr] + dataAsBytes[bl]) == 160

					if right {
						founds++
					}
				}
			}
		}
	}

	return founds
}
