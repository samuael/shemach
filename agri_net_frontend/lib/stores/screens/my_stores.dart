import 'package:agri_net_frontend/profile/widgets/address.dart';

import '../../libs.dart';

class MyStoresScreen extends StatefulWidget {
  static String RouteName = "stores";
  User user;
  MyStoresScreen({required this.user});

  @override
  State<MyStoresScreen> createState() => _MyStoresScreenState();
}

class _MyStoresScreenState extends State<MyStoresScreen> {
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
              if (state is MyStoresState) {
                return Flexible(
                    child: ListView.builder(
                        itemCount: state.myStores.length,
                        itemBuilder: (context, counter) {
                          return StoreView(state.myStores[counter]);
                        }));
              }

              return Container();
            },
          ),
        ),
      ),
    );
  }
}
