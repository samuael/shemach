import 'package:agri_net_frontend/profile/widgets/address.dart';

import '../../libs.dart';

class MyStoresScreen extends StatefulWidget {
  static String RouteName = "/stores";

  MyStoresScreen();

  @override
  State<MyStoresScreen> createState() => _MyStoresScreenState();
}

class _MyStoresScreenState extends State<MyStoresScreen> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).primaryColor,
        elevation: 0,
        title: Text(
          "Stores",
          style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
        ),
      ),
      body: SafeArea(
        child: Padding(
          padding: EdgeInsets.all(10),
          child: BlocBuilder<StoreBloc, StoreState>(
            builder: (context, state) {
              if (context.watch<StoreBloc>().state is MyStoresLoadedState) {
                return Flexible(
                    child: (context.watch<StoreBloc>().state
                                    as MyStoresLoadedState)
                                .myStores[StaticDataStore.ID] !=
                            null
                        ? ListView.builder(
                            itemCount: (context.watch<StoreBloc>().state
                                    as MyStoresLoadedState)
                                .myStores[StaticDataStore.ID]!
                                .length,
                            itemBuilder: (context, counter) {
                              return StoreView((context.watch<StoreBloc>().state
                                      as MyStoresLoadedState)
                                  .myStores[StaticDataStore.ID]![counter]);
                            })
                        : Center());
              }

              return Container();
            },
          ),
        ),
      ),
    );
  }
}
