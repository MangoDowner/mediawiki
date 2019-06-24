package php

import "html"

/**
 * Encode an attribute value for HTML output.
 * @param string $text
 * @return string HTML-encoded text fragment
 */
func EncodeAttribute(text string) (encValue string) {
	encValue = html.EscapeString(text)

	// Whitespace is normalized during attribute decoding,
	// so if we've been passed non-spaces we must encode them
	// ahead of time or they won't be preserved.
	encValue = Strtr(encValue, map[string]string{
			"\n": "&#10;",
			"\r": "&#13;",
			"\t": "&#9;",
		},
	)

	return encValue
}