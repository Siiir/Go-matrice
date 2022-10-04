package matrix

//Rows
func (M_p *Matrix) GetRow(rowInd uint) (accesor []float64) {
	return M_p.slice1D[rowInd*M_p.colN : (rowInd+1)*M_p.colN]
}

func (M_p *Matrix) CloneRow(rowInd uint) (clone []float64) {
	clone = make([]float64, M_p.colN)
	copy(clone, M_p.GetRow(rowInd))
	return
}

func (M_p *Matrix) SetRow(rowInd uint, newRow []float64) {
	copy(M_p.GetRow(rowInd), newRow)
}

func (M_p *Matrix) UncheckedDropFirstNRows(n uint) {
	M_p.slice1D = M_p.slice1D[n*M_p.colN:]
}
func (M_p *Matrix) DropFirstNRows(n uint) (panicMemoryErr interface{}) {
	defer func() {
		panicMemoryErr = recover()
	}()
	M_p.slice1D = M_p.slice1D[n*M_p.colN:]
	return
}

func (M_p *Matrix) UncheckedSetMemorySpanByRows(newRowNum uint) {
	M_p.slice1D = M_p.slice1D[0 : newRowNum*M_p.colN]
	//no panic => success =>
	M_p.rowN = newRowNum
}
func (M_p *Matrix) SetMemorySpanByRows(newRowNum uint) (panicMemoryErr interface{}) {
	defer func() {
		panicMemoryErr = recover()
	}()
	M_p.UncheckedSetMemorySpanByRows(newRowNum)

	return
}
func (M_p *Matrix) HideLastNRows(n uint) (notEnoughRowsInMatrix bool) {
	newRowN := int(M_p.rowN) - int(n)
	notEnoughRowsInMatrix = newRowN < 0
	if notEnoughRowsInMatrix {
		return
	}

	M_p.UncheckedSetMemorySpanByRows(uint(newRowN))
	return
}
func (M_p *Matrix) UnhideLastNRows(n uint) (panicMemoryErr interface{}) {
	return M_p.SetMemorySpanByRows(M_p.rowN + n)
}
