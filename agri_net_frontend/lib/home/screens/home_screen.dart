import '../../libs.dart';

class HomeScreen extends StatelessWidget {
  static const String RouteName = "homescreen";
  HomeScreen();

  @override
  Widget build(BuildContext context) {
    final productTypeProvider = BlocProvider.of<ProductTypeBloc>(context);
    final userProvider = BlocProvider.of<UserBloc>(context);
    final adminProvider = BlocProvider.of<AdminsBloc>(context);

    if (productTypeProvider.state is ProductTypeInit) {
      productTypeProvider.add(ProductTypesLoadEvent());
    }

    if (StaticDataStore == ROLE_AGENT ||
        StaticDataStore.ROLE == ROLE_MERCHANT) {
      final myProductsBlocProvider = BlocProvider.of<MyProductsBloc>(context);
      if (myProductsBlocProvider.state is MyProductInit) {
        myProductsBlocProvider.add(LoadMyProductsEvent());
      }
    }

    if (userProvider.state is Authenticated) {
      final theUser = (userProvider.state as Authenticated).user;
      if (StaticDataStore.ROLE == ROLE_SUPERADMIN) {
        adminProvider.add(GetAllAdminsEvent());
      }
      if (theUser is Admin) {
        final theAdmin = theUser as Admin;
        adminProvider.add(GetAllAgentsEvent(admin: theAdmin));
        adminProvider.add(GetAllMerchantsEvent(admin: theAdmin));
      }
    }

    final productsPostProvider = BlocProvider.of<ProductsBloc>(context);
    if (!(productsPostProvider.state is ProductsLoadSuccess)) {
      productsPostProvider.add(LoadProductsEvent());
    }
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).canvasColor,
        toolbarHeight: MediaQuery.of(context).size.height / 13,
        leading: AgriNetLogo(),
        title: UserScreenAppBarDrawer(),
      ),
      body: Row(children: [
        CollapsingSideBarDrawer(),
        ProductPostsList(),
      ]),
    );
  }
}
