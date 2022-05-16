import '../../libs.dart';

class UserBloc extends Bloc<UserEvent, UserState> implements Cubit<UserState> {
  UserBloc(UserState initialState) : super(initialState);

  Stream<UserState> mapEventToState(UserEvent userEvent) async* {
    if (userEvent is UserLoggedInSuccessEvent) {
      yield (UserLoggedInSuccessState());
    }
    if (userEvent is AgentLoggedInSuccessEvent) {
      yield (AgentLoggedInSuccessState());
    }
    if (userEvent is SuperAdminLoggedInSuccessState) {
      yield (SuperAdminLoggedInSuccessState());
    }
    if (userEvent is MerchantLoggedInSuccessState) {
      yield (MerchantLoggedInSuccessState());
    }
    ;
  }

  Future<UserState> whoLoggedIn(UserLoggedInSuccessEvent successEvent) async {
    final userState = successEvent.user;
    final userRole = successEvent.role;

    if (userState != null && userRole != null && userRole == "superadmin") {
      return (SuperAdminLoggedInSuccessState());
    }
    if (userState != null && userRole != null && userRole == "agent") {
      return (AgentLoggedInSuccessState());
    }
    if (userState != null && userRole != null && userRole == "merchant") {
      return (MerchantLoggedInSuccessState());
    }
    return UserWithNoRoleState();
  }
}
