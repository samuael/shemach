import 'package:flutter/services.dart';

import 'libs.dart';

void main() {
  WidgetsFlutterBinding.ensureInitialized();
  // Step 3
  SystemChrome.setPreferredOrientations([
    DeviceOrientation.portraitUp,
    DeviceOrientation.portraitDown,
  ]).then(
    (value) => runApp(
      MultiBlocProvider(
        providers: [
          BlocProvider<UserBloc>(
            create: (context) =>
                UserBloc(repo: UserRepo(provider: AuthProvider())),
          ),
          BlocProvider(create: (context) {
            return AdminsBloc(
                adminsRepo: AdminsRepo(adminsProvider: UserProvider()))
              ..add(GetAllAdminsEvent());
          }),
          BlocProvider(create: (context) {
            return ProductTypeBloc(
              ProductTypesRepository(ProductTypesProvider()),
            );
          }),
          BlocProvider(
            create: (context) {
              return StoreBloc(
                  storeRepo: StoreRepo(storeProvider: StoreProvider()));
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
            return IndexBloc();
          }),
          BlocProvider(create: (context) {
            return ProductsBloc(repo: ProductsRepo(ProductProvider()));
          }),
          BlocProvider(create: (context) {
            return UsersBloc(UsersRepo(UsersProvider()));
          }),
          BlocProvider(create: (context) {
            return TransactionBloc(TransactionRepo(TransactionProvider()));
          })
        ],
        child: MyHomePage(),
      ),
    ),
  );
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
          } else if (route == CreateTransactionScreen.RouteName) {
            ProductPost post =
                (setting.arguments as Map<String, ProductPost>)["post"]!;
            return MaterialPageRoute(builder: (context) {
              return CreateTransactionScreen(post);
            });
          } else if (route == RegisterAdminPage.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return RegisterAdminPage();
            });
          } else if (route == ConfirmationScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return ConfirmationScreen();
            });
          } else if (route == NotificationsScreen.RouteName) {
            return MaterialPageRoute(builder: (context) {
              return NotificationsScreen();
            });
          } else if (route == UserProfileScreen.RouteName) {
            User args = setting.arguments as User;
            return MaterialPageRoute(builder: (context) {
              return UserProfileScreen(
                requestedUser: args,
              );
            });
          } else if (route == UploadProductPostImages.RouteName) {
            ProductPost post =
                (setting.arguments as Map<String, dynamic>)["post"];
            return MaterialPageRoute(builder: (context) {
              return UploadProductPostImages(post);
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
