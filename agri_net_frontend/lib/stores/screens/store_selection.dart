import '../../libs.dart';

class StoreSelectionScreen extends StatefulWidget {
  static const String RouteName = "/store_selection_screen";
  final List<Store> stores;
  final Function callBack;

  StoreSelectionScreen(this.stores, this.callBack, {Key? key})
      : super(key: key);

  @override
  State<StoreSelectionScreen> createState() => StoreSelectionScreenState();
}

class StoreSelectionScreenState extends State<StoreSelectionScreen> {
  String text = "";

  Store? store;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        backgroundColor: Colors.transparent,
        body: Container(
          margin: EdgeInsets.symmetric(horizontal: 40, vertical: 80),
          child: ClipRRect(
            borderRadius: BorderRadius.circular(30),
            child: Container(
              color: Colors.white,
              padding: EdgeInsets.symmetric(
                horizontal: 10,
                vertical: 20,
              ),
              // margin: EdgeInsets.symmetric(horizontal: 40, vertical: 80),
              child: SingleChildScrollView(
                child: Column(
                  children: [
                    Container(
                      child: SingleChildScrollView(
                        child: Column(
                          mainAxisSize: MainAxisSize.min,
                          children: (context.read<StoreBloc>().state
                                  is MyStoresState)
                              ? (context.read<StoreBloc>().state
                                      as MyStoresState)
                                  .myStores
                                  .map((p) {
                                  return GestureDetector(
                                    onTap: () async {
                                      widget.callBack(p);
                                      await Future.delayed(
                                          Duration(seconds: 1));
                                      Navigator.of(context).pop();
                                    },
                                    child: Container(
                                      padding: EdgeInsets.symmetric(
                                        horizontal: 10,
                                      ),
                                      margin: EdgeInsets.symmetric(
                                        vertical: 2,
                                      ),
                                      decoration: BoxDecoration(
                                        borderRadius: BorderRadius.circular(
                                          5,
                                        ),
                                        border: Border.all(
                                          color: Theme.of(context)
                                              .primaryColorLight,
                                        ),
                                      ),
                                      child: Row(
                                        mainAxisAlignment:
                                            MainAxisAlignment.spaceBetween,
                                        children: [
                                          Column(children: [
                                            Text(
                                              p.storeName,
                                              style: TextStyle(
                                                fontStyle: FontStyle.italic,
                                                fontWeight: FontWeight.bold,
                                              ),
                                            ),
                                            Text(
                                              UnixTime(p.createdAt
                                                      .millisecondsSinceEpoch)
                                                  .toString(),
                                              style: TextStyle(
                                                color: Color.fromARGB(
                                                    137, 48, 47, 47),
                                              ),
                                            ),
                                          ]),
                                          Text(p.address.toString()),
                                          Icon(
                                            Icons.add,
                                            color:
                                                Theme.of(context).primaryColor,
                                          ),
                                        ],
                                      ),
                                    ),
                                  );
                                }).toList()
                              : [
                                  Center(
                                    child: Text(
                                      translate(
                                        lang,
                                        "\tSorry!!\n No store instance is found type found ",
                                      ),
                                      textAlign: TextAlign.center,
                                      style: TextStyle(
                                        fontWeight: FontWeight.bold,
                                        fontStyle: FontStyle.italic,
                                      ),
                                    ),
                                  ),
                                ],
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ),
          ),
        ));
  }
}
