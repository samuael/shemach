import '../../libs.dart';

class HomeScreen extends StatefulWidget {
  static const String RouteName = "homescreen";
  const HomeScreen({Key? key}) : super(key: key);

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: Scaffold(
        appBar: UserScreenAppBarDrawer(),
        body: Center(
          child: Container(
            // height: double.infinity,
            child: Row(
              children: [
                CollapsingSideBarDrawer(),
                // HomePage()
                UserScreenBody(),
              ],
            ),
          ),
        ),
        // bottomNavigationBar: BottomAppBar(
        //     child: Row(
        //   children: [
        //     Footer(),
        //   ],
        // )),
        // persistentFooterButtons: [Footer()],
      ),
    );
  }
}
