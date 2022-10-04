package matrix

import (
	"fmt"
)

func (M1_p *Matrix) ShapeAfterMult(M2_p *Matrix) (rowN uint, colN uint) {
	rowN = M1_p.rowN
	colN = M2_p.colN
	return
}

func (M1_p *Matrix) _tryProd(M2_p *Matrix) (M_p *Matrix) {
	// Exceptions
	if !M1_p.DoesProdExist(M2_p) {
		return nil
	}

	// Init
	M_p.rowN, M_p.colN = M1_p.ShapeAfterMult(M2_p)
	M_p.slice1D = make([]float64, M_p.rowN*M_p.colN) //This will be zero-init

	// The main algorithm that fills zero initialized M with appropiate scalar products
	for m := uint(0); m < M_p.rowN; m++ { //For every row `m` in M1
		M1_row := M1_p.GetRow(m)
		for n := uint(0); n < M_p.colN; n++ { //For each col `n` in M2
			for m2, rowVal := range M1_row {
				colVal := M2_p.Get(uint(m2), n)
				prod := rowVal * colVal
				M_p.slice1D[m*M_p.colN+n] += prod //Default val. is zero
			}
		}
	}
	return
}

func (M1_p *Matrix) TryProd(M2_p *Matrix) (M_p *Matrix, e error) {
	M_p = M1_p._tryProd(M2_p)
	if M_p == nil {
		return nil, fmt.Errorf(
			"impossible to multilply M1 and M2.\n"+
				"M1 row lenght is %v, whereas M2 col length is %v",
			M1_p.colN, M2_p.rowN,
		)

	}
	return
}

func (M1_p *Matrix) Reciprocal() (M2_p *Matrix)
