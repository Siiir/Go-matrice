package matrix

func (M1_p *Matrix) IsSquare() bool {
	return M1_p.rowN == M1_p.colN
}

func (M1_p *Matrix) DoesProdExist(M2_p *Matrix) bool {
	return M1_p.colN == M2_p.rowN
}

func (M_p *Matrix) isEmpty() bool { //Performs full check
	return M_p.rowN == 0 || M_p.colN == 0
}
