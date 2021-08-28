package main

import "fmt"

/**
* Definicion  de variables globales
*
* isDebug = Imprime en pantalla paso a paso
**/
var isDebug bool = false

/**
* Definicion de modelos MUTANTS
**/
type Mutant struct {
	x int
	y int
	t string
}

type Mutants struct {
	items []Mutant
}

func (x *Mutants) addMutant(newItem Mutant) []Mutant {
	x.items = append(x.items, newItem)
	return x.items
}

/**
* Struct de logica mutants
* matrix: arreglo bidireccional con ADN
* len: longitud de coincidencia
* mutants: atributo de almacenamiento de mutants, el resultado
**/
type mutantsModel struct {
	matrix  [][]string
	len     int
	mutants Mutants
}

/**
* Logica
**/
func (m mutantsModel) getMutants() Mutants {
	m.mutants = Mutants{[]Mutant{}}
	xLen := len(m.matrix)
	for x := 0; x < xLen; x++ {
		yLen := len(m.matrix[x])
		for y := 0; y < yLen; y++ {
			r := (x == 0 && y == 0) || !isDebug
			if r {
				if isDebug {
					fmt.Println("Position", x, y)
				}

				if m.validate(x, y, "Rigth") {
					m.mutants.addMutant(Mutant{x: x, y: y, t: "Rigth"})
				} else if m.validate(x, y, "Down") {
					m.mutants.addMutant(Mutant{x: x, y: y, t: "Down"})
				} else if m.validate(x, y, "DownVertical") {
					m.mutants.addMutant(Mutant{x: x, y: y, t: "DownVertical"})
				} else if m.validate(x, y, "Top") {
					m.mutants.addMutant(Mutant{x: x, y: y, t: "Top"})
				} else if m.validate(x, y, "TopVertical") {
					m.mutants.addMutant(Mutant{x: x, y: y, t: "TopVertical"})
				}
			}

		}
	}
	if isDebug {
		fmt.Println(" Result:", m.mutants)
	}
	return m.mutants
}

/**
* Validar con base a la posicion de ciclo, y depende de la orientacion factores en comun
**/
func (m mutantsModel) validate(x int, y int, t string) bool {
	if isDebug {
		fmt.Println(" >>> ", t)
	}

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

			if isDebug {
				fmt.Println(" >>>>>>>", m.matrix[xSS][ySS], " == ", m.matrix[x][y])
			}
			resultM := m.matrix[xSS][ySS] == m.matrix[x][y]
			if !resultM {
				isMutant = false
			}
		}
		if isDebug {
			fmt.Println(" >>>>>>> IsMutant ", isMutant)
		}
		if isMutant {
			return isMutant
		}
	}
	return false
}

/**
* Main
* Define el adn, ejecuta el llamo de analisis de la estructura
**/
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
	r := m.getMutants()
	fmt.Println("Mutants:", r)
}
