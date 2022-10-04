package matrix

import "fmt"

func NewZeroed(rowN, colN uint) *Matrix {
	M := &Matrix{
		rowN:    rowN,
		colN:    colN,
		slice1D: make([]float64, rowN*colN),
	}
	return M
}

func _tryNewSharing1DSlice(Sl []float64, rowN, colN uint) *Matrix {
	l := rowN * colN
	if l > uint(cap(Sl)) {
		return nil
	}
	return &Matrix{
		slice1D: Sl[:l],
		rowN:    rowN,
		colN:    colN,
	}
}
func TryNewSharing1DSlice(Sl []float64, rowN, colN uint) (M_p *Matrix, e error) {
	M_p = _tryNewSharing1DSlice(Sl, rowN, colN)
	if M_p == nil {
		e = fmt.Errorf(
			"given slice has to small capacity to contain matrix %v×%v."+
				"Hint: %v·%v = %v < %v = slice_capacity",
			rowN, colN,
			rowN, colN,
			rowN*colN,
			cap(Sl),
		)
		return
	}
	return
}
func NewFrom1DSlice(Sl []float64, rowN, colN uint) *Matrix {
	M_p := &Matrix{
		slice1D: make([]float64, rowN*colN),
		rowN:    rowN,
		colN:    colN,
	}
	copy(M_p.slice1D, Sl)
	return M_p
}

func NewFrom2DSlice(S [][]float64) (M *Matrix, e error) {
	// Validation
	e = Validate2DSliceAsMatrix(S)
	if e != nil {
		return
	}

	// Setting dimensions
	M.rowN = uint(len(S))
	if M.rowN == 0 {
		M.colN = 0
	} else {
		M.colN = uint(len(S[0]))
	}
	M.slice1D = make([]float64, M.rowN*M.colN)

	// Copying values
	for m := uint(0); m < M.rowN; m++ {
		copy(M.slice1D[(m*M.colN):((m+1)*M.colN)], S[m])
	}

	return
}
