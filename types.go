package matrix

// This needs documentation.
type Matrix struct {
	slice1D []float64
	rowN    uint
	colN    uint
}

type Submatrix struct {
	// Descriptor over Matrix
	colN    uint
	slice2D [][]float64
}

type Matrice interface {
	// Interface for all types of matrices {Matrix, Submatrix, ...}
}
