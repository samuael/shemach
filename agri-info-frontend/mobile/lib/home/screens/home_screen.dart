import "package:flutter_bloc/flutter_bloc.dart";
import '../../libs.dart';
import 'package:web_socket_channel/io.dart';
import 'dart:async';
import "dart:convert";

IOWebSocketChannel? channel;
Stream<dynamic>? broadcastStream;

class HomeScreen extends StatefulWidget {
  static const String RouteName = '/message_screen';
  const HomeScreen({Key? key}) : super(key: key);

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  bool runs = false;
  Stream<dynamic>? stream;
  bool closeWebsocket = false;
  List<String> titles = ["Products", "Messages", "Subscriptions", "Settings"];

  @override
  void initState() {
    super.initState();
    connectChannel(StaticDataStore.ID);
  }

  int lastid = 0;

  Future<void> connectChannel(int susbscriberID) async {
    // setState(() {
    channel = IOWebSocketChannel.connect(Uri.parse(''
        'ws://${StaticDataStore.HOST}:${StaticDataStore.PORT}/api/connection/subscriber/$susbscriberID'));
    broadcastStream = channel!.stream.asBroadcastStream();
    // Future.delayed(Duration(seconds: 10), () async {
    //   while (true) {
    //     await Future.delayed(Duration(seconds: 5));
    //     if (channel != null || channel!.closeCode != null) {
    //       channel = IOWebSocketChannel.connect(Uri.parse(''
    //           'ws://${StaticDataStore.HOST}:${StaticDataStore.PORT}/api/connection/subscriber/$susbscriberID'));
    //       broadcastStream = channel!.stream.asBroadcastStream();
    //     }
    //   }
    // });
  }

  // bool ishome = true;
  bool searching = false;
  double dopacity = 0.0;
  TextEditingController controller = TextEditingController();
  bool tried = false;

  @override
  void dispose() {
    if (channel != null) {
      channel!.sink.close();
    }
    super.dispose();
  }

  bool messagesLoaded = false;

  String title = translate(lang, "Agri Info");

  @override
  Widget build(BuildContext context) {
    setState(() {
      this.title = this.titles[context.watch<IndexBloc>().state - 1];
      Future.delayed(Duration(seconds: 5), () {
        this.title = translate(lang, "Agri Info");
      });
    });
    final messagesProvider = BlocProvider.of<MessagesBloc>(context);
    if (!((messagesProvider.state is MessagesLoaded) && messagesLoaded)) {
      messagesProvider.add(MessagesLoadEvent());
      messagesLoaded = true;
    }
    final authProvider = BlocProvider.of<AuthBloc>(context);
    // if (!tried) {
    //   if (broadcastStream == null) {
    //     connectChannel(
    //         (authProvider.state as AuthSubscriberAuthenticated).subscriber.id);
    //     Future.delayed(Duration(seconds: 3), () {
    //       if (broadcastStream != null) {
    //         tried = true;
    //       }
    //     });
    //   }
    // } else if (broadcastStream == null) {
    //   connectChannel(
    //       (authProvider.state as AuthSubscriberAuthenticated).subscriber.id);
    // }
    final productProvider = BlocProvider.of<ProductsBloc>(context);
    return Scaffold(
      //backgroundColor: Color(0xbae8e8),
      appBar: AppBar(
        title: searching
            ? AnimatedOpacity(
                duration: Duration(seconds: 2),
                opacity: dopacity,
                child: ClipRRect(
                  borderRadius: BorderRadius.circular(80),
                  child: Container(
                    decoration: BoxDecoration(
                      color: Theme.of(context).canvasColor,
                      border: Border.all(
                        color: Theme.of(context).canvasColor,
                      ),
                    ),
                    padding: EdgeInsets.symmetric(
                      horizontal: 20,
                    ),
                    child: TextField(
                      controller: controller,
                      decoration: InputDecoration(
                        suffixIcon: Icon(
                          Icons.search,
                          color: Theme.of(context).primaryColor,
                          size: 15,
                        ),
                      ),
                      onChanged: (text) {
                        if (text.length > 0) {
                          productProvider.add(SearchProductEvent(text));
                        } else {
                          productProvider.add(ProductsLoadEvent());
                        }
                      },
                    ),
                  ),
                ),
              )
            : Text(
                title,
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                ),
              ),
        centerTitle: true,
        actions: [
          Container(
            padding: EdgeInsets.only(right: 10),
            child: IconButton(
              icon: Icon(searching ? Icons.search_off : Icons.search),
              onPressed: () {
                setState(() {
                  this.searching = !searching;
                  if (searching) {
                    this.dopacity = 1.0;
                  } else {
                    this.dopacity = 0.0;
                  }
                });
              },
              color: searching ? Colors.red : Colors.white,
              // size: 30,
            ),
          ),
        ],
        elevation: 0,
      ),
      drawer: NavigationDrawer(),
      body: SingleChildScrollView(
        child: Column(
          children: [
            Container(
              color: Theme.of(context).primaryColor,
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                mainAxisSize: MainAxisSize.max,
                children: [
                  Container(
                    padding: EdgeInsets.symmetric(vertical: 10, horizontal: 20),
                    child: InkWell(
                        onTap: () {
                          setState(() {
                            context.read<IndexBloc>().add(1);
                          });
                        },
                        child: Row(children: [
                          Icon(
                            Icons.home,
                            color: context.watch<IndexBloc>().state == 1
                                ? Colors.white
                                : Colors.white54,
                          ),
                        ])),
                  ),
                  SizedBox(width: 10),
                  Container(
                    padding: EdgeInsets.symmetric(vertical: 10, horizontal: 20),
                    child: InkWell(
                        onTap: () {
                          setState(() {
                            context.read<IndexBloc>().add(2);
                          });
                        },
                        child: Row(children: [
                          // Text(
                          //   "Messages ",
                          //   style: TextStyle(
                          //     fontWeight: FontWeight.bold,
                          //     color: context.watch<IndexBloc>().state == 2
                          //         ? Colors.white54
                          //         : Colors.white,
                          //     fontFamily: "Roboto",
                          //   ),
                          // ),
                          Icon(
                            Icons.message,
                            color: context.watch<IndexBloc>().state == 2
                                ? Colors.white
                                : Colors.white54,
                          ),
                          context.watch<MessagesBloc>().countUnreadMessage() > 0
                              ? ClipRRect(
                                  borderRadius: BorderRadius.circular(10),
                                  child: Container(
                                    padding: EdgeInsets.all(5),
                                    color: Colors.red,
                                    child: Text(
                                      "${context.watch<MessagesBloc>().countUnreadMessage()}",
                                      style: TextStyle(
                                        color: Colors.white,
                                        fontWeight: FontWeight.bold,
                                      ),
                                    ),
                                  ),
                                )
                              : SizedBox(),
                        ])),
                  ),
                  Container(
                    padding: EdgeInsets.symmetric(vertical: 10, horizontal: 20),
                    child: InkWell(
                        onTap: () {
                          setState(() {
                            context.read<IndexBloc>().add(3);
                          });
                        },
                        child: Row(children: [
                          Icon(
                            Icons.subscriptions,
                            color: context.watch<IndexBloc>().state == 3
                                ? Colors.white
                                : Colors.white54,
                          ),
                        ])),
                  ),
                  Container(
                    padding: EdgeInsets.symmetric(vertical: 10, horizontal: 20),
                    child: InkWell(
                      onTap: () {
                        setState(() {
                          context.read<IndexBloc>().add(4);
                        });
                      },
                      child: Row(
                        children: [
                          Icon(
                            Icons.settings,
                            color: context.watch<IndexBloc>().state == 4
                                ? Colors.white
                                : Colors.white54,
                          ),
                        ],
                      ),
                    ),
                  ),
                ],
              ),
            ),
            StreamBuilder(
                stream: broadcastStream,
                builder: (context, snapshot) {
                  if (snapshot.hasData) {
                    try {
                      final response = MessageResponse.fromJson(
                          jsonDecode(snapshot.data as String)!);
                      if (response.type == 1) {
                        if ((response.body as Message).id > 0 &&
                            (response.body as Message).id != lastid) {
                          context
                              .read<MessagesBloc>()
                              .add(MessageAdd(response.body as Message));
                          lastid = (response.body as Message).id;
                        }
                      } else if (response.type == 2) {
                        if ((response.body as ProductUpdate).productID > 0 &&
                            (response.body as ProductUpdate).cost > 0) {
                          context.read<ProductsBloc>().add(ProductUpdateEvent(
                              response.body as ProductUpdate));
                        }
                      }
                    } catch (e, a) {}
                  }
                  return Container(
                    color: Theme.of(context).primaryColorLight,
                    height: MediaQuery.of(context).size.height * 0.80,
                    child: SingleChildScrollView(
                      scrollDirection: Axis.vertical,
                      child: context.watch<IndexBloc>().state == 1
                          ? Products()
                          : (context.watch<IndexBloc>().state == 2
                              ? ChatMessages()
                              : (context.watch<IndexBloc>().state == 3
                                  ? SubscriptionPage()
                                  : ChatMessages())),
                    ),
                  );
                }),
          ],
        ),
      ),
    );
  }
}
