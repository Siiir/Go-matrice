package matrix

func (M *Matrix) AddScalar(scalar float64) {
	for i, stop := 0, len(M.slice1D); i < stop; i++ {
		M.slice1D[i] += scalar
	}
}
func (M *Matrix) MultByScalar(scalar float64) {
	for i, stop := 0, len(M.slice1D); i < stop; i++ {
		M.slice1D[i] *= scalar
	}
}
func (M_p *Matrix) MultRowByScalar(rowInd uint, scalar float64) {
	row := M_p.GetRow(rowInd)
	for n := range row {
		row[n] *= scalar
	}
}
func (M1_p *Matrix) ProdWithScalar(scalar float64) *Matrix {
	M_p := M1_p.Clone()
	M_p.MultByScalar(scalar)
	return M_p
}
