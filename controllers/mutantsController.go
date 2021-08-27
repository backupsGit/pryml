package mutantsController

import (
	"fmt"
)

type mutanModel struct {
	x int
	y int
	t string
}

type mutansModel struct {
	mutants []mutanModel
	matrix  [][]string
	len     int
}

func (x mutansModel) setMutant() {
	fmt.Println(x)
}

func (x mutansModel) getMutants() {

}
