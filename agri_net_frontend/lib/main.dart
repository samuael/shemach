import 'package:agri_net_frontend/admins/widgets/register_admin_form.dart';

import 'libs.dart';

void main() {
  runApp(MultiBlocProvider(providers: [
    BlocProvider<UserBloc>(
      create: (context) => UserBloc(repo: UserRepo(provider: AuthProvider())),
    ),
    BlocProvider(create: (context) {
      return AdminsBloc(adminsRepo: AdminsRepo(adminsProvider: UserProvider()))
        ..add(GetAllAdminsEvent());
    }),
    BlocProvider(create: (context) {
      return StoreBloc(storeRepo: StoreRepo(storeProvider: StoreProvider()));
    }),
    BlocProvider(create: (context) {
      return ProductBloc(
        repo: ProductRepo(provider: ProductProvider()),
      )..add(GetProductListEvent());
    }),
  ], child: MyHomePage()));
}

class MyHomePage extends StatefulWidget {
  // static int once = 0;
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
          primaryColor: Color.fromARGB(
              255, 11, 100, 14), //  MaterialColor(primary, swatch),
        ),
        initialRoute: AuthScreen.RouteName,
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
          } else if (route == RegisteredMerchantsScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return RegisteredMerchantsScreen();
            });
          } else if (route == RegisteredAgentsScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return RegisteredAgentsScreen();
            });
          } else if (route == HomeScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return HomeScreen();
            });
          } else if (route == AdminsScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return AdminsScreen();
            });
          } else if (route == RegisterAdminPage.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return RegisterAdminPage();
            });
          } else if (route == UserProfileScreen.RouteName) {
            User args = setting.arguments as User;
            return MaterialPageRoute(builder: (context) {
              return UserProfileScreen(
                requestedUser: args,
              );
            });
          }
        });
  }
}
