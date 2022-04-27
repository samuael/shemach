class StaticDataStore {
  static const SCHEME = "http://";
  static const HOST ="127.0.0.1";
  static const PORT = "8080";

  static get getURI{
    return SCHEME+HOST+":"+PORT;
  }
  
}