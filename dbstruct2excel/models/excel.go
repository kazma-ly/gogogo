package models

type (
	// Font font style
	Font struct {
		Bold      bool    `json:"bold"`
		Italic    bool    `json:"italic"`
		Underline string  `json:"underline"`
		Family    string  `json:"family"`
		Size      float64 `json:"size"`
		Strike    bool    `json:"strike"`
		Color     string  `json:"color"`
	}

	// Fill fill style
	Fill struct {
		Type    string   `json:"type"`
		Pattern int      `json:"pattern"`
		Color   []string `json:"color"`
		Shading int      `json:"shading"`
	}

	// Border border style
	Border struct {
		Type  string `json:"type"`
		Color string `json:"color"`
		Style int    `json:"style"`
	}

	// Style style
	Style struct {
		Border []Border `json:"border"`
		Fill   Fill     `json:"fill"`
		Font   *Font    `json:"font"`
	}
)
