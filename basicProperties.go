package matrix

// Public wrapped API
func (M_p *Matrix) Get(rowInd uint, colInd uint) float64 {
	return M_p.slice1D[rowInd*M_p.colN+colInd]
}
func (M_p *Matrix) Set(rowInd uint, colInd uint, value float64) {
	M_p.slice1D[rowInd*M_p.colN+colInd] = value
}
func (M_p *Matrix) SequentialInd(rowInd, colInd uint) uint {
	return rowInd*M_p.colN + colInd
}

// Field access
func (M_p *Matrix) RowNum() uint                               { return M_p.rowN }
func (M_p *Matrix) ColNum() uint                               { return M_p.colN }
func (M_p *Matrix) GetUnderlyingArr() (accesor []float64)      { return M_p.slice1D }
func (M_p *Matrix) GetWholeUnderlyingArr() (accesor []float64) { return M_p.slice1D[:cap(M_p.slice1D)] }
