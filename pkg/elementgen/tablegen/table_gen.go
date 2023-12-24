package tablegen

import (
	"fmt"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/customerr"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/direction"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/textconfig"
	"github.com/ffleader1/GoReinvoice/pkg/inputdata"
	"math"
)

type TableLine struct {
	From        int
	To          int
	Strokewidth int
}

func GenerateTableObject(id string, x, y, width, height, lineWidth float64, tableExtra inputdata.Table) (TableObject, error) {
	sumCol := .0
	sumRow := .0

	for _, c := range tableExtra.ColumnRatio {
		sumCol += c
	}
	if sumCol != 1 {
		return TableObject{}, fmt.Errorf("%w: Sum Colum Ratio Incorrect: %v, but expect 1", customerr.ErrInvalidTableSize, sumCol)
	}

	for _, r := range tableExtra.RowRatio {
		sumRow += r
	}
	if sumRow != 1 {
		return TableObject{}, fmt.Errorf("%w: Sum Row Ratio Incorrect: %v, but expect 1", customerr.ErrInvalidTableSize, sumRow)
	}

	xList, widths := genLocationAndLength(x, width, tableExtra.ColumnRatio)

	yList, heights := genLocationAndLength(y, height, tableExtra.RowRatio)

	cellMap := MakeTableObject(id, xList, widths, yList, heights, lineWidth)

	for k, v := range tableExtra.MergeCell {
		if err := cellMap.Merge(k, v); err != nil {
			return TableObject{}, err
		}
	}

	for k, v := range tableExtra.HiddenEdge {
		cellMap.HideEdge(k, v)
	}

	for k, v := range tableExtra.CellText {
		tf := textconfig.TextConfig{
			Text:                v.Text,
			FontSize:            v.FontSize,
			FontFamily:          v.FontFamily,
			HorizontalAlignment: direction.ToHorizontalAlignment(v.TextAlign),
			VerticalAlignment:   direction.ToVerticalAlignment(v.VerticalAlign),
		}
		cellMap.AddTextConfig(k, tf)
	}

	return cellMap, nil
}

func genLocationAndLength(s, l float64, ratio []float64) ([]float64, []float64) {
	lengths := make([]float64, 0)
	for _, r := range ratio {
		lengths = append(lengths, math.Round(l*r))
	}
	starts := []float64{s}
	for idx, ll := range lengths {
		if idx < len(lengths)-1 {
			starts = append(starts, starts[idx]+ll)
		}
	}
	return starts, lengths
}
