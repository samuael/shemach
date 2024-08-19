package translation

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

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
	"you will recieve a message": {
		"amh": "",
		"tig": "you will recieve a message",
		"oro": "you will recieve a message",
	},
	"shemach": {
		"amh": "አግሪ-ኢንፎ",
		"oro": "አግሪ-ኢንፎ",
		"tig": "አግሪ-ኢንፎ",
	},
	"user with this phone is already in confirmation process": {
		"amh": "ይህ የስልክ ቁጥር በማረጋገጫ ሂደት ላይ ነው።",
		"oro": "ይህ የስልክ ቁጥር በማረጋገጫ ሂደት ላይ ነው።",
		"tig": "ይህ የስልክ ቁጥር በማረጋገጫ ሂደት ላይ ነው።",
	},
	"internal server problem, plese try again": {
		"amh": "የውስጥ ችግር; እባክዎ እንደገና ሞክሩ!",
		"oro": "የውስጥ ችግር; እባክዎ እንደገና ሞክሩ!",
		"tig": "የውስጥ ችግር; እባክዎ እንደገና ሞክሩ!",
	},
	"account with this phone already exist": {
		"amh": "በዚህ ስልክ ቁጥር የተመዘገበ አካውንት አለ",
		"oro": "በዚህ ስልክ ቁጥር የተመዘገበ አካውንት አለ",
		"tig": "በዚህ ስልክ ቁጥር የተመዘገበ አካውንት አለ",
	},
	"you will recieve an sms a message containing the confirmation code\nplease confirm your account with in 30 minutes.": {
		"amh": "የማረጋገጫ መልእክት የያዘ የሴኤምኤስ መልእክት ይደርስዎታል። እባክዎ በ 30 ደቂቃዎች ውስጥ ቁጥሩን በማስገባት ባለቤትነትዎን ያረጋግጡ።",
		"oro": "የማረጋገጫ መልእክት የያዘ የሴኤምኤስ መልእክት ይደርስዎታል። እባክዎ በ 30 ደቂቃዎች ውስጥ ቁጥሩን በማስገባት ባለቤትነትዎን ያረጋግጡ።",
		"tig": "የማረጋገጫ መልእክት የያዘ የሴኤምኤስ መልእክት ይደርስዎታል። እባክዎ በ 30 ደቂቃዎች ውስጥ ቁጥሩን በማስገባት ባለቤትነትዎን ያረጋግጡ።",
	},
	"this is your login confirmation code from agri-info systems": {
		"amh": "ይህ የአግሪ-ኢንፎ የማረጋገጫ ቁጥርዎ ነው።",
		"tig": "ይህ የአግሪ-ኢንፎ የማረጋገጫ ቁጥርዎ ነው።",
		"oro": "ይህ የአግሪ-ኢንፎ የማረጋገጫ ቁጥርዎ ነው።",
	},
	"you are registered succesfuly. welcome!": {
		"amh": "bemigeba temezgbewal",
		"tig": "nay temezgebna dehan",
		"oro": "sinmezegebi succesfulida",
	},
	"please complete your registration!": {
		"amh": "ebakwo mejemeria mzgebawn yichersu",
		"oro": "ebakwo mejemeria mzgebawn yichersu",
		"tig": "ebakwo mejemeria mzgebawn yichersu",
	},
}

func TranslateIt(sentence string) string {
	str := sentence
	sentence = strings.ToTitle((DICTIONARY[strings.ToLower(sentence)])["amh"])
	if sentence == "" {
		var er error
		sentence, er = translateText("am", str)
		if er != nil {
			println(er.Error())
			sentence = ""
		}
	}
	if strings.EqualFold(sentence, "") {
		return str
	}
	return sentence
}

// Translate  function to change the word to the needed Language Representation
func Translate(lang string, sentence string) string {
	str := sentence
	switch strings.ToLower(lang) {
	case "en", "eng":
		return sentence
	case "amh", "am", "amharic", "amhara":
		sentence = strings.ToTitle((DICTIONARY[strings.ToLower(sentence)])["amh"])
		if sentence == "" {
			var er error
			sentence, er = translateText("am", str)
			if er != nil {
				sentence = ""
			}
		}
	case "oro", "or", "oromifa", "oromo":
		sentence = strings.ToTitle((DICTIONARY[strings.ToLower(sentence)])["oro"])
		return sentence
	case "tigr", "tig", "tigray", "tigrigna":
		sentence = strings.ToTitle((DICTIONARY[strings.ToLower(sentence)])["tig"])
		return sentence
	}
	if sentence == "" {
		return str
	}
	return sentence
}

// translateText
func translateText(targetLanguage, text string) (string, error) {
	// text := "The Go Gopher is cute"
	ctx := context.Background()

	lang, err := language.Parse(targetLanguage)
	if err != nil {
		return "", fmt.Errorf("language.Parse: %v", err)
	}

	client, err := translate.NewClient(ctx)
	if err != nil {
		return "", err
	}
	defer client.Close()

	resp, err := client.Translate(ctx, []string{text}, lang, nil)
	if err != nil {
		return "", fmt.Errorf("Translate: %v", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("Translate returned empty response to text: %s", text)
	}
	return resp[0].Text, nil
}
