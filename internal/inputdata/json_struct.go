package inputdata

type PdfInput struct {
	Type       string           `json:"type"`
	Version    int              `json:"version"`
	Source     string           `json:"source"`
	SourceSize string           `json:"source_size"`
	Rotation   string           `json:"rotation"`
	Elements   []Element        `json:"elements"`
	AppState   AppState         `json:"appState"`
	Files      map[string]File  `json:"files"`
	Tables     map[string]Table `json:"tables"`
	Fonts      []Font           `json:"fonts"`
	Paper      Paper            `json:"paper"`
}

type Element struct {
	ID              string        `json:"id"`
	Type            string        `json:"type"`
	X               int           `json:"x"`
	Y               int           `json:"y"`
	Width           float64       `json:"width"`
	Height          float64       `json:"height"`
	Angle           int           `json:"angle"`
	StrokeColor     string        `json:"strokeColor"`
	BackgroundColor string        `json:"backgroundColor"`
	FillStyle       string        `json:"fillStyle"`
	StrokeWidth     int           `json:"strokeWidth"`
	StrokeStyle     string        `json:"strokeStyle"`
	Roughness       int           `json:"roughness"`
	Opacity         int           `json:"opacity"`
	GroupIds        []interface{} `json:"groupIds"`
	FrameID         interface{}   `json:"frameId"`
	Roundness       interface{}   `json:"roundness"`
	Seed            int           `json:"seed"`
	Version         int           `json:"version"`
	VersionNonce    int           `json:"versionNonce"`
	IsDeleted       bool          `json:"isDeleted"`
	BoundElements   interface{}   `json:"boundElements"`
	Updated         int64         `json:"updated"`
	Link            interface{}   `json:"link"`
	Locked          bool          `json:"locked"`
	Text            string        `json:"text,omitempty"`
	FontSize        int           `json:"fontSize,omitempty"`
	FontFamily      int           `json:"fontFamily,omitempty"`
	TextAlign       string        `json:"textAlign,omitempty"`
	VerticalAlign   string        `json:"verticalAlign,omitempty"`
	Baseline        int           `json:"baseline,omitempty"`
	ContainerID     string        `json:"containerId,omitempty"`
	OriginalText    string        `json:"originalText,omitempty"`
	LineHeight      float64       `json:"lineHeight,omitempty"`
	Status          string        `json:"status,omitempty"`
	FileID          string        `json:"fileId,omitempty"`
	Scale           []int         `json:"scale,omitempty"`
}

type File struct {
	MimeType      string `json:"mimeType"`
	ID            string `json:"id"`
	DataURL       string `json:"dataURL"`
	Created       int64  `json:"created"`
	LastRetrieved int64  `json:"lastRetrieved"`
}

type AppState struct {
	GridSize            interface{} `json:"gridSize"`
	ViewBackgroundColor string      `json:"viewBackgroundColor"`
}

type Table struct {
	ColumnRatio []float64         `json:"columnRatio"`
	RowRatio    []float64         `json:"rowRatio"`
	MergeCell   map[string]string `json:"mergeCell"`
	HiddenEdge  map[string]string `json:"hiddenEdge"`
	CellText    map[string]string `json:"cellText"`
}

type Font struct {
	FamilyName string `json:"familyName"`
	Style      string `json:"style"`
	DataURL    string `json:"dataURL"`
}

type Paper struct {
	Size        string `json:"size"`
	Unit        string `json:"unit"`
	Orientation string `json:"orientation"`
}
