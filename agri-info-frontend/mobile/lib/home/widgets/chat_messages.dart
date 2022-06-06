import "../../libs.dart";
import "package:flutter_bloc/flutter_bloc.dart";

class ChatMessages extends StatefulWidget {
  const ChatMessages({Key? key}) : super(key: key);

  @override
  State<ChatMessages> createState() => _ChatMessagesState();
}

class _ChatMessagesState extends State<ChatMessages> {
  @override
  Widget build(BuildContext context) {
    print("Chat Messages");
    return Container(
      color: Theme.of(context).primaryColorLight,
      height: MediaQuery.of(context).size.height * 0.9,
      width: MediaQuery.of(context).size.width,
      child: BlocBuilder<MessagesBloc, MessageState>(
        builder: (context, state) {
          if (state is MessagesLoaded) {
            return SingleChildScrollView(
              child: Column(children: [
                ...state.messages.map((e) => MessageItem(e)).toList()
              ]),
            );
          }
          return Center(
              child: Column(
            mainAxisSize: MainAxisSize.max,
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Text(
                "No Message Is Loaded",
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                ),
              ),
              IconButton(
                onPressed: () {
                  context.read<MessagesBloc>().add(MessagesLoadEvent());
                },
                icon: Icon(Icons.refresh_outlined),
              ),
            ],
          )); 
        },
      ),
    );
  }
}
