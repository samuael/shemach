import '../../libs.dart';

class UserRepo {
  AuthProvider? provider;

  UserRepo({@required this.provider});

  Future<UsersLoginResponse> loginAdmin(String email, String password) async {
    return await this.provider!.loginAdmin(email, password);
  }

  Future<MessageOnly> forgotPassword(String email) async {
    return this.provider!.forgotPassword(email);
  }
}
