import '../../libs.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import "dart:convert";

class MessagesBloc extends Bloc<MessageEvent, MessageState> {
  final MessagesRepository repository;

  MessagesBloc(this.repository) : super(MessagesInit()) {
    on<MessagesLoadEvent>((event, emit) async {
      final messresponse = await this.repository.fetchMessages();
      if (messresponse.messages.length > 0 && messresponse.statusCode == 200) {
        emit(MessagesLoaded(messresponse.messages));
      }
    });
    on<MessageAdd>((event, emit) {
      final dstate = this.state;
      if (dstate is MessagesLoaded) {
        for (final mes in dstate.messages){
          if (mes.id ==event.message.id){
            return;
          }
        }
        dstate.messages.add(event.message);
        emit(MessagesInit());
        emit(dstate);
      } else {
        emit(MessagesLoaded([event.message]));
      }
    });
    on<MessageDrop>((event, emit) {
      final dstate = this.state;
      if (dstate is MessagesLoaded) {
        dstate.messages.removeWhere((e) {
          return e.id == event.message.id;
        });
        emit(MessagesInit());
        emit(dstate);
      }
    });
    on<SetMessageSeenEvent>((event, emit) {
      final dstate = this.state;
      if (dstate is MessagesLoaded) {
        for (int a = 0; a < dstate.messages.length; a++) {
          if (dstate.messages[a].id == event.productID) {
            dstate.messages[a].seen = true;
          }
        }
        emit(MessagesInit());
        emit(dstate);
      }
    });
  }

  int countUnreadMessage() {
    int counter = 0;
    if (this.state is MessagesLoaded) {
      for (final val in (this.state as MessagesLoaded).messages) {
        if (!(val.seen)) {
          counter++;
        }
      }
    }
    return counter;
  }
}
