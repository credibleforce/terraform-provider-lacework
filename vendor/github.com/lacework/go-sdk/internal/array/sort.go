package array

import (
	"sort"
	"strings"
)

// Sort2D can be used 2d Arrays used for Table Outputs to sort headers
func Sort2D(slice [][]string) {
	for range slice {
		sort.Slice(slice[:], func(i, j int) bool {
			elem := slice[i][0]
			next := slice[j][0]
			switch strings.Compare(elem, next) {
			case -1:
				return true
			case 1:
				return false
			default:
				// When equal compare next element
				for x := 1; x < len(slice[i]); x++ {
					secondaryElem := slice[i][x]
					nextSecondaryElem := slice[j][x]
					switch strings.Compare(secondaryElem, nextSecondaryElem) {
					case -1:
						return true
					case 1:
						return false
					default:
						continue
					}
				}
				return false
			}
		})
	}
}
