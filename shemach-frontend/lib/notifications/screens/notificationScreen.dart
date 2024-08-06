import "../../libs.dart";

class NotificationsScreen extends StatefulWidget {
  static const RouteName = "/notification/routes";
  const NotificationsScreen({Key? key}) : super(key: key);

  @override
  State<NotificationsScreen> createState() => _NotificationsScreenState();
}

class _NotificationsScreenState extends State<NotificationsScreen> {
  int selectedIndex = 0;
  @override
  Widget build(BuildContext context) {
    return Container(
        // height: 100,
        // width: 100,
        // color: Theme.of(context).primaryColor,
        child: Column(children: [
      Container(
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(
            20,
          ),
          color: Theme.of(context).primaryColor,
        ),
        padding: EdgeInsets.symmetric(
          horizontal: 5,
          vertical: 10,
        ),
        margin: EdgeInsets.only(
          left: 40,
          top: 0,
          bottom: 0,
        ),
        child: SingleChildScrollView(
          scrollDirection: Axis.horizontal,
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
            children: [
              GestureDetector(
                onTap: () {
                  setState(() {
                    selectedIndex = 0;
                  });
                },
                child: Container(
                  padding: EdgeInsets.symmetric(
                    horizontal: 5,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(
                      20,
                    ),
                    color: selectedIndex == 0
                        ? Colors.white
                        : Theme.of(context).primaryColor,
                  ),
                  child: Row(
                    children: [
                      Text(
                        translate(lang, " My Transactions "),
                        style: TextStyle(
                            fontWeight: FontWeight.bold,
                            fontStyle: FontStyle.italic,
                            color: selectedIndex == 0
                                ? Theme.of(context).primaryColor
                                : Colors.white),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      (context.watch<NotificationsBloc>().state
                              is NotificationsLoadSuccess)
                          ? ClipRRect(
                            borderRadius: BorderRadius.circular(10), 

                              child: Container(
                                padding: EdgeInsets.symmetric(
                                    vertical: 5, horizontal: 5),
                                color: Colors.white,
                                child: Text(
                                    "${(context.watch<NotificationsBloc>().state as NotificationsLoadSuccess).transactions.length}"),
                              ),
                            )
                          : SizedBox(),
                    ],
                  ),
                ),
              ),
              Container(
                child: Text(
                  " | ",
                  style: TextStyle(
                    color: Colors.white,
                  ),
                ),
              ),
              GestureDetector(
                onTap: () {
                  setState(() {
                    selectedIndex = 1;
                  });
                },
                child: Container(
                  padding: EdgeInsets.symmetric(
                    horizontal: 5,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(
                      20,
                    ),
                    color: selectedIndex == 1
                        ? Colors.white
                        : Theme.of(context).primaryColor,
                  ),
                  child: Row(
                    children: [
                      Text(
                        translate(lang, " Transactions "),
                        style: TextStyle(
                            fontWeight: FontWeight.bold,
                            fontStyle: FontStyle.italic,
                            color: selectedIndex == 1
                                ? Theme.of(context).primaryColor
                                : Colors.white),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      (context.watch<NotificationsBloc>().state
                              is NotificationsLoadSuccess)
                          ? ClipRRect(
                            borderRadius: BorderRadius.circular(10), 

                              child: Container(
                                padding: EdgeInsets.symmetric(
                                    vertical: 5, horizontal: 5),
                                color: Colors.white,
                                child: Text(
                                  "${(context.watch<NotificationsBloc>().state as NotificationsLoadSuccess).transaction_notifications}",
                                  style: TextStyle(
                                      fontWeight: FontWeight.bold,
                                      color: Theme.of(context).primaryColor),
                                ),
                              ),
                            )
                          : SizedBox(),
                    ],
                  ),
                ),
              ),
              Container(
                child: Text(
                  " | ",
                  style: TextStyle(
                    color: Colors.white,
                  ),
                ),
              ),
              GestureDetector(
                onTap: () {
                  setState(() {
                    selectedIndex = 2;
                  });
                },
                child: Container(
                  padding: EdgeInsets.symmetric(
                    horizontal: 5,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(
                      20,
                    ),
                    color: selectedIndex == 2
                        ? Colors.white
                        : Theme.of(context).primaryColor,
                  ),
                  child: Row(
                    children: [
                      Text(
                        translate(lang, " Kebd "),
                        style: TextStyle(
                            fontWeight: FontWeight.bold,
                            fontStyle: FontStyle.italic,
                            color: selectedIndex == 2
                                ? Theme.of(context).primaryColor
                                : Colors.white),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      (context.watch<NotificationsBloc>().state
                              is NotificationsLoadSuccess)
                          ? ClipRRect(
                            borderRadius: BorderRadius.circular(10), 

                              child: Container(
                                padding: EdgeInsets.symmetric(
                                    vertical: 5, horizontal: 5),
                                color: Colors.white,
                                child: Text(
                                  "${(context.watch<NotificationsBloc>().state as NotificationsLoadSuccess).kebd_notifications}",
                                  style: TextStyle(
                                      fontWeight: FontWeight.bold,
                                      color: Theme.of(context).primaryColor),
                                ),
                              ),
                            )
                          : SizedBox(),
                    ],
                  ),
                ),
              ),
              Container(
                child: Text(
                  " | ",
                  style: TextStyle(
                    color: Colors.white,
                  ),
                ),
              ),
              GestureDetector(
                onTap: () {
                  setState(() {
                    selectedIndex = 3;
                  });
                },
                child: Container(
                  padding: EdgeInsets.symmetric(
                    horizontal: 5,
                  ),
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(
                      20,
                    ),
                    color: selectedIndex == 3
                        ? Colors.white
                        : Theme.of(context).primaryColor,
                  ),
                  child: Row(
                    children: [
                      Text(
                        translate(lang, " Guarantee "),
                        style: TextStyle(
                            fontWeight: FontWeight.bold,
                            fontStyle: FontStyle.italic,
                            color: selectedIndex == 3
                                ? Theme.of(context).primaryColor
                                : Colors.white),
                      ),
                      SizedBox(
                        width: 5,
                      ),
                      (context.watch<NotificationsBloc>().state
                              is NotificationsLoadSuccess)
                          ? ClipRRect(
                            borderRadius: BorderRadius.circular(10), 
                              child: Container(
                                padding: EdgeInsets.symmetric(
                                    vertical: 5, horizontal: 5),
                                color: Colors.white,
                                child: Text(
                                  "${(context.watch<NotificationsBloc>().state as NotificationsLoadSuccess).guarantee_notifications}",
                                  style: TextStyle(
                                      fontWeight: FontWeight.bold,
                                      color: Theme.of(context).primaryColor),
                                ),
                              ),
                            )
                          : SizedBox(),
                    ],
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
      // ---------------------------------------

      selectedIndex == 0
          ? TransactionsList()
          : (selectedIndex == 1
              ? TransactionNotificationView()
              : (selectedIndex == 2
                  ? KebdNotificationView()
                  : GuaranteeNotificationView()))
    ]));
  }
}
