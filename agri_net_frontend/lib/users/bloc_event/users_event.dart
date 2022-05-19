import '../../libs.dart';

class UsersEvent {}

class GetAllUsersEvent extends UsersEvent {}

class AllUsersRetrievedEvent extends UsersEvent {
  List<User> usersList = [];
  AllUsersRetrievedEvent({required this.usersList});
}

class NoUserFound extends UsersEvent {}

class CreateNewUserEvent extends UsersEvent {
  int id;
  String firstname;
  String lastname;
  String phone;
  String email;
  String lang;
  String imgurl;
  CreateNewUserEvent({required this.id,required this.firstname,required this.lastname,required this.phone,required this.email,required this.imgurl,required this.lang})
}

class SomethingWrongEvent extends UsersEvent {}
