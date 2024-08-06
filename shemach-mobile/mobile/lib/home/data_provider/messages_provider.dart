import "package:http/http.dart";
import "../../libs.dart";
import "dart:convert";

class MessagesProvider {
  Client client = Client();

  Future<MessagesResponse> fetchMessages() async {
    try {
      var response = await client.get(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          scheme: StaticDataStore.SCHEME,
          path: "/api/messages",
        ),
        headers: StaticDataStore.headers,
      );
      print(response.statusCode);
      print(jsonDecode(response.body));
      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        final result = MessagesResponse.fromJson(body);
        result.statusCode = response.statusCode;
        return result;
      } else if (response.statusCode < 500 && response.statusCode >= 200) {
        final json = jsonDecode(response.body);
        return MessagesResponse(
          msg: json["msg"] ?? '',
          statusCode: response.statusCode,
          messages: [],
        );
      }
      return MessagesResponse(
          msg: 'internal problem, please try again!',
          statusCode: 500,
          messages: []);
    } catch (e, a) {
      print(e.toString());
      return MessagesResponse(
        msg: "connection problem!",
        messages: [],
        statusCode: 999,
      );
    }
  }
}
