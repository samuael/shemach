import 'package:path/path.dart';

import 'libs.dart';

void main() {
  runApp(MultiBlocProvider(providers: [
    BlocProvider<UserBloc>(
      create: (context) => UserBloc(repo: UserRepo(provider: AuthProvider())),
    ),
    BlocProvider(create: (context) {
      return AdminsBloc(
          adminsRepo: AdminsRepo(adminsProvider: AdminProvider()));
    }),
    BlocProvider(create: (context) {
      return AgentsBloc(agentRepo: AgentRepo(agentProvider: AgentProvider()));
    }),
    BlocProvider(create: (context) {
      return MercahntsBloc(
          merchantRepo: MerchantRepo(merchantProvider: MerchantProvider()));
    }),
    BlocProvider(create: (context) {
      return ProductTypeBloc(
        ProductTypesRepository(ProductTypesProvider()),
      );
    }),
    BlocProvider(
      create: (context) {
        return StoreBloc(storeRepo: StoreRepo(storeProvider: StoreProvider()));
      },
    ),
    BlocProvider(create: (context) {
      return MyProductsBloc(
        ProductsRepo(
          ProductProvider(),
        ),
      );
    }),
    BlocProvider(create: (context) {
      return ProductsBloc(repo: ProductsRepo(ProductProvider()));
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
        debugShowCheckedModeBanner: false,
        title: 'Agri-Net',
        theme: ThemeData(
          primaryColor: Colors.green, //  MaterialColor(primary, swatch),
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
          } else if (route == StoreSelectionScreen.RouteName) {
            final arguments = setting.arguments as Map<String, dynamic>;
            // ProductTypeState state = (arguments["state"]) as ProductTypeState;
            List<Store> stores = arguments["stores"] as List<Store>;
            Function callBack = (arguments["callback"] as Function);
            return MaterialPageRoute(builder: (context) {
              return StoreSelectionScreen(stores, callBack);
            });
          } else if (route == ProductPostDetailScreen.RouteName) {
            final ProductPost post = ((setting.arguments
                as Map<String, dynamic>)["post"] as ProductPost);
            return MaterialPageRoute(builder: (context) {
              return ProductPostDetailScreen(post);
            });
          } else if (route == ProductTypeSelectionScreen.RouteName) {
            final arguments = setting.arguments as Map<String, dynamic>;
            ProductTypeState state = (arguments["state"]) as ProductTypeState;
            List<ProductType> products =
                arguments["products"] as List<ProductType>;
            Function callBack = (arguments["callback"] as Function);
            String text = (arguments["text"] as String);

            return MaterialPageRoute(builder: (context) {
              return ProductTypeSelectionScreen(
                  state, products, callBack, text);
            });
          }
        });
  }
}
