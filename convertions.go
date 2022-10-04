package matrix

func (M_p *Matrix) DescribeWith2DSlice() (Sl_2D [][]float64) {
	for m := uint(0); m < M_p.rowN; m++ {
		Sl_2D[m] = M_p.GetRow(m)
	}
	return
}

func (M_p *Matrix) Into2DSlice() (Sl_2D [][]float64) {
	for m := uint(0); m < M_p.rowN; m++ {
		Sl_2D[m] = M_p.CloneRow(m)
	}
	return
}
