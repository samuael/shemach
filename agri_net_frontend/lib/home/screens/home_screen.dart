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

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).canvasColor,
        toolbarHeight: MediaQuery.of(context).size.height / 13,
        leading: AgriNetLogo(),
        title: UserScreenAppBarDrawer(),
      ),
      body: CollapsingSideBarDrawer(),
    );
  }
}
