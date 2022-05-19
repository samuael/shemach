import '../../libs.dart';

class ProductResponse {
  // @required
  int statusCode;
  String msg;

  // not required
  Product? product;

  // Constructor
  ProductResponse({required this.statusCode, required this.msg, this.product});
}

class Message {
  int statusCode;
  String msg;
  Message(this.statusCode, this.msg);
}
