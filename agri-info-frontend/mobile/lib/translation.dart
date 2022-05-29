final dictionary = {
  "register"  : {
    "amh"  : "temezgeb",
    "oro"  : "temezgeb",
    "tig"  : "temezgeb",
  },
  "confirm" : {
    "amh" : "aregagt",
    "oro" : "aregagt",
    "tig" : "aregagt",
  }
};

String lang="amh";

String translate(String lang , String sentence) {
  final str = sentence.trim();
	switch ( lang.toLowerCase() ) {
	case "en": case "eng":{
		return sentence;
  }
	case "amh":case "am":case "amharic":case "amhara":{
    final val = dictionary[sentence.toLowerCase()];
		sentence = (val !=null) ? (val["amh"]??'') : sentence;
    break;
  }
	case "oro": case "or":  case "oromifa": case"oromo":{
		final val = dictionary[sentence.toLowerCase()];
		sentence = (val !=null) ? (val["oro"]??'') : sentence;
    break;
  }
	case "tigr": case "tig": case "tigray": case "tigrigna":{
		final val = dictionary[sentence.toLowerCase()];
		sentence = (val !=null) ? (val["tig"]??'') : sentence;
    break;
  }
	}
	if (sentence == "") {
		return str;
	}
	return sentence;
}