class StaticDataStore {
  static const SCHEME = "http";
  static const HOST ="192.168.43.208";
  static const PORT = 8080;
  static const TOKEN = ""; 

  static get getURI{
    return SCHEME+HOST+":$PORT";
  }
  
}