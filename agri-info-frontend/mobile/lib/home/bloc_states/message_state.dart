import "../../libs.dart";

class MessageState{}

class MessagesLoaded extends MessageState{
  List<Message> messages ;
  MessagesLoaded(this.messages);
}

class MessagesInit extends MessageState{}

class MessageLoadingFailed extends MessageState{
  MessageResponse response;
  MessageLoadingFailed(this.response);
}