final dictionary = {
  "register"  : {
    "amh"  : "temezgeb",
    "oro"  : "temezgeb",
    "tig"  : "temezgeb",
  }
};


String translate(String lang , String sentence) {
  final str = sentence;
	switch ( lang.toLowerCase() ) {
	case "en": case "eng":{
		return sentence;
  }
	case "amh":case "am":case "amharic":case "amhara":{
		sentence = ((dictionary[sentence.toLowerCase()])!["amh"])!;
    break;
  }
	case "oro": case "or":  case "oromifa": case"oromo":{
		sentence = (dictionary[sentence.toLowerCase()]!["oro"])!;
		return sentence;
  }
	case "tigr": case "tig": case "tigray": case "tigrigna":{
		sentence = (dictionary[sentence])!["tig"]!;
		return sentence;
  }
	}
	if (sentence == "") {
		return str;
	}
	return sentence;
}