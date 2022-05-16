import '../../libs.dart';

class HomeScreen extends StatefulWidget {
  static const String RouteName = "homescreen";
  HomeScreen({Key? key}) : super(key: key);

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
            child: Row(
              children: [
                CollapsingSideBarDrawer(),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
