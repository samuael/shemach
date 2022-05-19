import 'package:agri_net_frontend/auth/model/modelUser.dart';

class UsersState {}

class GetAllUsersState extends UsersState {}

class LoadingUsersState extends UsersState {}

class AllUsersRetrievedState extends UsersState {
  List<User> usersList = [];
  AllUsersRetrievedState({required this.usersList});
}

class NoUserFoundState extends UsersState {}

class CreateNewUserState extends UsersState {
  User newUser;
  CreateNewUserState({required this.newUser});
}

class SthWentWrongState extends UsersState {}

// AdminsLoadingFailed
class AdminsLoadingFailed extends UsersState {}
