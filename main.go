package main

import "fmt"

func main() {
	var tm = [][]string{
		{"A", "A", "A", "A", "04", "05", "06", "07", "08"},
		{"A", "11", "12", "13", "14", "15", "16", "17", "18"},
		{"A", "21", "B", "23", "24", "25", "26", "27", "28"},
		{"A", "31", "32", "B", "34", "35", "36", "37", "38"},
		{"40", "41", "42", "43", "B", "45", "46", "47", "48"},
		{"50", "51", "52", "53", "54", "B", "56", "57", "58"},
		{"60", "61", "62", "63", "64", "65", "66", "67", "68"},
		{"70", "71", "72", "73", "74", "75", "76", "77", "78"},
		{"80", "81", "82", "83", "84", "85", "86", "87", "88"}}
	m := mutantsModel{len: 3, matrix: tm}
	m.getMutants()
}

type mutantsModel struct {
	matrix  [][]string
	len     int
	mutants [][]string
}

func (m mutantsModel) getMutants() {
	var result []mutantModel

	xLen := len(m.matrix)
	for x := 0; x < xLen; x++ {
		yLen := len(m.matrix[x])
		for y := 0; y < yLen; y++ {
			r := x == 0 && y == 0
			if r {
				//fmt.Println("Position", x, y)
				if m.validate(x, y, "Rigth") {
					result = []mutantModel{{1, 2, "Rigth"}}
				}
				if m.validate(x, y, "Down") {
					result = []mutantModel{{1, 2, "Down"}}
				}
				if m.validate(x, y, "DownVertical") {
					result = []mutantModel{{1, 2, "DownVertical"}}
				}
				if m.validate(x, y, "Top") {
					result = []mutantModel{{1, 2, "Top"}}
				}
				if m.validate(x, y, "TopVertical") {
					result = []mutantModel{{1, 2, "TopVertical"}}
				}
			}

		}
	}
	fmt.Println(" Result:", result)
}

func (m mutantsModel) validate(x int, y int, t string) bool {
	fmt.Println(" >>> ", t)
	yS := y + m.len
	xS := x + m.len
	xS2 := x - m.len

	var resultIni bool
	switch {
	case t == "Rigth":
		resultIni = yS < len(m.matrix[x])
	case t == "Down":
		resultIni = xS < len(m.matrix)
	case t == "DownVertical":
		resultIni = xS < len(m.matrix) && yS < len(m.matrix[x])
	case t == "Top":
		resultIni = xS2 > 0 && xS2 < len(m.matrix)
	case t == "TopVertical":
		resultIni = xS2 > 0 && xS2 < len(m.matrix) && yS < len(m.matrix[x])
	default:
		resultIni = false
	}
	if resultIni {
		var isMutant bool = true
		for i := 1; i <= m.len; i++ {
			var ySS int
			var xSS int

			switch {
			case t == "Rigth":
				ySS = y + i
				xSS = x
			case t == "Down":
				ySS = y
				xSS = x + i
			case t == "DownVertical":
				ySS = y + i
				xSS = x + i
			case t == "Top":
				ySS = y
				xSS = x - i
			case t == "TopVertical":
				ySS = y + i
				xSS = x - i
			}

			//fmt.Println(" >>>>>>>", m.matrix[xSS][ySS])
			fmt.Println(" >>>>>>>", m.matrix[xSS][ySS], " == ", m.matrix[x][y])
			resultM := m.matrix[xSS][ySS] == m.matrix[x][y]
			if !resultM {
				isMutant = false
			}
		}
		fmt.Println(" >>>>>>>", isMutant)
		if isMutant {
			return isMutant
		}
	}
	return false
}

type mutantModel struct {
	x int
	y int
	t string
}

func newMutantModel(x int, y int, t string) *mutantModel {
	m := mutantModel{x: x, y: y, t: t}
	return &m
}
