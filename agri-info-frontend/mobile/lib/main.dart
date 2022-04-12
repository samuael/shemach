import "libs.dart" ;


void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter Demo',
      theme: ThemeData(
        primarySwatch: Colors.brown,
      ),
      initialRoute: AuthScreen.RouteName,
      routes: {
        AuthScreen.RouteName : (context){
          return AuthScreen();
        }
      },
    );
  }
}

