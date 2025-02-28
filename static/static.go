package static

import (
	_ "embed"
)

//go:embed html-summary.html
var HtmlTemplate string

//go:embed web_fingerprint_v4.json
var FingerPrintV4 []byte
