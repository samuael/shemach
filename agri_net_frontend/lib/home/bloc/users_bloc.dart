import '../../libs.dart';

class UsersBloc extends Bloc<UsersEvent, UsersState> {
  UsersRepo repo;
  UsersBloc(this.repo) : super(UsersInit()) {
    on<LoadUserByIDEvent>((event, emit) async {
      if (event.id <= 0) {
        return;
      }
      if (!(this.state is UsersLoadedState)) {
        emit(UsersLoading());
      } else {
        if (this.getUserByID(event.id) != null) {
          emit(this.state);
          return;
        }
      }
      final userResponse = await this.repo.provider.getUserByID(event.id);
      if (userResponse.statusCode == 200) {
        if (this.state is UsersLoadedState) {
          (this.state as UsersLoadedState).users[event.id] = userResponse.user!;
        } else {
          emit(UsersLoadedState({event.id: userResponse.user!}, merchants: {}));
        }
      } else {
        emit(UsersLoadFailedState());
      }
    });

    on<AddUserInstanceEvent>((event, emit) {
      if (this.state is UsersLoadedState) {
        (this.state as UsersLoadedState).users[event.user.id] = event.user;
      } else {
        emit(UsersLoadedState({event.user.id: event.user}, merchants: {}));
      }
    });

    on<AddUsersEvent>((event, emit) {
      for (final user in event.users) {
        if (this.state is UsersLoadedState) {
          (this.state as UsersLoadedState).users[user.id] = user;
        } else {
          emit(UsersLoadedState({user.id: user}, merchants: {}));
        }
      }
    });

    on<LoadMerchantByStoreIDEvent>((event, emit) async {
      if (event.storeID <= 0) {
        return;
      }
      if (!(this.state is UsersLoadedState)) {
        emit(UsersLoading());
      } else {
        if (this.getMerchantByStoreID(event.storeID) != null) {
          emit(this.state);
          return;
        }
      }
      final userResponse =
          await this.repo.provider.getMerchantByStoreID(event.storeID);
      print("The response is ${userResponse.statusCode} : ${userResponse.msg}");
      if (userResponse.statusCode == 200) {
        if (this.state is UsersLoadedState) {
          (this.state as UsersLoadedState).users[userResponse.user!.id] =
              userResponse.user!;
          (this.state as UsersLoadedState).merchants[event.storeID] =
              (userResponse.user! as Merchant);
        } else {
          emit(UsersLoadedState({userResponse.user!.id: userResponse.user!},
              merchants: {event.storeID: (userResponse.user as Merchant)}));
        }
      }
    });
  }

  User? getMerchantByStoreID(int storeID) {
    if (this.state is UsersLoadedState) {
      for (final sid in (this.state as UsersLoadedState).merchants.keys) {
        if (storeID == sid) {
          return (this.state as UsersLoadedState).merchants[sid];
        }
      }
    }
    this.add(LoadMerchantByStoreIDEvent(storeID));
    return null;
  }

  User? getUserByID(int userid) {
    if (this.state is UsersLoadedState) {
      for (int a in (this.state as UsersLoadedState).users.keys) {
        if (a == userid) {
          return (this.state as UsersLoadedState).users[userid]!;
        }
      }
    }
    return null;
  }
}
