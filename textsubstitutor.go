package textsubstitutor

import (
	"strings"
)

type Data struct {
	Key   string
	Value string
}

type TextSubstitutor struct {
	StartsWith string
	EndsWith   string
}

func NewTextSubstitutor(startsWith string, endsWith string) *TextSubstitutor {
	return &TextSubstitutor{
		StartsWith: startsWith,
		EndsWith:   endsWith,
	}
}

func (ur *TextSubstitutor) ReplaceData(text string, data []Data) string {
	for _, d := range data {
		placeholder := ur.StartsWith + d.Key + ur.EndsWith
		text = strings.Replace(text, placeholder, d.Value, -1)
	}

	return text
}
