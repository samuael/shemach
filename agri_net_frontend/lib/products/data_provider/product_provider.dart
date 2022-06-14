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
      final body = jsonDecode(response.body)??{} as Map<String, dynamic>;
      print(body);
      return ProductPostResponse.fromJson(body);
    } catch (e, a) {
      print(e.toString());
      return ProductPostResponse(
          statusCode: 999, msg: "Connection issue!!!", crop: null);
    }
  }
}
