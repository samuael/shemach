import '../../libs.dart';

class NewStoreResponse {
  int statusCode;
  String msg;
  Store? newStore;
  NewStoreResponse(
      {required this.statusCode, required this.msg, this.newStore});
}
