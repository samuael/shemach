
import "../../libs.dart";

class NotificationsBloc extends Bloc< NotificationEvent , NotificationState > {
  NotificationsBloc() : super(NotificationsInit()) {
    on<NotificationsLoadEvent>((event , emit ){
      // 
    });
  }
}