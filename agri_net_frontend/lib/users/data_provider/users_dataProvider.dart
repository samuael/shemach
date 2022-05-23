import 'dart:convert';

import 'package:agri_net_frontend/auth/model/modelUser.dart';
import 'package:http/http.dart';

import '../../libs.dart';

class UsersProvider {
  // api provider
  static Client client = Client();
  UsersProvider();

  Future<List<User>> getAdmins() async {
    final headers = {"Authorization": "Bearer ${StaticDataStore.USER_TOKEN}"};
    List<User> usersList = [];
    var respo = await client.get(
        Uri(
          scheme: "http",
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          path: "/api/admins",
        ),
        headers: headers);
    if (respo.statusCode == 200) {
      List<dynamic> users = [];
      users = json.decode(respo.body);
      for (int i = 0; i < users.length; i++) {
        Map<String, dynamic> tempMap = users[i];
        usersList.add((User.fromJson(tempMap)));
      }
      return usersList;
    } else {
      // If that call was not successful, throw an error.
      throw Exception('Failed to load post');
    }
  }

  Future<User> createNewUser(
    int id,
    String firstname,
    String lastname,
    String phone,
    String email,
    String lang,
    String imgurl,
  ) async {
    final headers = {"Authorization": "Bearer ${StaticDataStore.USER_TOKEN}"};
    var respo = await client.post(
        Uri(
          scheme: "http",
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          path: "/api/admins",
        ),
        body: jsonEncode(
          {
            "id": id,
            "firstname": firstname,
            "lastname": lastname,
            "email": email,
            "phone": phone,
            "imgurl": imgurl,
            "lang": lang
          },
        ),
        headers: headers);
    if (respo.statusCode == 200) {
      final body = jsonDecode(respo.body);
      return User.fromJson(body["user"] as Map<String, dynamic>);
    }
  }
}
