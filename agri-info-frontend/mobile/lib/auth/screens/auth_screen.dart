import "../../libs.dart";

class AuthScreen extends StatefulWidget {
  static const String RouteName = "/auth_screen";

  const AuthScreen({Key? key}) : super(key: key);

  @override
  State<AuthScreen> createState() => _AuthScreenState();
}

class _AuthScreenState extends State<AuthScreen> {
  bool right = false;
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text("Naty"),
        elevation: 0,
        centerTitle: true,
      ),
      drawer: HomeNavigationDrawer(),
      body: Container(
        child: Center(
          child: Row(
            children: [
              AnimatedContainer(
                duration: Duration(
                  milliseconds: 400,
                ),
                color: Colors.yellow,
                width: MediaQuery.of(context).size.width * (right ? 0.3 : 0.7),
                child: Center(
                  child: ElevatedButton(
                    onPressed: () {
                      setState(() {
                        right = false;
                      });
                    },
                    child: Text("Click Me"),
                  ),
                ),
              ),
              AnimatedContainer(
                duration: Duration(
                  milliseconds: 400,
                ),
                width: MediaQuery.of(context).size.width * (right ? 0.7 : 0.3),
                color: Colors.blue,
                child: Center(
                  child: ElevatedButton(
                    onPressed: () {
                      setState(() {
                        right = true;
                      });
                    },
                    child: Text("clieck meeey"),
                  ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
