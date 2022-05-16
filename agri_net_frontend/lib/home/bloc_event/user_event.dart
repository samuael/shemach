import '../../libs.dart';

class UserEvent {}

class UserLoggedInEvent extends UserEvent {
  User user;
  String role;
  UserLoggedInEvent({required this.user, required this.role});
}

class UserWithNoRoleEvent extends UserEvent {}

class AgentLoggedInEvent extends UserEvent {}

class MerchantLoggedInEvent extends UserEvent {}

class AdminLoggedInEvent extends UserEvent {}

class SuperAdminLoggedInEvent extends UserEvent {}
