import '../../libs.dart';

class UserEvent {}

class UserLoggedInSuccessEvent extends UserEvent {
  User user;
  String role;
  UserLoggedInSuccessEvent({required this.user, required this.role});
}

class AgentLoggedInSuccessEvent extends UserEvent {}

class MerchantLoggedInSuccessEvent extends UserEvent {}

class SuperAdminLoggedInSucceeeEvent extends UserEvent {}

class UserWithNoRoleEvent extends UserEvent {}
