import 'dart:convert';

import 'package:http/http.dart';

import '../../libs.dart';

class ProductProvider {
  // api provider
  static Client client = Client();
  ProductProvider();

  Future<List<Product>> getProducts() async {
    final header = {"Authorization": "Bearer ${StaticDataStore.USER_TOKEN}"};
    List<Product> products = [];

    var respo = await client.get(
        Uri(
          scheme: "http",
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          path: "/api/products",
        ),
        headers: {"Content-Type": "application/json"});
    print(respo.body);
    if (respo.statusCode == 200 || respo.statusCode == 201) {
      Map<String, dynamic> bodyMap = jsonDecode(respo.body);
      List<dynamic> tempProducts = bodyMap["products"];
      for (int i = 0; i < tempProducts.length; i++) {
        products.add((Product.fromJson(tempProducts[i])));
      }
      print("Success\n\n");
      return products;
    } else {
      print("Failure\n\n");
      throw Exception('Failed to load post');
    }
  }

  Future<ProductResponse> createProduct(int unitId, String productName,
      String productionArea, double currentPrice) async {
    final header = {"Authorization": "Bearer${StaticDataStore.USER_TOKEN}"};
    try {
      var res = await client.post(
          Uri(
            scheme: "http",
            host: StaticDataStore.HOST,
            port: StaticDataStore.PORT,
            path: "/api/superadmin/product/new",
          ),
          body: {
            "unit_id": unitId,
            "name": productName,
            "production_area": productionArea,
            "current_price": currentPrice
          },
          headers: header);

      if (res.statusCode == 200 || res.statusCode == 201) {
        // copy the json resonse and
        var body = jsonDecode(res.body);
        return ProductResponse(
            statusCode: res.statusCode,
            msg: "${body["msg"]}",
            product: Product.fromJson(body["product"]));
      } else {
        var body = jsonDecode(res.body) as Map<String, dynamic>;
        return ProductResponse(
            statusCode: res.statusCode, msg: "${body["msg"]}");
      }
    } catch (e) {
      return ProductResponse(
          statusCode: 999, msg: "Sorry something went wrong");
    }
  }
}
