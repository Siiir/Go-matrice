package matrix

import (
	"fmt"
	"strings"
)

//Describing functions
func (M_p *Matrix) RawPrint() {
	for m := uint(0); m < M_p.rowN; m++ {
		for n := uint(0); n < M_p.colN; n++ {
			fmt.Printf("%v", M_p.Get(m, n))
		}
		fmt.Println()
	}
}

func (M_p *Matrix) SprintDims() string {
	return fmt.Sprintf("%v×%v", M_p.rowN, M_p.colN)
}
func (M_p *Matrix) AsciiSprintDims() string {
	return fmt.Sprintf("%vx%v", M_p.rowN, M_p.colN)
}
func (M_p *Matrix) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("matrix %s:\n", M_p.SprintDims()))
	for m := uint(0); m < M_p.rowN; m++ {
		row := M_p.GetRow(m)
		sb.WriteString(fmt.Sprintf("%v\n", row)) // Let's see how this wil work.
	}
	return sb.String()
}
