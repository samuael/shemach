import '../../libs.dart';

class Product {
  int pid;
  String productName;
  String location;
  double amounte;
  double price;

  Product(
      {required this.pid,
      required this.productName,
      required this.location,
      required this.amounte,
      required this.price});

  factory Product.fromJson(Map<String, dynamic> json) {
    return Product(
        pid: int.parse("${json['id']}"),
        productName: json["productName"],
        location: json["location"],
        amounte: double.parse(json['amounte']),
        price: double.parse(json['price']));
  }

  Map<String, dynamic> toJson() {
    return {
      "pid": this.pid,
      "productName": this.productName,
      "location": this.location,
      "amounte": this.amounte,
      "price": this.price
    };
  }
}
