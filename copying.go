package matrix

import "fmt"

func (M_p0 *Matrix) UncheckedCopyTo(M_p1 *Matrix) {
	for m := uint(0); m < M_p0.rowN; m++ {
		//Unoptimized yet. Implement with generic copy(...) for slices
		for n := uint(0); n < M_p0.colN; n++ {
			M_p1.Set(m, n, M_p0.Get(m, n))
		}
	}
}

func (M_p0 *Matrix) CopyTo(M_p1 *Matrix) error {
	if !M_p1.CanContain(M_p0) {
		return fmt.Errorf(
			"can fit a matrix %s in matrix %s",
			M_p0.SprintDims(), M_p1.SprintDims(),
		)
	}
	M_p0.UncheckedCopyTo(M_p1)
	return nil
}

func (M_p *Matrix) Clone() *Matrix {
	M := &Matrix{
		slice1D: make([]float64, len(M_p.slice1D)),
		rowN:    M_p.rowN,
		colN:    M_p.colN,
	}
	copy(M.slice1D, M_p.slice1D)
	return M
}
