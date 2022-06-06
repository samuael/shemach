class StaticDataStore {
  static const SCHEME = "http://";
  static const HOST ="192.168.43.208";
  static const PORT = 8080;
  static String TOKEN = ""; 
  static int ID=0;

  static get getURI{
    return SCHEME+HOST+":$PORT";
  }

  static Map<String, String>  get headers {
    return {
      "Authorization": "Bearer ${TOKEN}"
    };
  }
  
}