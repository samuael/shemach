import 'package:http/http.dart';
import 'package:mobile/libs.dart';
import "dart:convert";

class AuthDataProvider {
  Client client = Client();

  Future<ResponseSubscription> registerUser(
      String fullname, String phone, int role, String lang) async {
    try {
      var response = await client.post(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
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
Future<AuthenticationResponse>  loginSubscriber(String phone) async {
  try {
      var response = await client.post(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          scheme: StaticDataStore.SCHEME,
          path: "/api/subscription/login",
        ),
        body: jsonEncode(
          {"phone" : phone},
        ),
      );
      print(response.statusCode );
      if (response.statusCode == 201 || response.statusCode == 200){
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return AuthenticationResponse.fromJson(json);
      }else if (response.statusCode < 500 && response.statusCode >=200 ){
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return AuthenticationResponse.fromJson(json);
      }else {
        return AuthenticationResponse( msg: "Server Problem", errors:{}, statusCode: 500 );
      }
    } catch (e, a) {
      print(e.toString());
      return  AuthenticationResponse( msg: "Network connection problem",errors:{},statusCode:999);
    }
}
  Future<AuthenticationResponse>  register( RegistrationInput input) async {
      try {
      var response = await client.post(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          scheme: StaticDataStore.SCHEME,
          path: "/api/info/register",
        ),
        body: jsonEncode(
          input.toJson(),
        ),
      );
      print(response.statusCode );
      if (response.statusCode == 201 || response.statusCode == 200){
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return AuthenticationResponse.fromJson(json);
      }else if (response.statusCode < 500 && response.statusCode >=200 ){
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return AuthenticationResponse.fromJson(json);
      }else {
        return AuthenticationResponse( msg: "Server Problem", errors:{}, statusCode: 500 );
      }
    } catch (e, a) {
      print(e.toString());
      return  AuthenticationResponse( msg: "Network connection problem",errors:{},statusCode:999);
    }
  }


  Future<SubscriberAuthenticationRespnse>  confirmRegistration(SubscriberConfirmation input) async{
    try {
      var response = await client.post(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          scheme: StaticDataStore.SCHEME,
          path: "/api/subscription/registration/confirm",
        ),
        body: jsonEncode(
          input.toJson(),
        ),
      );
      print(response.statusCode );
      if (response.statusCode == 201  || response.statusCode == 200){
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return SubscriberAuthenticationRespnse.fromJson(json);
      }else if (response.statusCode < 500 && response.statusCode >=200 ){
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return SubscriberAuthenticationRespnse.fromJson(json);
      }else {
        return SubscriberAuthenticationRespnse( msg: "Server Problem", statusCode: 500, token : '', subscriber: null);
      }
    } catch (e, a) {
      print(e.toString());
      return  SubscriberAuthenticationRespnse( msg: "Network connection problem", token: '', subscriber: null,statusCode:999);
    }
  }

  Future<SubscriberAuthenticationRespnse>  confirmLogin(SubscriberConfirmation input) async{
    try {
      var response = await client.post(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          scheme: StaticDataStore.SCHEME,
          path: "/api/subscription/confirm",
        ),
        body: jsonEncode(
          input.toJson(),
        ),
      );
      print(response.statusCode );
      if (response.statusCode == 201  || response.statusCode == 200){
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return SubscriberAuthenticationRespnse.fromJson(json);
      }else if (response.statusCode < 500 && response.statusCode >=200 ){
        final json = jsonDecode(response.body) as Map<String, dynamic>;
        return SubscriberAuthenticationRespnse.fromJson(json);
      }else {
        return SubscriberAuthenticationRespnse( msg: "Server Problem", statusCode: 500, token : '', subscriber: null);
      }
    } catch (e, a) {
      print(e.toString());
      return  SubscriberAuthenticationRespnse( msg: "Network connection problem", token: '', subscriber: null,statusCode:999);
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
