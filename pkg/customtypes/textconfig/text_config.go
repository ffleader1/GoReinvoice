package textconfig

import (
	"github.com/ffleader1/GoReinvoice/pkg/customtypes/direction"
	"strings"
)

type TextConfig struct {
	Text                string
	FontSize            int
	FontFamily          int
	HorizontalAlignment direction.HorizontalAlignment
	VerticalAlignment   direction.VerticalAlignment
}

func (tc TextConfig) AlignmentString() string {
	return tc.HorizontalAlignment.String() + tc.VerticalAlignment.String()
}

func (tc TextConfig) TextWithPlaceholder(placeHolderMap map[string]string) string {
	for k, v := range placeHolderMap {
		holder := "{{" + k + "}}"
		if strings.Contains(tc.Text, holder) {
			tc.Text = strings.ReplaceAll(tc.Text, holder, v)
		}
	}

	return tc.Text
}
