import '../../libs.dart';

class HomeScreen extends StatelessWidget {
  static const String RouteName = "homescreen";
  HomeScreen();

  @override
  Widget build(BuildContext context) {
    final productTypeProvider = BlocProvider.of<ProductTypeBloc>(context);
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
      body: Stack(
        children: [
          context.watch<IndexBloc>().state == 0
              ? ProductPostsList()
              : NotificationsScreen(),
          CollapsingSideBarDrawer(),
        ],
      ),
    );
  }
}
