package goreinvoice

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
