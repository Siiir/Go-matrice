package matrix

import "github.com/Siiir/vector"

func (Sm_p *Submatrix) GetRow(rowInd uint) (accesor []float64) {
	return Sm_p.slice2D[rowInd]
}

func (Sm_p *Submatrix) CloneRow(rowInd uint) (accesor []float64) {
	return vector.Clone(Sm_p.slice2D[rowInd])
}
