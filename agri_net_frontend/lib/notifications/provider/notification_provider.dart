import "package:http/http.dart";
import "../../libs.dart";
import "dart:convert";

class NotificationProvider {
  Client client = Client();

  Future<TransactionNotifications> getMyTransactionNotifications() async {
    try {
      final Map<String, String> headers = {
        "authorization": StaticDataStore.HEADERS["authorization"]!
      };
      var response = await client.get(
          Uri(
            scheme: "http",
            host: StaticDataStore.HOST,
            port: StaticDataStore.PORT,
            path: "/api/cxp/mytransactions",
          ),
          headers: headers);
      if (response.statusCode >= 100 && response.statusCode < 500) {
        final bosy = jsonDecode(response.body);
        return TransactionNotifications.fromJson(bosy);
      } else {
        return TransactionNotifications(
            statusCode: response.statusCode,
            msg: STATUS_CODES[response.statusCode] ?? "",
            transactionNotifications: []);
      }
    } catch (e) {
      print(e.toString());
      return TransactionNotifications(
          statusCode: 999,
          msg: STATUS_CODES[999] ?? "",
          transactionNotifications: []);
    }
  }
}
