package inputdata

type PdfInput struct {
	Type       string           `json:"type" yaml:"type"`
	Version    int              `json:"version" yaml:"version"`
	Source     string           `json:"source" yaml:"source"`
	SourceSize string           `json:"source_size" yaml:"source_size"`
	Rotation   string           `json:"rotation" yaml:"rotation"`
	Resource   string           `json:"resource" yaml:"resource"`
	Elements   []Element        `json:"elements" yaml:"elements"`
	AppState   AppState         `json:"appState" yaml:"appState"`
	Files      map[string]File  `json:"files" yaml:"files"`
	Tables     map[string]Table `json:"tables" yaml:"tables"`
	Fonts      []Font           `json:"fonts" yaml:"fonts"`
	Paper      Paper            `json:"paper" yaml:"paper"`
}

type Element struct {
	ID              string        `json:"id" yaml:"id"`
	Type            string        `json:"type" yaml:"type"`
	X               int           `json:"x" yaml:"x"`
	Y               int           `json:"y" yaml:"y"`
	Width           float64       `json:"width" yaml:"width"`
	Height          float64       `json:"height" yaml:"height"`
	Angle           float64       `json:"angle" yaml:"angle"`
	StrokeColor     string        `json:"strokeColor" yaml:"strokeColor"`
	BackgroundColor string        `json:"backgroundColor" yaml:"backgroundColor"`
	FillStyle       string        `json:"fillStyle" yaml:"fillStyle"`
	StrokeWidth     int           `json:"strokeWidth" yaml:"strokeWidth"`
	StrokeStyle     string        `json:"strokeStyle" yaml:"strokeStyle"`
	Roughness       int           `json:"roughness" yaml:"roughness"`
	Opacity         int           `json:"opacity" yaml:"opacity"`
	GroupIds        []interface{} `json:"groupIds" yaml:"groupIds"`
	FrameID         interface{}   `json:"frameId" yaml:"frameId"`
	Roundness       interface{}   `json:"roundness" yaml:"roundness"`
	Seed            int           `json:"seed" yaml:"seed"`
	Version         int           `json:"version" yaml:"version"`
	VersionNonce    int           `json:"versionNonce" yaml:"versionNonce"`
	IsDeleted       bool          `json:"isDeleted" yaml:"isDeleted"`
	BoundElements   interface{}   `json:"boundElements" yaml:"boundElements"`
	Updated         int64         `json:"updated" yaml:"updated"`
	Link            interface{}   `json:"link" yaml:"link"`
	Locked          bool          `json:"locked" yaml:"locked"`
	Text            string        `json:"text,omitempty" yaml:"text"`
	FontSize        int           `json:"fontSize,omitempty" yaml:"fontSize,omitempty"`
	FontFamily      int           `json:"fontFamily,omitempty" yaml:"fontFamily,omitempty"`
	TextAlign       string        `json:"textAlign,omitempty" yaml:"textAlign,omitempty"`
	VerticalAlign   string        `json:"verticalAlign,omitempty" yaml:"verticalAlign,omitempty"`
	Baseline        int           `json:"baseline,omitempty" yaml:"baseline,omitempty"`
	ContainerID     string        `json:"containerId,omitempty" yaml:"containerId,omitempty"`
	OriginalText    string        `json:"originalText,omitempty" yaml:"originalText,omitempty"`
	LineHeight      float64       `json:"lineHeight,omitempty" yaml:"lineHeight,omitempty"`
	Status          string        `json:"status,omitempty" yaml:"status,omitempty"`
	FileID          string        `json:"fileId,omitempty" yaml:"fileId,omitempty"`
	Scale           []float64     `json:"scale,omitempty" yaml:"scale,omitempty"`
	Point           [][]int       `json:"point,omitempty" yaml:"point,omitempty"`
}

type File struct {
	MimeType      string `json:"mimeType" yaml:"mimeType"`
	ID            string `json:"id" yaml:"id"`
	DataURL       string `json:"dataURL" yaml:"dataURL"`
	Created       int64  `json:"created" yaml:"created"`
	LastRetrieved int64  `json:"lastRetrieved" yaml:"lastRetrieved"`
}

type AppState struct {
	GridSize            interface{} `json:"gridSize" yaml:"gridSize"`
	ViewBackgroundColor string      `json:"viewBackgroundColor" yaml:"iewBackgroundColor"`
}

type Table struct {
	ColumnRatio []float64         `json:"columnRatio" yaml:"columnRatio"`
	RowRatio    []float64         `json:"rowRatio" yaml:"rowRatio"`
	MergeCell   map[string]string `json:"mergeCell" yaml:"mergeCell"`
	HiddenEdge  map[string]string `json:"hiddenEdge" yaml:"hiddenEdge"`
	CellText    map[string]string `json:"cellText" yaml:"cellText"`
}

type Font struct {
	FamilyName string `json:"familyName" yaml:"familyName"`
	Style      string `json:"style" yaml:"style"`
	DataURL    string `json:"dataURL" yaml:"dataURL"`
}

type Paper struct {
	Size        string `json:"size" yaml:"size"`
	Unit        string `json:"unit" yaml:"unit"`
	Orientation string `json:"orientation" yaml:"orientation"`
}
