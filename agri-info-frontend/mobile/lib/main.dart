import 'package:flutter_bloc/flutter_bloc.dart';
import "libs.dart";

// class
void main() {
  runApp(MultiBlocProvider(providers: [
    BlocProvider<AuthBloc>(
      create: (context) => AuthBloc(
        AuthRepository(AuthDataProvider()),
      ),
    ),
    BlocProvider<ProductsBloc>(create: (context) {
      return ProductsBloc(ProductsRepository(ProductsProvider()));
    }),
    BlocProvider<MessagesBloc>(create: (context){
      return MessagesBloc(MessagesRepository(MessagesProvider()));
    }),
    BlocProvider<IndexBloc>(create: (context){
      return IndexBloc();
    })
  ], child: MyApp()));
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Agri-Net',
      theme: ThemeData(
          // primarySwatch: Colors.black45, //  MaterialColor(primary, swatch),
          ),
      initialRoute: AuthScreen.RouteName,
      onGenerateRoute: (setting) {
        switch (setting.name) {
          case RegistrationScreen.RouteName:
            {
              return MaterialPageRoute(builder: (context) {
                return RegistrationScreen();
              });
            }
          case AuthScreen.RouteName:
            {
              return MaterialPageRoute(builder: (context) {
                return AuthScreen();
              });
            }
          // case SubscriptionScreen.RouteName:
          //   {
          //     return MaterialPageRoute(builder: (context) {
          //       return SubscriptionScreen();
          //     });
          //   }
          case ConfirmationScreen.RouteName:
            {
              return MaterialPageRoute(builder: (context) {
                final String phone =
                    (setting.arguments as Map<String, dynamic>)["phone"];
                final String fullname =
                    (setting.arguments as Map<String, dynamic>)["fullname"];
                bool islogin = ((setting.arguments
                    as Map<String, dynamic>)["islogin"]) as bool;
                return ConfirmationScreen(phone, fullname, islogin: islogin);
              });
            }
          case HomeScreen.RouteName:
            {
              return MaterialPageRoute(builder: (context) {
                return HomeScreen();
              });
            }
        }
      },
    );
  }
}
