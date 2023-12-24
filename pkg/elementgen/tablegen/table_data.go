package tablegen

import (
	"errors"
	"fmt"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/direction"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/fpdfpoint"
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/textconfig"
	"strings"
)

const MergedCellNamePrefix = "+"

type Column string

var ErrInvalidCellToMerge = errors.New("invalid cell to merge")

func numberToColumn(num int) Column {
	if num <= 0 {
		return ""
	}

	result := ""

	if f := (num - 1) / 26; f > 0 {
		result = fmt.Sprintf("%c", 'a'+f-1)
	}

	result += fmt.Sprintf("%c", 'a'+(num-1)%26)

	return Column(strings.ToUpper(result))
}

func (c Column) int() int {

	cLower := strings.ToLower(string(c))
	if len(cLower) == 0 {
		return 0
	}
	if len(c) == 1 {
		return int(cLower[0]) - 'a' + 1
	}
	first := int(cLower[0]) - 'a' + 1
	second := int(cLower[1]) - 'a' + 1

	return first*26 + second
}

type CellEdge struct {
	TopEdge    fpdfpoint.Line
	BottomEdge fpdfpoint.Line
	LeftEdge   fpdfpoint.Line
	RightEdge  fpdfpoint.Line
	CornerAndEdgeInterface
}

type CornerAndEdgeInterface interface {
	TopLeftCorner() fpdfpoint.Point
	TopRightCorner() fpdfpoint.Point
	BottomLeftCorner() fpdfpoint.Point
	BottomRightCorner() fpdfpoint.Point
	CardinalString() string
	HideEdge(edgeOption string) CellEdge
	Translation(x, y float64) CellEdge
	MaxLineWidth() int
}

func (c CellEdge) TopLeftCorner() fpdfpoint.Point {
	return c.TopEdge.A
}

func (c CellEdge) TopRightCorner() fpdfpoint.Point {
	return c.TopEdge.B
}

func (c CellEdge) BottomLeftCorner() fpdfpoint.Point {
	return c.BottomEdge.A
}

func (c CellEdge) BottomRightCorner() fpdfpoint.Point {
	return c.BottomEdge.B
}

func (c CellEdge) CardinalString() string {
	str := ""
	if c.TopEdge.IsShown() {
		str += direction.CardinalTop.String()
	}
	if c.BottomEdge.IsShown() {
		str += direction.CardinalBottom.String()
	}
	if c.LeftEdge.IsShown() {
		str += direction.CardinalLeft.String()
	}
	if c.RightEdge.IsShown() {
		str += direction.CardinalRight.String()
	}
	return str
}

func (c CellEdge) HideEdge(edgeOption string) CellEdge {
	if direction.IsCardinalTop(edgeOption) {
		c.TopEdge = c.TopEdge.Hide()
	}

	if direction.IsCardinalBottom(edgeOption) {
		c.BottomEdge = c.BottomEdge.Hide()
	}

	if direction.IsCardinalLeft(edgeOption) {
		c.LeftEdge = c.LeftEdge.Hide()
	}

	if direction.IsCardinalRight(edgeOption) {
		c.RightEdge = c.RightEdge.Hide()
	}

	return c
}

func (c CellEdge) Translation(x, y float64) CellEdge {
	c.TopEdge = c.TopEdge.Translation(x, y)
	c.BottomEdge = c.BottomEdge.Translation(x, y)
	c.LeftEdge = c.LeftEdge.Translation(x, y)
	c.RightEdge = c.RightEdge.Translation(x, y)
	return c
}

func (c CellEdge) MaxLineWidth() float64 {
	maxWidth := 0.0
	if maxWidth > c.TopEdge.Width {
		maxWidth = c.TopEdge.Width
	}
	if maxWidth > c.BottomEdge.Width {
		maxWidth = c.BottomEdge.Width
	}
	if maxWidth > c.LeftEdge.Width {
		maxWidth = c.LeftEdge.Width
	}
	if maxWidth > c.RightEdge.Width {
		maxWidth = c.RightEdge.Width
	}

	return maxWidth
}

func NewCellEdge(Ax, Ay, Cx, Cy, lineWidth float64) CellEdge {
	CornerA := fpdfpoint.Point{
		X: Ax,
		Y: Ay,
	}

	CornerB := fpdfpoint.Point{
		X: Cx,
		Y: Ay,
	}

	CornerC := fpdfpoint.Point{
		X: Cx,
		Y: Cy,
	}

	CornerD := fpdfpoint.Point{
		X: Ax,
		Y: Cy,
	}
	return CellEdge{
		TopEdge: fpdfpoint.Line{
			A:      CornerA,
			B:      CornerB,
			Width:  lineWidth,
			Status: fpdfpoint.Visible,
		},
		RightEdge: fpdfpoint.Line{
			A:      CornerB,
			B:      CornerC,
			Width:  lineWidth,
			Status: fpdfpoint.Visible,
		},
		BottomEdge: fpdfpoint.Line{
			A:      CornerD,
			B:      CornerC,
			Width:  lineWidth,
			Status: fpdfpoint.Visible,
		},
		LeftEdge: fpdfpoint.Line{
			A:      CornerA,
			B:      CornerD,
			Width:  lineWidth,
			Status: fpdfpoint.Visible,
		},
	}
}

type SingleCell struct {
	Column Column
	Row    int
	CellEdge
}

func NewCellAddress(col, row int, x, y, w, h, lineWidth float64) SingleCell {
	return SingleCell{
		Column:   numberToColumn(col),
		Row:      row,
		CellEdge: NewCellEdge(x, y, x+w, y+h, lineWidth),
	}
}

func (sc SingleCell) Name() string {
	return fmt.Sprintf("%v%v", sc.Column, sc.Row)
}

type MergedCell struct {
	Columns []Column
	Rows    []int
	CellEdge
}

func (mc MergedCell) Name() string {
	return fmt.Sprintf("%v%v%v", MergedCellNamePrefix, mc.Columns[0], mc.Rows[0])
}

func (mc MergedCell) NameAllSingle() []string {
	singles := make([]string, 0)
	for _, c := range mc.Columns {
		for _, r := range mc.Rows {
			singles = append(singles, fmt.Sprintf("%v%v", c, r))
		}
	}
	return singles
}

func NewMergedCell(CellTopLeft, CellBottomRight SingleCell) MergedCell {
	pt1A := CellTopLeft.TopLeftCorner()
	pt2A := CellTopLeft.BottomRightCorner()

	pt1B := CellBottomRight.TopLeftCorner()
	pt2B := CellBottomRight.BottomRightCorner()

	minX := pt1A.X
	if minX > pt1B.X {
		minX = pt1B.X
	}

	minY := pt1A.Y
	if minY > pt1B.Y {
		minY = pt1B.Y
	}

	maxX := pt2A.X
	if maxX < pt2B.X {
		maxX = pt2B.X
	}

	maxY := pt2A.Y
	if maxY < pt2B.Y {
		maxY = pt2B.Y
	}

	columns := make([]Column, 0)
	for i := CellTopLeft.Column.int(); i <= CellBottomRight.Column.int(); i++ {
		columns = append(columns, numberToColumn(i))
	}

	rows := make([]int, 0)
	for i := CellTopLeft.Row; i <= CellBottomRight.Row; i++ {
		rows = append(rows, i)
	}

	lineWidth := CellTopLeft.MaxLineWidth()
	brLineWidth := CellBottomRight.MaxLineWidth()

	if lineWidth < brLineWidth {
		lineWidth = brLineWidth
	}
	return MergedCell{
		columns,
		rows,
		NewCellEdge(minX, minY, maxX, maxY, lineWidth),
	}
}

func (mc MergedCell) ContainCell(c SingleCell) bool {
	mtl := mc.TopLeftCorner()
	mbr := mc.BottomRightCorner()

	ctl := c.TopLeftCorner()
	cbr := c.BottomRightCorner()

	return mtl.X <= ctl.X && mtl.Y <= ctl.Y && mbr.X >= cbr.X && mbr.Y >= cbr.Y
}

type TaggedUnionCell struct {
	*SingleCell
	*MergedCell
	textconfig.TextConfig
	CornerAndEdgeInterface
}

func NewSingleTaggedUnionCell(col, row int, x, y, w, h, lineWidth float64) TaggedUnionCell {
	single := NewCellAddress(col, row, x, y, w, h, lineWidth)
	return TaggedUnionCell{
		SingleCell: &single,
		MergedCell: nil,
	}
}

func (tuc TaggedUnionCell) Name() string {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.Name()
	}
	return tuc.SingleCell.Name()
}

func (tuc TaggedUnionCell) CardinalString() string {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.CardinalString()
	}
	return tuc.SingleCell.CardinalString()
}

func (tuc TaggedUnionCell) TopLeftCorner() fpdfpoint.Point {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.TopLeftCorner()
	}
	return tuc.SingleCell.TopLeftCorner()
}

func (tuc TaggedUnionCell) TopRightCorner() fpdfpoint.Point {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.TopRightCorner()
	}
	return tuc.SingleCell.TopRightCorner()
}

func (tuc TaggedUnionCell) BottomLeftCorner() fpdfpoint.Point {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.BottomLeftCorner()
	}
	return tuc.SingleCell.BottomLeftCorner()
}

func (tuc TaggedUnionCell) BottomRightCorner() fpdfpoint.Point {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.BottomRightCorner()
	}
	return tuc.SingleCell.BottomRightCorner()
}

func (tuc TaggedUnionCell) HideEdge(edgeOption string) TaggedUnionCell {
	if tuc.MergedCell != nil {
		tuc.MergedCell.CellEdge = tuc.MergedCell.HideEdge(edgeOption)
		return tuc
	}

	tuc.SingleCell.CellEdge = tuc.SingleCell.HideEdge(edgeOption)
	return tuc
}

func (tuc TaggedUnionCell) MaxLineWidth() float64 {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.MaxLineWidth()
	}

	return tuc.SingleCell.MaxLineWidth()
}

func (tuc TaggedUnionCell) WidthForFpdf() float64 {
	topLeft := tuc.TopLeftCorner()
	bottomRight := tuc.BottomRightCorner()
	return bottomRight.X - topLeft.X
}

func (tuc TaggedUnionCell) HeightForFpdf() float64 {
	topLeft := tuc.TopLeftCorner()
	bottomRight := tuc.BottomRightCorner()

	line := 1
	for _, r := range tuc.Text {
		if r == '\n' {
			line++
		}
	}

	return (bottomRight.Y - topLeft.Y) / float64(line)
}

type TableObject struct {
	ID      string
	CellMap map[string]TaggedUnionCell
}

//type CellMap map[string]TaggedUnionCell

func MakeTableObject(id string, XList, widths, YList, heights []float64, lineWidth float64) TableObject {
	cm := make(map[string]TaggedUnionCell)
	for i := 0; i < len(XList); i++ {
		for j := 0; j < len(YList); j++ {
			single := NewSingleTaggedUnionCell(i+1, j+1, XList[i], YList[j], widths[i], heights[j], lineWidth)
			cm[single.Name()] = single
		}
	}
	return TableObject{id, cm}
}

func (to TableObject) Merge(TopLeftName, BottomRightName string) error {
	topLeftCell := to.CellMap[TopLeftName]
	bottomRightCell := to.CellMap[BottomRightName]
	if topLeftCell.SingleCell == nil || bottomRightCell.SingleCell == nil {
		return ErrInvalidCellToMerge
	}

	mCell := NewMergedCell(*topLeftCell.SingleCell, *bottomRightCell.SingleCell)
	singleCellNames := mCell.NameAllSingle()

	to.CellMap[mCell.Name()] = TaggedUnionCell{
		SingleCell: nil,
		MergedCell: &mCell,
	}

	for _, n := range singleCellNames {
		delete(to.CellMap, n)
	}

	return nil
}

func (to TableObject) HideEdge(cellName, edgeToHide string) {
	var cellVal TaggedUnionCell

	for k, v := range to.CellMap {
		if k == cellName || k == MergedCellNamePrefix+cellName {
			cellVal = v
			break
		}
	}

	cellVal = cellVal.HideEdge(edgeToHide)

	to.CellMap[cellVal.Name()] = cellVal
}

func (to TableObject) Translation(x, y float64) TableObject {
	ncm := make(map[string]TaggedUnionCell)

	for k, tuc := range to.CellMap {
		if tuc.MergedCell != nil {
			tuc.MergedCell.CellEdge = tuc.MergedCell.Translation(x, y)
			ncm[k] = tuc
			continue
		}

		tuc.SingleCell.CellEdge = tuc.SingleCell.Translation(x, y)
		ncm[k] = tuc
	}

	return to
}

func (to TableObject) AddTextConfig(cellName string, textConfigToAdd textconfig.TextConfig) {
	var cellVal TaggedUnionCell

	for k, v := range to.CellMap {
		if k == cellName || k == MergedCellNamePrefix+cellName {
			cellVal = v
			break
		}
	}

	cellVal.TextConfig = textConfigToAdd
	to.CellMap[cellVal.Name()] = cellVal
}
