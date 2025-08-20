import "../../libs.dart";

class ProductsResponse {
  int statusCode;
  List<ProductType>  products;
  String msg;

  ProductsResponse({required this.statusCode, required this.products, required this.msg});

  factory ProductsResponse.fromJson(Map<String , dynamic> json){
    return ProductsResponse(
      msg : json["msg"]?? '' ,
      statusCode : json["status_code"]??999,
      products : ProductType.fromListOfJSON(json["products"]), 
    );
  }
}