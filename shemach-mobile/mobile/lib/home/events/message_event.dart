import "../../libs.dart";

class MessageEvent {}

class MessageAdd extends MessageEvent {
  Message message;
  MessageAdd(this.message);
}

class MessageDrop extends MessageEvent {
  Message message;
  MessageDrop(this.message);
}

class MessagesLoadEvent extends MessageEvent {}

class SetMessageSeenEvent extends MessageEvent {
  int productID;
  SetMessageSeenEvent(this.productID);
}
