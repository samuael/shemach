import "../../libs.dart";

class UsersLoginResponse {
  int statusCode;
  String msg;
  Admin? user;

  UsersLoginResponse({
    required this.statusCode,
    required this.msg,
    this.user,
  });
}
