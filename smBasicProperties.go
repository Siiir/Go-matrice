package matrix

func (Sm_p *Submatrix) RowNum() uint {
	return uint(len(Sm_p.slice2D))
}
func (Sm_p *Submatrix) ColNum() uint {
	return Sm_p.colN
}

func (Sm_p *Submatrix) Get(rowInd uint, colInd uint) float64 {
	return Sm_p.slice2D[rowInd][colInd]
}

func (Sm_p *Submatrix) Set(rowInd uint, colInd uint, value float64) {
	Sm_p.slice2D[rowInd][colInd] = value
}
func (Sm_p *Submatrix) SequentialInd(rowInd, colInd uint) uint {
	return rowInd*Sm_p.colN + colInd
}
