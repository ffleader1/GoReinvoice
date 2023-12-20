package tablegen

import (
	"GoReinvoice/internal/customtypes/direction"
	"GoReinvoice/internal/elementgen/resuable/pointgen"
	"errors"
	"fmt"

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
	TopEdge    pointgen.Line
	BottomEdge pointgen.Line
	LeftEdge   pointgen.Line
	RightEdge  pointgen.Line
	CornerAndEdgeInterface
}

type CornerAndEdgeInterface interface {
	TopLeftCorner() pointgen.Point
	TopRightCorner() pointgen.Point
	BottomLeftCorner() pointgen.Point
	BottomRightCorner() pointgen.Point
	CardinalString() string
	HideEdge(edgeOption string)
	MaxLineWidth() int
}

func (c *CellEdge) TopLeftCorner() pointgen.Point {
	return c.TopEdge.A
}

func (c *CellEdge) TopRightCorner() pointgen.Point {
	return c.TopEdge.B
}

func (c *CellEdge) BottomLeftCorner() pointgen.Point {
	return c.BottomEdge.A
}

func (c *CellEdge) BottomRightCorner() pointgen.Point {
	return c.BottomEdge.B
}

func (c *CellEdge) CardinalString() string {
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

func (c *CellEdge) HideEdge(edgeOption string) {
	if direction.IsCardinalTop(edgeOption) {
		c.TopEdge.Hide()
	}

	if direction.IsCardinalBottom(edgeOption) {
		c.BottomEdge.Hide()
	}

	if direction.IsCardinalLeft(edgeOption) {
		c.LeftEdge.Hide()
	}

	if direction.IsCardinalRight(edgeOption) {
		c.RightEdge.Hide()
	}
}

func (c *CellEdge) MaxLineWidth() int {
	maxWidth := 0
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

func NewCellEdge(Ax, Ay, Cx, Cy, lineWidth int) CellEdge {
	CornerA := pointgen.Point{
		X: Ax,
		Y: Ay,
	}

	CornerB := pointgen.Point{
		X: Cx,
		Y: Ay,
	}

	CornerC := pointgen.Point{
		X: Cx,
		Y: Cy,
	}

	CornerD := pointgen.Point{
		X: Ax,
		Y: Cy,
	}
	return CellEdge{
		TopEdge: pointgen.Line{
			A:      CornerA,
			B:      CornerB,
			Width:  lineWidth,
			Status: pointgen.Visible,
		},
		RightEdge: pointgen.Line{
			A:      CornerB,
			B:      CornerC,
			Width:  lineWidth,
			Status: pointgen.Visible,
		},
		BottomEdge: pointgen.Line{
			A:      CornerD,
			B:      CornerC,
			Width:  lineWidth,
			Status: pointgen.Visible,
		},
		LeftEdge: pointgen.Line{
			A:      CornerA,
			B:      CornerD,
			Width:  lineWidth,
			Status: pointgen.Visible,
		},
	}
}

type SingleCell struct {
	Column Column
	Row    int
	CellEdge
}

func NewCellAddress(col, row, x, y, w, h, lineWidth int) SingleCell {
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
	Text string
	CornerAndEdgeInterface
}

func NewSingleTaggedUnionCell(col, row, x, y, w, h, lineWidth int) TaggedUnionCell {
	single := NewCellAddress(col, row, x, y, w, h, lineWidth)
	return TaggedUnionCell{
		SingleCell: &single,
		MergedCell: nil,
	}
}

func (tuc *TaggedUnionCell) Name() string {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.Name()
	}
	return tuc.SingleCell.Name()
}

func (tuc *TaggedUnionCell) CardinalString() string {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.CardinalString()
	}
	return tuc.SingleCell.CardinalString()
}

func (tuc *TaggedUnionCell) TopLeftCorner() pointgen.Point {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.TopLeftCorner()
	}
	return tuc.SingleCell.TopLeftCorner()
}

func (tuc *TaggedUnionCell) TopRightCorner() pointgen.Point {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.TopRightCorner()
	}
	return tuc.SingleCell.TopRightCorner()
}

func (tuc *TaggedUnionCell) BottomLeftCorner() pointgen.Point {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.BottomLeftCorner()
	}
	return tuc.SingleCell.BottomLeftCorner()
}

func (tuc *TaggedUnionCell) BottomRightCorner() pointgen.Point {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.BottomRightCorner()
	}
	return tuc.SingleCell.BottomRightCorner()
}

func (tuc *TaggedUnionCell) HideEdge(edgeOption string) {
	if tuc.MergedCell != nil {
		tuc.MergedCell.HideEdge(edgeOption)
		return
	}

	tuc.SingleCell.HideEdge(edgeOption)
}

func (tuc *TaggedUnionCell) MaxLineWidth() int {
	if tuc.MergedCell != nil {
		return tuc.MergedCell.MaxLineWidth()
	}

	return tuc.SingleCell.MaxLineWidth()
}

func (tuc *TaggedUnionCell) WidthForFpdf() float64 {
	topLeft := tuc.TopLeftCorner()
	bottomRight := tuc.BottomRightCorner()
	return float64(bottomRight.X - topLeft.X)
}

func (tuc *TaggedUnionCell) HeightForFpdf() float64 {
	topLeft := tuc.TopLeftCorner()
	bottomRight := tuc.BottomRightCorner()

	line := 1
	for _, r := range tuc.Text {
		if r == '\n' {
			line++
		}
	}

	return float64(bottomRight.Y-topLeft.Y) / float64(line)
}

type CellMap map[string]TaggedUnionCell

func MakeCellMap(XList, widths, YList, heights []int, lineWidth int) CellMap {
	cm := make(map[string]TaggedUnionCell)
	for i := 0; i < len(XList); i++ {
		for j := 0; j < len(YList); j++ {
			single := NewSingleTaggedUnionCell(i+1, j+1, XList[i], YList[j], widths[i], heights[j], lineWidth)
			cm[single.Name()] = single
		}
	}
	return cm
}

func (cm CellMap) Merge(TopLeftName, BottomRightName string) error {
	topLeftCell := cm[TopLeftName]
	bottomRightCell := cm[BottomRightName]
	if topLeftCell.SingleCell == nil || bottomRightCell.SingleCell == nil {
		return ErrInvalidCellToMerge
	}

	mCell := NewMergedCell(*topLeftCell.SingleCell, *bottomRightCell.SingleCell)
	singleCellNames := mCell.NameAllSingle()

	cm[mCell.Name()] = TaggedUnionCell{
		SingleCell: nil,
		MergedCell: &mCell,
	}

	for _, n := range singleCellNames {
		delete(cm, n)
	}

	return nil
}

func (cm CellMap) HideEdge(cellName, edgeToHide string) {
	var cellVal TaggedUnionCell

	for k, v := range cm {
		if k == cellName || k == MergedCellNamePrefix+cellName {
			cellVal = v
			break
		}
	}

	cellVal.HideEdge(edgeToHide)

	cm[cellVal.Name()] = cellVal
}

func (cm CellMap) AddText(cellName, textToAdd string) {
	var cellVal TaggedUnionCell

	for k, v := range cm {
		if k == cellName || k == MergedCellNamePrefix+cellName {
			cellVal = v
			break
		}
	}

	cellVal.Text = strings.TrimSpace(textToAdd)
	cm[cellVal.Name()] = cellVal
}
