package matrix

func (M_p *Matrix) ifEmptyThanMake0x0() {
	if M_p.isEmpty() {
		M_p.rowN = 0
		M_p.colN = 0
	}
}
