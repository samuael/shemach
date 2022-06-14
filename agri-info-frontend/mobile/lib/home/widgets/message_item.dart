import "../../libs.dart";
import 'package:visibility_detector/visibility_detector.dart';
import "package:flutter_bloc/flutter_bloc.dart";

class MessageItem extends StatefulWidget {
  Message message;
  MessageItem(this.message, {Key? key}) : super(key: key);

  @override
  State<MessageItem> createState() => _MessageItemState();
}

class _MessageItemState extends State<MessageItem> {
  @override
  Widget build(BuildContext context) {
    return VisibilityDetector(
      key: UniqueKey(),
      onVisibilityChanged: (VisibilityInfo info) {
        if (info.visibleFraction == 1.0) {
          context
              .read<MessagesBloc>()
              .add(SetMessageSeenEvent(widget.message.id));
        }
      },
      child: Padding(
        padding: EdgeInsets.symmetric(
          vertical: 10,
          horizontal: 10,
        ),
        child: ClipRRect(
          borderRadius: BorderRadius.only(
              topRight: Radius.circular(20),
              topLeft: Radius.circular(20),
              bottomRight: Radius.circular(20)),
          child: Container(
            color: Theme.of(context).canvasColor,
            width: MediaQuery.of(context).size.width,
            padding: EdgeInsets.symmetric(horizontal: 20, vertical: 20),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Column(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    // Column(
                    //   children: [
                    //     Container(
                    //       margin: EdgeInsets.symmetric(vertical: 8),
                    //       child: Container(
                    //         child: Text(
                    //           "Abebe Kebede",
                    //           style: TextStyle(
                    //               fontFamily: "Roboto",
                    //               fontWeight: FontWeight.bold),
                    //         ),
                    //       ),
                    //     ),
                    //   ],
                    // ),
                    Text(
                      widget.message.data,
                      style: TextStyle(
                        fontFamily: "ROboto",
                        fontWeight: FontWeight.bold,
                      ),
                    ),
                    Container(
                      decoration: BoxDecoration(
                        border: Border.all(
                          color: Colors.black54,
                        ),
                        borderRadius: BorderRadius.circular(10),
                      ),
                      margin: EdgeInsets.only(top: 5),
                      padding: EdgeInsets.symmetric(horizontal: 4),
                      child: Text(
                        "ከ " +
                            UnixTime(widget.message.createdAt).toString() +
                            " " +
                            "በፊት",
                        style: TextStyle(fontSize: 13),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
