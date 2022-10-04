package matrix

// Columns
func (M_p *Matrix) SetCol(newRow []float64)

func (M_p *Matrix) CloneColumn(n uint) (col []float64) {
	col = make([]float64, M_p.rowN)
	for m := uint(0); m < M_p.rowN; m++ {
		col[m] = M_p.Get(m, n)
	}
	return col
}
