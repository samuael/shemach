package translation

import "strings"

var LANGUAGES = []string{"amh", "oro", "tig"}

// DICTIONARY FOR SAVING DIFFERENT LANGUAGES
var DICTIONARY = map[string]map[string]string{
	"login": {
		"amh": "ግba",
		"oro": "ግba",
		"tig": "ግba",
	},
	"this is your confirmation code from agri-info systems": {
		"amh": "ይህ የአግሪ-ኢንፎ የማረጋገጫ ቁጥርዎ ነው።",
		"oro": "ይህ የአግሪ-ኢንፎ የማረጋገጫ ቁጥርዎ ነው።",
		"tig": "ይህ የአግሪ-ኢንፎ የማረጋገጫ ቁጥርዎ ነው።",
	},
}

// Translate  function to change the word to the needed Language Representation
func Translate(lang string, sentence string) string {
	switch strings.ToLower(lang) {
	case "en", "eng":
		return sentence
	case "amh", "am", "amharic", "amhara":
		return strings.ToTitle((DICTIONARY["amh"])[strings.ToLower(sentence)])
	case "oro", "or", "oromifa", "oromo":
		return strings.ToTitle((DICTIONARY["oromifa"])[strings.ToLower(sentence)])
	}
	return sentence
}
