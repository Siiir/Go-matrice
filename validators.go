package matrix

import "fmt"

func Validate2DSliceAsMatrix(S [][]float64) error {
	h := len(S)
	if h == 0 {
		return nil
	} else {
		w := len(S[0])
		for i := 1; i < h; i++ {
			l := len(S[i])
			if w != l {
				return fmt.Errorf("rows [0:%d) have len=%d, whereas row %d has len=%d", i, w, i, l)
			}
		}
	}
	return nil
}
