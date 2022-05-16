import 'package:agri_net_frontend/contracts/screens/screens.dart';

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
    ),
    BlocProvider(create: (context) {
      return ProductBloc(
        repo: ProductRepo(provider: ProductProvider()),
      );
    }),
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
        initialRoute: HomeScreen.RouteName,
        onGenerateRoute: (setting) {
          final route = setting.name;
          if (route == AuthScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return AuthScreen();
            });
          } else if (route == ProductScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return ProductScreen();
            });
          } else if (route == HomeScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return HomeScreen();
            });
          } else if (route == ProfileScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return ProfileScreen();
            });
          } else if (route == ContractScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return ContractScreen();
            });
          } else if (route == NotificationScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return NotificationScreen();
            });
          }
        });
  }
}
