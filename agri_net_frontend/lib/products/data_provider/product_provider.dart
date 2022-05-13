import 'dart:convert';
import 'dart:math';

import 'package:http/http.dart';

import '../../libs.dart';

class ProductProvider {
  // api provider
  static Client client = Client();
  ProductProvider();

  Future<Product> getProducts() async {
    var respo = await client.get(
      Uri(
        scheme: "http",
        host: StaticDataStore.HOST,
        port: StaticDataStore.PORT,
        path: "/api/product/get",
      ),
    );
    return Product.fromJson(json.decode(respo.body));
  }

  Future<ProductResponse> createProduct(int pid, String productName,
      String location, double amounte, double price) async {
    try {
      var res = await client.post(
          Uri(
            scheme: "http",
            host: StaticDataStore.HOST,
            port: StaticDataStore.PORT,
            path: "/api/product/create",
          ),
          body: {
            "pid": pid,
            "productName": productName,
            "location": location,
            "amounte": amounte,
            "price": price
          },
          headers: {
            "Content-Type": "application/json"
          });

      if (res.statusCode == 200 || res.statusCode == 201) {
        // copy the json resonse and
        var body = jsonDecode(res.body) as Map<String, dynamic>;
        // final map = body.map<Map<String, dynamic>>((elem) {
        //   return (elem as Map<String, dynamic>);
        // }).toList();
        // return map;
        return ProductResponse(
            statusCode: res.statusCode,
            // msg: "$body["msg"]",
            msg: "${body["msg"]}",
            product: Product.fromJson(body["msg"] as Map<String, dynamic>));
      } else {
        var body = jsonDecode(res.body) as Map<String, dynamic>;
        return ProductResponse(
            statusCode: res.statusCode, msg: "${body["msg"]}");
      }
    } catch (e, a) {
      return ProductResponse(
          statusCode: 999, msg: "Sorry something went wrong");
    }
  }

  // getALLProduct() {}
  // getProductById() {}
  // updateProduct(int pid) {}
  // deleteProduct(int pid) {}
}
