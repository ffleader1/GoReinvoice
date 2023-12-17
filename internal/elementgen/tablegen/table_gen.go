package tablegen

import (
	"GoReinvoice/internal/inputdata"
	"errors"
	"math"
)

var ErrInvalidTableSize = errors.New("invalid table size")

type TableLine struct {
	From        int
	To          int
	Strokewidth int
}

func GenerateTableLine(x, y, width, height int, tableExtra inputdata.Table) (CellMap, error) {
	sumCol := .0
	sumRow := .0

	for _, c := range tableExtra.ColumnRatio {
		sumCol += c
	}
	if sumCol != 1 {
		return nil, ErrInvalidTableSize
	}

	for _, r := range tableExtra.RowRatio {
		sumRow += r
	}
	if sumRow != 1 {
		return nil, ErrInvalidTableSize
	}

	xList, widths := genLocationAndLength(x, width, tableExtra.ColumnRatio)

	yList, heights := genLocationAndLength(y, height, tableExtra.RowRatio)

	cellMap := MakeCellMap(xList, widths, yList, heights)

	for k, v := range tableExtra.MergeCell {
		if err := cellMap.Merge(k, v); err != nil {
			return nil, err
		}
	}

	return cellMap, nil
}

func genLocationAndLength(s, l int, ratio []float64) ([]int, []int) {
	lengths := make([]int, 0)
	for _, r := range ratio {
		lengths = append(lengths, int(math.Round(float64(l)*r)))
	}
	starts := []int{s}
	for idx, ll := range lengths {
		if idx < len(lengths)-1 {
			starts = append(starts, starts[idx]+ll)
		}
	}
	return starts, lengths
}