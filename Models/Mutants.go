package Models

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
type MutantsModel struct {
	X           int
	Matrix      [][]string
	Len         int
	ListMutants Mutants
}

/**
* Validar con base a la posicion de ciclo, y depende de la orientacion factores en comun
**/
func (m MutantsModel) validate(x int, y int, t string) bool {
	if isDebug {
		fmt.Println(" >>> ", t)
	}

	yS := y + m.Len
	xS := x + m.Len
	xS2 := x - m.Len
	var resultIni bool
	switch {
	case t == "Rigth":
		resultIni = yS < len(m.Matrix[x])
	case t == "Down":
		resultIni = xS < len(m.Matrix)
	case t == "DownVertical":
		resultIni = xS < len(m.Matrix) && yS < len(m.Matrix[x])
	case t == "Top":
		resultIni = xS2 > 0 && xS2 < len(m.Matrix)
	case t == "TopVertical":
		resultIni = xS2 > 0 && xS2 < len(m.Matrix) && yS < len(m.Matrix[x])
	default:
		resultIni = false
	}
	if resultIni {
		var isMutant bool = true
		for i := 1; i <= m.Len; i++ {
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
				fmt.Println(" >>>>>>>", m.Matrix[xSS][ySS], " == ", m.Matrix[x][y])
			}
			resultM := m.Matrix[xSS][ySS] == m.Matrix[x][y]
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
* Logica
**/
func (m MutantsModel) GetMutants() Mutants {
	m.ListMutants = Mutants{}
	xLen := len(m.Matrix)
	for x := 0; x < xLen; x++ {
		yLen := len(m.Matrix[x])
		for y := 0; y < yLen; y++ {
			r := (x == 0 && y == 0) || !isDebug
			if r {
				if isDebug {
					fmt.Println("Position", x, y)
				}

				if m.validate(x, y, "Rigth") {
					m.ListMutants.addMutant(Mutant{x: x, y: y, t: "Rigth"})
				} else if m.validate(x, y, "Down") {
					m.ListMutants.addMutant(Mutant{x: x, y: y, t: "Down"})
				} else if m.validate(x, y, "DownVertical") {
					m.ListMutants.addMutant(Mutant{x: x, y: y, t: "DownVertical"})
				} else if m.validate(x, y, "Top") {
					m.ListMutants.addMutant(Mutant{x: x, y: y, t: "Top"})
				} else if m.validate(x, y, "TopVertical") {
					m.ListMutants.addMutant(Mutant{x: x, y: y, t: "TopVertical"})
				}
			}

		}
	}
	if isDebug {
		fmt.Println(" Result:", m.ListMutants)
	}
	return m.ListMutants
}
