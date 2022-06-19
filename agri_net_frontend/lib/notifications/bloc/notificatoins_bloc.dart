import "../../libs.dart";

class NotificationsBloc extends Bloc<NotificationEvent, NotificationState> {
  NotificationRepository repo;
  NotificationsBloc(this.repo) : super(NotificationsInit()) {
    on<NotificationsLoadEvent>((event, emit) async {
      final response = await this.repo.getMyTransactionNotifications();
      hoolders -= 1;
      if (response.statusCode == 200) {
        final dstate = NotificationsLoadSuccess(response.transactionNotifications);
        dstate.generateCount();
        emit(dstate);
      } else {
        emit(NotificationsLoadFailure());
      }
    });
  }

  bool looping = true;
  void stopTransactionNotificationsLoop() {
    looping = false;
    hoolders -= 1;
  }

  static int hoolders = 0;
  Future<void> startLoadTransactionNotificationsLoop() async {
    if (hoolders >= 1) {
      return;
    }
    hoolders += 1;
    this.looping = true;
    while (looping) {
      await Future.delayed(Duration(minutes: 1), () {
        this.add(NotificationsLoadEvent());
      });
      await Future.delayed(Duration(seconds: 40), () {});
    }
  }
}
