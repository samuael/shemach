import 'package:http/http.dart';
import 'package:mobile/libs.dart';
import "dart:convert";

class AuthenticationDataProvider {
  Client client = Client();

  Future<ResponseSubscription> registerUser(
      String fullname, String phone, int role, String lang) async {
    try {
      var response = await client.post(
        Uri(
          host: StaticDataStore.HOST,
          port: int.parse(StaticDataStore.PORT),
          scheme: StaticDataStore.SCHEME,
          path: "/api/info/register",
        ),
        body: jsonEncode(
          {
            "fullname": fullname,
            "phone": phone,
            "role": role,
            "lang": lang,
          },
        ),
      );
      print(response.statusCode);
      if (response.statusCode == 200) {
        final body = jsonDecode(response.body);
        final respons = ResponseSubscription.fromJson(body);
        return respons;
      } else if (response.statusCode == 500) {
        return ResponseSubscription(
            msg: "internal server error",
            statusCode: response.statusCode,
            errors: {});
      } else {
        return ResponseSubscription(
            msg: "", statusCode: response.statusCode, errors: {});
      }
    } catch (e, a) {
      return ResponseSubscription(
          msg: "connection problem", statusCode: 999, errors: {});
    }
  }

  // Future<ResponseSubscription> loginSubscription(
  //     String phone, String confirmation) async {
  //   try {
  //     var response = await client.post(
  //       Uri(
  //         host: StaticDataStore.HOST,
  //         port: int.parse(StaticDataStore.PORT),
  //         scheme: StaticDataStore.SCHEME,
  //         path: "/api/subscription/registration/confirm",
  //       ),
  //       body: jsonEncode(
  //         {
  //           "phone": phone,
  //           "confirmation": confirmation,
  //         },
  //       ),
  //     );

  //     print("${response.statusCode} ${(response.body).toString()}");
  //     if (response.statusCode == 200){

  //     }else if (response.statusCode == 500 ){

  //     }else {

  //     }
  //   } catch (e, a) {
  //     return 
  //   }
  // }
}
