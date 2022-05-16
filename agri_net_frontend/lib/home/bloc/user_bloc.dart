import '../../libs.dart';

class UserBloc extends Bloc<UserEvent, UserState> implements Cubit<UserState> {
  UserBloc(UserState initialState) : super(initialState);

  Stream<UserState> mapEventToState(UserEvent userEvent) async* {
    if (userEvent is UserLoggedInEvent) {
      yield (UserLoggedInState());
    }
    if (userEvent is AgentLoggedInEvent) {
      yield (AgentLoggedInState());
    }
    if (userEvent is SuperAdminLoggedInState) {
      yield (SuperAdminLoggedInState());
    }
    if (userEvent is MerchantLoggedInState) {
      yield (MerchantLoggedInState());
    }
    ;
  }

  Future<UserState> whoLoggedIn(UserLoggedInEvent successEvent) async {
    final userState = successEvent.user;
    final userRole = successEvent.role;

    if (userState != null && userRole != null && userRole == "superadmin") {
      return (SuperAdminLoggedInState());
    }
    if (userState != null && userRole != null && userRole == "agent") {
      return (AgentLoggedInState());
    }
    if (userState != null && userRole != null && userRole == "merchant") {
      return (MerchantLoggedInState());
    }
    return UserWithNoRoleState();
  }
}
