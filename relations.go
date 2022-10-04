package matrix

func (M1_p *Matrix) EqShape(M2_p *Matrix) bool {
	return M1_p.rowN == M2_p.rowN && M1_p.colN == M2_p.colN
}
func (M1_p *Matrix) CanContain(M2_p *Matrix) bool {
	return M1_p.rowN >= M2_p.rowN && M1_p.colN >= M2_p.colN
}
