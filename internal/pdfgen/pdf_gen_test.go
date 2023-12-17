package pdfgen

import (
	"GoReinvoice/internal/inputdata"
	"fmt"
	"testing"
)

func TestGenPdf(t *testing.T) {
	input := inputdata.PdfInput{
		Type:       "table",
		Version:    0,
		Source:     "",
		SourceSize: "",
		Rotation:   "",
		Elements: []inputdata.Element{
			{
				ID:     "elm1",
				Type:   "table",
				X:      20,
				Y:      30,
				Width:  180,
				Height: 100,
			},
		},
		AppState: inputdata.AppState{},
		Files:    nil,
		Tables: map[string]inputdata.Table{
			"elm1": {
				ColumnRatio: []float64{0.2, 0.2, 0.2, 0.2, 0.2},
				RowRatio:    []float64{0.25, 0.25, 0.25, 0.25},
				MergeCell:   map[string]string{"A3": "C4", "D1": "E1"},
				HiddenEdge:  map[string]string{"A3": "LB", "E4": "TR"},
			},
		},
	}
	fmt.Println("Start write file")
	GenPdf(input)
}