import '../../libs.dart';

class AdminsState {}

// Super-Admin
class AdminsStateInIt extends AdminsState {}

class LoadingAdminsState extends AdminsState {}

class AdminsLoadedState extends AdminsState {
  List<Admin> adminsList = [];
  AdminsLoadedState({required this.adminsList});
}

class AdminsLoadingFailed extends AdminsState {
  int statusCode;
  String msg;
  AdminsLoadingFailed({required this.statusCode, required this.msg});
}

class UserDeleteSuccess extends AdminsState {
  int statusCode;
  String msg;
  UserDeleteSuccess({required this.statusCode, required this.msg});
}

class UserDeleteFailed extends AdminsState {
  int statusCode;
  String msg;
  UserDeleteFailed({required this.statusCode, required this.msg});
}
