import "libs.dart" ;


void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Agri-Net',
      theme: ThemeData(
        primarySwatch: Colors.tikur_arenguade, //  MaterialColor(primary, swatch),
      ),
      initialRoute: AuthScreen.RouteName,
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

