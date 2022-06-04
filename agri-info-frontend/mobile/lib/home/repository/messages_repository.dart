import "../../libs.dart";

class MessagesRepository {
  final MessagesProvider provider ;
  MessagesRepository(this.provider);

  Future<MessagesResponse>  fetchMessages() async {
    return this.provider.fetchMessages();
  }

  
}