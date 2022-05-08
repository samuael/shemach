import 'libs.dart';

void main() {
  runApp(MultiBlocProvider(providers: [
    BlocProvider(
      create: (context) {
        return AuthBloc(
          repo: AuthRepo(
            provider: AuthProvider(),
          ),
        );
      },
    )
  ], child: MyHomePage()));
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({Key? key}) : super(key: key);

  @override
  State<MyHomePage> createState() => MyHomePageState();
}

class MyHomePageState extends State<MyHomePage> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
        title: 'Agri-Net',
        theme: ThemeData(
          primaryColor: Colors.green, //  MaterialColor(primary, swatch),
          canvasColor: Colors.white,
        ),
        initialRoute: AuthScreen.RouteName,
        onGenerateRoute: (setting) {
          final route = setting.name;
          if (route == AuthScreen.RouteName) {
            // final param = setting.arguments as Map<String, Product>;
            return MaterialPageRoute(builder: (context) {
              return AuthScreen();
            });
          } else if (route == HomeScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return HomeScreen();
            });
          }
        });
  }
}
