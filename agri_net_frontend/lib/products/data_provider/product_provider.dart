import "package:http/http.dart";
import "../../libs.dart";
import "dart:convert";

class ProductProvider {
  Client client = Client();

  Future<ProductPostResponse> createProductPost(ProductPostInput input) async {
    try {
      Map<String, String> headers = {
        "authorization": StaticDataStore.HEADERS["authorization"]!
      };

      var response = await client.post(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          scheme: StaticDataStore.SCHEME,
          path: "/api/cxp/post/new",
        ),
        headers: headers,
        body: jsonEncode(input.toJson()),
      );
      print(response.body);
      final body = jsonDecode(response.body) ?? {} as Map<String, dynamic>;
      print(body);
      return ProductPostResponse.fromJson(body);
    } catch (e, a) {
      print(e.toString());
      return ProductPostResponse(
          statusCode: 999, msg: "Connection issue!!!", crop: null);
    }
  }

  Future<ProductsResponse> loadMyProductPosts() async {
    try {
      final Map<String, String> headers = {
        "authorization": StaticDataStore.HEADERS["authorization"]!
      };
      var response = await client.get(
          Uri(
            scheme: "http",
            host: StaticDataStore.HOST,
            port: StaticDataStore.PORT,
            path: "/api/cxp/posts",
          ),
          headers: headers);
      print(response.statusCode);
      print(response.body);
      if (response.statusCode >= 100 && response.statusCode < 500) {
        final bosy = jsonDecode(response.body);
        print(bosy);
        return ProductsResponse.fromJson(bosy);
      } else {
        return ProductsResponse(
            statusCode: response.statusCode,
            msg: STATUS_CODES[response.statusCode] ?? "",
            posts: []);
      }
    } catch (e, a) {
      print(e.toString());
      return ProductsResponse(
          statusCode: 999, msg: "connection issue!", posts: []);
    }
  }

  Future<ProductsResponse> loadProducts(int offset, int limit) async {
    try {
      final Map<String, String> headers = {
        "authorization": StaticDataStore.HEADERS["authorization"]!
      };
      var response = await client.get(
          Uri(
            scheme: "http",
            host: StaticDataStore.HOST,
            port: StaticDataStore.PORT,
            path: "/api/posts",
            queryParameters: {
              "offset": "$offset",
              "limit": "$limit",
            },
          ),
          headers: headers);
      if (response.statusCode >= 100 && response.statusCode < 500) {
        final bosy = jsonDecode(response.body);
        print(bosy);
        return ProductsResponse.fromJson(bosy);
      } else {
        return ProductsResponse(
            statusCode: response.statusCode,
            msg: STATUS_CODES[response.statusCode] ?? "",
            posts: []);
      }
    } catch (e, a) {
      print(e.toString());
      return ProductsResponse(
          statusCode: 999, msg: "connection issue!", posts: []);
    }
  }
}
