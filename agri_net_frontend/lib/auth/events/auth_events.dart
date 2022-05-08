import "../../libs.dart";

class AuthEvent {}

class AuthLoginEvent extends AuthEvent {
  String email;
  String password;
  AuthLoginEvent(this.email, this.password);
}

class AuthForgotPasswordEvent extends AuthEvent {
  String email;
  AuthForgotPasswordEvent(this.email);
}

class AdminLoginInProgressEvent extends AuthEvent {}

class AuthStateInitEvent extends AuthEvent {}

class AuthAdminLoggedInEvent extends AuthEvent {
  Admin state;
  AuthAdminLoggedInEvent(this.state);
}

class AuthAdminLoginNotSuccesfulEvent extends AuthEvent {
  String Msg;
  AuthAdminLoginNotSuccesfulEvent(this.Msg);
}
