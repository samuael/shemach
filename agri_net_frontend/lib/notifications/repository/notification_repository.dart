import "../../libs.dart";

class NotificationRepository {
  NotificationProvider provider;

  NotificationRepository(this.provider);

  Future<TransactionNotifications> getMyTransactionNotifications() async {
    return this.provider.getMyTransactionNotifications();
  }
}
