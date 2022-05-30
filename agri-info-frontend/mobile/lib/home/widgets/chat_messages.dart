import "../../libs.dart";

class ChatMessages extends StatefulWidget {
  const ChatMessages({ Key? key }) : super(key: key);

  @override
  State<ChatMessages> createState() => _ChatMessagesState();
}

class _ChatMessagesState extends State<ChatMessages> {
  @override
  Widget build(BuildContext context) {
    print("Chat Messages");
    return Container(
      color : Colors.blue ,
      height: MediaQuery.of(context).size.height*0.7, 
      width : MediaQuery.of(context).size.width, 
      child : Column(
        children : [

        ]
      )
    );
  }
}