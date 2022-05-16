import '../../libs.dart';

const String ROLE_AGENT = "agent";
const String ROLE_SUPERADMIN = "superadmin";
const String ROLE_MERCHANT = "merchant";
const String ROLE_INFOADMIN = "infoadmin";
const String ROLE_ADMIN = "admin";

class StaticDataStore {
  // static const String HOST = "10.5.202.116";

  static const String HOST = "10.5.228.227";
  static const int PORT = 8080;

  static const String SCHEME = "http://";
  static String get URI {
    return "$SCHEME$HOST:$PORT";
  }

  static String TOKEN = "";
  static Map<String, String> HEADERS = {};
  static DeviceType DType = DeviceType.Unknown;

  static bool isEmail(String email) {
    return RegExp(
            r"^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+")
        .hasMatch(email);
  }

  static final Map<String, String> LanguageMap = {
    "eng": "English",
    "english": "English",
    "en": "English",
    "am": "Amharic",
    "amh": "Amharic",
    "amharic": "Amharic",
  };
}
