import '../../libs.dart';

class HomeScreen extends StatefulWidget {
  static const String RouteName = '/message_screen';
  const HomeScreen({Key? key}) : super(key: key);

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  bool ishome = true;
  bool isSelectedMessage = true;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      //backgroundColor: Color(0xbae8e8),
      appBar: AppBar(
        title: Text(
          "Agri Info",
          style: TextStyle(
            fontWeight: FontWeight.bold,
          ),
        ),
        actions: <Widget>[
          IconButton(
            icon: Icon(
              Icons.more_vert,
              color: Colors.white,
            ),
            onPressed: () {
              // do something
            },
          ),
        ],
        elevation: 0,
        centerTitle: true,
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
                            ishome = !ishome;
                          });
                        },
                        child: Row(children: [
                                Text(
                                  "Products ",
                                  style: TextStyle(
                                    fontWeight: FontWeight.bold,
                                    color:  ishome ? Theme.of(context).canvasColor : Colors.white,
                                    fontFamily: "Roboto",
                                  ),
                                ),
                                Icon(
                                  Icons.home,
                                  color: ishome ? Theme.of(context).canvasColor : Colors.white,
                                ),
                              ])
                            
                      ),
                    ),
                    SizedBox(width: 10),
                    Container(
                      padding: EdgeInsets.symmetric(vertical: 10, horizontal: 20),
                      child: InkWell(
                        onTap: () {
                          setState(() {
                            ishome = !ishome;
                          });
                        },
                        child: Row(
                          children: [
                                Text(
                                  "Messages ",
                                  style: TextStyle(
                                    fontWeight: FontWeight.bold,
                                    color:  ishome ?   Colors.white:Theme.of(context).canvasColor,
                                    fontFamily: "Roboto",
                                  ),
                                ),
                                Icon(
                                  Icons.message,
                                  color:ishome ?   Colors.white:Theme.of(context).canvasColor,
                                ),
                                ClipRRect(
                                  borderRadius: BorderRadius.circular(10),
                                  child: Container(
                                    padding : EdgeInsets.all(5),
                                    color: Colors.red,
                                    child: Text(
                                      "5",
                                      style: TextStyle(
                                        color: Colors.white,
                                        fontWeight: FontWeight.bold,
                                      ),
                                    ),
                                  ),
                                ),
                              ])
                            
                      ),
                    ),
                  ],
                )),
            Container(
              height: MediaQuery.of(context).size.height * 0.85,
              child: SingleChildScrollView(
                scrollDirection: Axis.vertical,
                child: Column(
                  children: [
                    // ProductItem(),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
