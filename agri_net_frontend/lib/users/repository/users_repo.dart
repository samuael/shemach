import '../../libs.dart';

class UsersRepo {
  UsersProvider usersProvider;
  UsersRepo({required this.usersProvider});

  Future<List<User>> getAdmins() {
    return usersProvider.getAdmins();
  }

  Future<User> postUser(
    int id,
    String firstname,
    String lastname,
    String phone,
    String email,
    String lang,
    String imgurl,
  ) {
    return usersProvider.createNewUser(user);
  }
}
