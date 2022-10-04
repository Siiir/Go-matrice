package matrix

//Matrix-Matrix
func (M1_p *Matrix) UncheckedAddOrSub(M2_p *Matrix, doSub bool) {
	if doSub {
		for i, v := range M2_p.slice1D {
			M1_p.slice1D[i] -= v
		}
	} else {
		for i, v := range M2_p.slice1D {
			M1_p.slice1D[i] += v
		}
	}
}

func (M1_p *Matrix) TryAddOrSub(M2_p *Matrix, doSub bool) (impossible bool) {
	// The check.
	if !M1_p.EqShape(M2_p) {
		impossible = true
		return
	}
	// The algorithm
	M1_p.UncheckedAddOrSub(M2_p, doSub)

	return
}
func (M1_p *Matrix) TryAdd(M2_p *Matrix) (impossible bool) {
	return M1_p.TryAddOrSub(M2_p, false)
}
func (M1_p *Matrix) TrySub(M2_p *Matrix) (impossible bool) {
	return M1_p.TryAddOrSub(M2_p, true)
}

//In the air
//Matrix-Matrix
func (M1_p *Matrix) UncheckedSumOrDiff(M2_p *Matrix, doDiff bool) *Matrix {
	M_p := M1_p.Clone()
	M_p.UncheckedAddOrSub(M2_p, doDiff)
	return M_p
}
func (M1_p *Matrix) TrySumOrDiff(M2_p *Matrix, doDiff bool) (M_p *Matrix, impossible bool) {
	// The check.
	if !M1_p.EqShape(M2_p) {
		impossible = true
		return
	}
	// The algorithm
	M_p = M1_p.UncheckedSumOrDiff(M2_p, doDiff)
	return
}
func (M1_p *Matrix) TrySum(M2_p *Matrix) (M_p *Matrix, impossible bool) {
	return M1_p.TrySumOrDiff(M2_p, false)
}
func (M1_p *Matrix) TryDiff(M2_p *Matrix) (M_p *Matrix, impossible bool) {
	return M1_p.TrySumOrDiff(M2_p, true)
}
