import '../../libs.dart';

class HomeScreen extends StatefulWidget {
  static const String RouteName = "homescreen";
  HomeScreen();

  @override
  State<HomeScreen> createState() {
    return HomeScreenState();
  }
}

class HomeScreenState extends State<HomeScreen> {
  TransactionBloc? transactionProvider;
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

    if (StaticDataStore.ROLE == ROLE_MERCHANT ||
        StaticDataStore.ROLE == ROLE_AGENT) {
      transactionProvider = BlocProvider.of<TransactionBloc>(context);
      transactionProvider?.add(TransactionLoadEvent());
      transactionProvider?.startLoadTransactionsLoop();
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

  @override
  void dispose() {
    super.dispose();
    if (this.transactionProvider != null) {
      this.transactionProvider!.stopLoop();
    }
  }
}
