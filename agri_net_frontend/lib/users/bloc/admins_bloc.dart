import '../../libs.dart';

class AdminsBloc extends Bloc<UsersEvent, UsersState>
    implements Cubit<UsersState> {
  UsersRepo usersRepo;
  AdminsBloc({required this.usersRepo}) : super(GetAllUsersState());

  @override
  Stream<UsersState> mapEventToState(UsersEvent usersEvent) async* {
    if (usersEvent is GetAllUsersEvent) {
      final admins = await usersRepo.getAdmins();
      if (admins.length > 0) {
        yield AllUsersRetrievedState(usersList: admins);
        print("Admins Loading Was succesful \n\n\n\n");
      } else {
        yield AdminsLoadingFailed();
      }
    }
    if (usersEvent is CreateNewUserEvent) {}
  }

  Future<UsersState> registerAdmin(
      CreateNewUserEvent createNewUserEvent) async {
    final product = await usersRepo.postUser(
        createNewUserEvent.id,
        createNewUserEvent.firstname,
        createNewUserEvent.lastname,
        createNewUserEvent.email,
        createNewUserEvent.imgurl,
        createNewUserEvent.phone,
        createNewUserEvent.imgurl,
        createNewUserEvent.lang);
  }
}
