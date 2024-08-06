import "package:http/http.dart";
import "../../libs.dart";
import "dart:convert";

class ProductsProvider {
  Client client = Client();

  Future<ProductsResponse> loadProducts() async {
    try {
      var response = await client.get(Uri(
        host: StaticDataStore.HOST,
        port: StaticDataStore.PORT,
        scheme: StaticDataStore.SCHEME,
        path: "/api/products",
      ));
      print(response.statusCode);
      if (response.statusCode == 200) {
        final result = ProductsResponse.fromJson(jsonDecode(response.body));
        result.statusCode = response.statusCode;
        return result;
      } else if (response.statusCode < 500 && response.statusCode >= 200) {
        return ProductsResponse(
            msg: jsonDecode(response.body)["msg"],
            statusCode: response.statusCode,
            products: []);
      } else {
        return ProductsResponse(
            msg: "Server Problem, please try again!",
            statusCode: 500,
            products: []);
      }
    } catch (e, a) {
      return ProductsResponse(
          msg: "Connection issue!", statusCode: 999, products: []);
    }
  }

  Future<ProductsResponse> searchProducts(String text) async {
    try {
      var response = await client.get(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          scheme: StaticDataStore.SCHEME,
          path: "/api/product/search",
          queryParameters: {"text": text},
        ),
        headers: StaticDataStore.headers,
      );
      print(response.statusCode);
      if (response.statusCode == 200) {
        final products = ((jsonDecode(response.body) ?? [])
                .map<Map<String, dynamic>>((e) => e as Map<String, dynamic>)
                .toList())
            .map<ProductType>((e) {
          return ProductType.fromJson(e);
        }).toList();
        final result = ProductsResponse(
            statusCode: response.statusCode,
            products: products,
            msg: "products found");
        result.statusCode = response.statusCode;
        return result;
      } else if (response.statusCode < 500 && response.statusCode >= 200) {
        return ProductsResponse(
            msg: "not found", statusCode: response.statusCode, products: []);
      } else {
        return ProductsResponse(
            msg: "Server Problem, please try again!",
            statusCode: 500,
            products: []);
      }
    } catch (e, a) {
      return ProductsResponse(
          msg: "Connection issue!", statusCode: 999, products: []);
    }
  }

  Future<StatusAndMessage> subscribeForProduct(int productID) async {
    try {
      final headers = StaticDataStore.headers;
      var response = await client.get(
        Uri(
            host: StaticDataStore.HOST,
            port: StaticDataStore.PORT,
            scheme: StaticDataStore.SCHEME,
            path: "/api/product/subscribe",
            queryParameters: {"product_id": "$productID"}),
        headers: headers,
      );
      print(response.statusCode);
      print(jsonDecode(response.body));
      if (response.statusCode == 200) {
        return StatusAndMessage.fromJson(jsonDecode(response.body));
      } else if (response.statusCode < 500 && response.statusCode >= 200) {
        return StatusAndMessage(response.statusCode, "can't subscribe");
      } else {
        return StatusAndMessage(response.statusCode, "can't subscribe");
      }
    } catch (e, a) {
      return StatusAndMessage(999, "Connection issue!");
    }
  }

  Future<StatusAndMessage> unSubscribeForProduct(int productID) async {
    try {
      final headers = StaticDataStore.headers;
      var response = await client.get(
        Uri(
            host: StaticDataStore.HOST,
            port: StaticDataStore.PORT,
            scheme: StaticDataStore.SCHEME,
            path: "/api/product/unsubscribe",
            queryParameters: {"product_id": "$productID"}),
        headers: headers,
      );
      print(response.statusCode);
      if (response.statusCode == 200) {
        return StatusAndMessage.fromJson(jsonDecode(response.body));
      } else if (response.statusCode < 500 && response.statusCode >= 200) {
        return StatusAndMessage(response.statusCode, "can't subscribe");
      } else {
        return StatusAndMessage(response.statusCode, "can't subscribe");
      }
    } catch (e, a) {
      return StatusAndMessage(999, "Connection issue!");
    }
  }
}
