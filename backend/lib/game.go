package lib

import (
	"fmt"
	"math"
)

var lowSegment = map[uint8][]float64{
	10: {
		1.2, 1.2, 1.2, 1.2, 0, 1.5, 1.2,
		1.2, 1.2, 0,
	},
	20: {
		1.2, 1.2, 1.2, 1.2, 0, 1.5, 1.2,
		1.2, 1.2, 0, 1.2, 1.2, 1.2, 1.2,
		0, 1.5, 1.2, 1.2, 1.2, 0,
	},
	30: {
		1.2, 1.2, 1.2, 1.2, 0, 1.5, 1.2,
		1.2, 1.2, 0, 1.2, 1.2, 1.2, 1.2,
		0, 1.5, 1.2, 1.2, 1.2, 0, 1.2,
		1.2, 1.2, 1.2, 0, 1.5, 1.2, 1.2,
		1.2, 0,
	},
	40: {
		1.2, 1.2, 1.2, 1.2, 0, 1.5, 1.2,
		1.2, 1.2, 0, 1.2, 1.2, 1.2, 1.2,
		0, 1.5, 1.2, 1.2, 1.2, 0, 1.2,
		1.2, 1.2, 1.2, 0, 1.5, 1.2, 1.2,
		1.2, 0, 1.2, 1.2, 1.2, 1.2, 0,
		1.5, 1.2, 1.2, 1.2, 0,
	},
	50: {
		1.2, 1.2, 1.2, 1.2, 0, 1.5, 1.2,
		1.2, 1.2, 0, 1.2, 1.2, 1.2, 1.2,
		0, 1.5, 1.2, 1.2, 1.2, 0, 1.2,
		1.2, 1.2, 1.2, 0, 1.5, 1.2, 1.2, 1.2, 0, 1.2, 1.2, 1.2, 1.2, 0,
		1.5, 1.2, 1.2, 1.2, 0, 1.2,
		1.2, 1.2, 1.2, 0, 1.5, 1.2,
		1.2, 1.2, 0,
	},
}

var mediumSegment = map[uint8][]float64{
	10: {
		0, 1.5, 0, 3, 0, 1.9, 0, 1.5, 2, 0,
	},
	20: {
		0, 2, 0, 2, 0, 2, 0, 1.5, 0, 2,
		0, 2, 0, 2, 0, 1.5, 0, 3, 0, 1.8,
	},
	30: {
		0, 1.5, 0, 2, 0, 3, 0, 1.5, 0, 2,
		0, 4, 0, 1.5, 0, 2, 0, 1.7, 0, 2,
		0, 1.5, 0, 2, 0, 1.5, 0, 2, 0, 0,
	},
	40: {
		0, 1.5, 0, 2, 0, 3, 0, 1.5, 0, 2,
		0, 3, 0, 1.5, 0, 2, 0, 1.5, 0, 2,
		0, 3, 0, 1.5, 0, 2, 0, 1.6, 0, 2,
		0, 3, 0, 1.5, 0, 2, 0, 1.5, 0, 0,
	},
	50: {
		0, 1.5, 0, 2, 0, 1.5, 0, 3, 0, 1.5,
		0, 2, 0, 1.5, 0, 1.5, 0, 3, 0, 2,
		0, 1.5, 0, 1.5, 0, 2, 0, 1.5, 0, 1.5,
		0, 2, 0, 3, 0, 1.5, 0, 1.5, 0, 2,
		0, 1.5, 0, 5, 0, 2, 0, 1.5, 0, 0,
	},
}

var highSegment = map[uint8][]float64{
	10: {0, 0, 0, 0, 0, 0, 0, 0, 0, 9.9},
	20: {
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 19.8,
	},
	30: {
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 29.7,
	},
	40: {
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 39.6,
	},
	50: {
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 49.5,
	},
}

func WheelResult(risk string, segments uint8, index float64) float64 {
	var result float64
	if risk == "low" {
		result = lowSegment[segments][int(index)]
	} else if risk == "medium" {
		result = mediumSegment[segments][int(index)]
	} else if risk == "high" {
		result = highSegment[segments][int(index)]
	}

	return result
}

func GetDragonTowerLevel(difficulty string, serverSeed string, clientSeed string, nonce int) [][]int {
	if difficulty == "easy" {
		var level [][]int
		for i := range 9 {
			row := []int{1, 1, 1, 1}
			for _ = range 3 {
				f, _, _ := RandomUInts(serverSeed, clientSeed, nonce, i+1)
				index := math.Floor(f * 4)
				row[int(index)] = 0
			}
			level = append(level, row)

		}
		fmt.Println("level", level)
		return level
	} else if difficulty == "medium" {
		var level [][]int
		for i := range 9 {
			row := []int{1, 1, 1}
			for _ = range 3 {
				f, _, _ := RandomUInts(serverSeed, clientSeed, nonce, i+1)
				index := math.Floor(f * 3)
				row[int(index)] = 0
			}
			level = append(level, row)
		}
		fmt.Println("level", level)
		return level
	} else if difficulty == "hard" {
		var level [][]int
		for i := range 9 {
			row := []int{1, 1}
			for _ = range 3 {
				f, _, _ := RandomUInts(serverSeed, clientSeed, nonce, i+1)
				index := math.Floor(f * 2)
				row[int(index)] = 0
			}
			level = append(level, row)
		}
		fmt.Println("level", level)
		return level
	} else if difficulty == "expert" {
		var level [][]int
		for i := range 9 {
			row := []int{0, 0, 0}
			for _ = range 3 {
				f, _, _ := RandomUInts(serverSeed, clientSeed, nonce, i+1)
				index := math.Floor(f * 3)
				row[int(index)] = 1
			}
			level = append(level, row)
		}
		fmt.Println("level", level)
		return level
	} else if difficulty == "master" {
		var level [][]int
		for i := range 9 {
			row := []int{0, 0, 0, 0}
			for _ = range 3 {
				f, _, _ := RandomUInts(serverSeed, clientSeed, nonce, i+1)
				index := math.Floor(f * 4)
				row[int(index)] = 1
			}
			level = append(level, row)
		}
		fmt.Println("level", level)
		return level
	}
	return [][]int{}
}

func GetMinesSet(minesCount int, serverSeed string, clientSeed string, nonce int) map[int]bool {
	minesSet := make(map[int]bool)
	i := 0
	for len(minesSet) < minesCount {
		f, _, _ := RandomUInts(serverSeed, clientSeed, nonce, i)
		index := math.Floor(f * 25)
		minesSet[int(index)] = true
		i++
	}
	fmt.Println("minesSet", minesSet)
	return minesSet
}
