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
      // routes: {
      //   AuthScreen.RouteName : (context){
      //     return AuthScreen();
      //   }
      // },
      onGenerateRoute: (setting){
        switch(setting.name){
          case RegistrationScreen.RouteName:{
            return MaterialPageRoute(builder: (context){
              return RegistrationScreen();
            });
          }
          case AuthScreen.RouteName : {
            return MaterialPageRoute(builder: (context){
              return AuthScreen();
            });
          }
          case ConfirmationScreen.RouteName : {
            return MaterialPageRoute( builder : (context){
              final String phone = (setting.arguments as Map<String,dynamic>)["phone"];
              return ConfirmationScreen(phone);
            });
          }
        }
      },
    );
  }
}

