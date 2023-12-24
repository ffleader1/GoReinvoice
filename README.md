# GoReinvoice

This library help create invoice .pdf file from .json or .yaml file, eliminating the need for hard-code the templates.

Sample Code

```go
package main

import (
	"github.com/ffleader1/GoReinvoice/pkg/inputdata"
	"github.com/ffleader1/GoReinvoice/pkg/pdfgen"
	"log"
)

func main() {
	inputYaml, err := inputdata.ReadData("../../resource/yaml/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	pdfDataYaml := pdfgen.NewPdfData(inputYaml)
	pdfDataYaml.GenPdf(map[string]string{
		"ref1":            "occho 1",
		"ref2":            "occho 2",
		"tax_id":          "taxid1234",
		"invoice_suffix":  "edv",
		"payment_id":      "PMID1669692935128",
		"order_number":    "7164cebfd3f06a6322eeeb6d",
		"total_price_str": "10,000.00",
		"total_price":     "1000000",
		"name":            "Cuong Nguyen",
		"phone":           "0473894823",
		"comp_code":       "cmpcd789",
	}, "test_gen_invoice_yaml.pdf")
}
```
Text data can be filled with placeholders like ```{{user_name}}```.

After reading the data from the .yaml or .json file, pass a ```map[string]string``` to ```GenPdf```, along with the output file location to generate the pdf with place holder data replaced.

Currently, the following type of elements supported:
- text
- image (png, jpg)
- basic shape (line, ellipse)
- code (code128, QR)
- table
- copy (a clone of a set of other elements)


A sample .yaml input file with detailed comments is shown below.
```yaml
version: 2
resource: "../" # Set the relative location to search for file and font
elements:
  - id: 1-logo
    type: image # Type Image. See Files below for more details
    x: 17 # Top left image X coordinate
    y: 12 # Top left image Y coordinate
    width: 300 # Image width
    height: 80 # Image height
    scale:
      - 0.12 # Scale width
      - 0.12 # Scale height
  - id: 2-text
    type: text # Type Text
    x: 17 # Top left text block X coordinate
    y: 27 # Top left text block Y coordinate
    width: 100 # Text block width
    height: 25 # Text block height
    text: "Bill Payment for {{user_name}}" # Text to fill
    fontSize: 14 # Font size
    fontFamily: 1 # Font family. See Fonts below for more details
    textAlign: left # Horizontal alignment in text block. Available values: left, center, right. Default: Center
    verticalAlign: top # Vertical alignment in text block. Available values: top, middle, bottom. Default: Middle
  - id: 3-ellipse
    type: ellipse  # Type Ellipse
    x: 19  # Top left ellipse X coordinate
    y: 77  # Top left ellipse Y coordinate
    width: 4 # Ellipse horizontal diameter
    height: 4 # Ellipse vertical diameter
    angle: 0 # Ellipse rotated angle
    strokeWidth: 2 # Stroke width to draw the ellipse
  - id: 4-table
    type: table # Type Table. See Tables below for more information
    x: 19 # Top left table X coordinate
    y: 83 # Top left table Y coordinate
    width: 180 # Table width
    height: 30 # Table height
  - id: 5-code128
    type: code128  # Type Code128
    x: 110 # Top left code image X coordinate
    y: 117 # Top left code image Y coordinate
    width: 89 # Code image width
    height: 10 # Code image height
    text: "|{{tax_id}}\n{{invoice_suffix}}\n{{total_price}}" # Code data
  - id: 6-qrcode
    type: qrcode  # Type QR Code
    x: 110 # Top left code image X coordinate
    y: 117 # Top left code image Y coordinate
    width: 60 # Code image width
    height: 60 # Code image height
    text: "|{{tax_id}}\n{{invoice_suffix}}\n{{total_price}}" # Code data
  - id: 7-line
    type: line # Type Line
    x: 18 # Default top left line X coordinate
    y: 143 # Default top left line Y coordinate
    strokeWidth: 2 # Stroke width to draw the line
    points: # An array with 2 elements, defining the start and end point of the line.
      # Values are added to the x and y above.
      - - 0 # Add this value to X to get the starting X coordinate of the line
        - 0 # Add this value to Y to get the starting Y coordinate of the line
      - - 180  # Add this value to X to get the ending X coordinate of the line
        - 0 # Add this value to Y to get the ending Y coordinate of the line
  - id: 8-copy
    type: copy # Type Copy. See Copies below for more information
    x: 0 # Add this value to every X value of the copied elements
    y: 135 # Add this value to every Y value of the copied elements
tables: # A map, with the key is the element id and the value is table configuration
  4-table:
    columnRatio: # Width ratio of each column. Must add up to 1
      - 0.3
      - 0.2
      - 0.5
    rowRatio: # Height ratio of each row. Must add up to 1
      - 0.4
      - 0.3
      - 0.3
    mergeCell: # A map with key is the top left cell and value is the bottom right cell of the cell block getting merged
      # For example, A2:B3 means 4 cells A2, A3, B2, B3 is merged into cell A2
      A2: B3
    hiddenEdge: # A map with key is the cell and value is its hidden borders
      # For example, A3: LB means its left (L) and bottom (B) borders are hidden.
      # Possible values are: L for left, R for right, T for top and B for bottom.
      # For merged cell, the original top left cell is treated as the merged cell name
      A2: LB
    cellText:
      A1:
        text: Name # Text to fill in the cell
        fontFamily: 1 # Font size
        fontSize: 14 # Font family. See Fonts below for more details
        textAlign: left # Horizontal alignment in cell block. Available values: left, center, right. Default: Center
        verticalAlign: top # Vertical alignment in cell block. Available values: top, middle, bottom. Default: Middle
      C3:
        text: "{{total}}"
        fontFamily: 1
        fontSize: 14
files: # Map of image file element id and file config value
  1-logo:
    dataURL: image/go_rust.png # File path relative to Resource above
fonts: # List of config values
  - familyName: sarabun # Font family name
    style: B # Specify font style. "B" (bold), "I" (italic), "U" (underscore), "S" (strike-out) or any combination.
    dataURL: font/sarabun_bold.ttf # Font file path relative to Resource above
  - familyName: sarabun
    style: ""
    dataURL: font/sarabun_regular.ttf
copies: # Map of copy element id and file config value
  8-copy: # List of element ids to copy
    - 1-logo
    - 2-text
    - 3-ellipse
paper: # Paper config. Currently using the size A4, measurement millimeter (mm) and orientation Portrait (P)
  size: A4
  unit: mm
  orientation: P
```