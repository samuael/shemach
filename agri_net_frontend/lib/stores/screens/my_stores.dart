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
                          return Card(
                            child: Padding(
                              padding: const EdgeInsets.all(10.0),
                              child: Column(
                                mainAxisAlignment:
                                    MainAxisAlignment.spaceEvenly,
                                children: [
                                  Row(
                                    children: [
                                      Text("Store : ",
                                          style: TextStyle(
                                              fontSize: 20,
                                              fontWeight: FontWeight.bold)),
                                      Text(
                                        state.myStores[counter].storeName,
                                        style: TextStyle(
                                          fontSize: 20,
                                        ),
                                      )
                                    ],
                                  ),
                                  Row(
                                    children: [
                                      Text("Owner : ",
                                          style: TextStyle(
                                              fontSize: 18,
                                              fontWeight: FontWeight.bold)),
                                      Text(
                                        widget.user.firstname +
                                            " " +
                                            widget.user.lastname,
                                        style: TextStyle(
                                          fontSize: 18,
                                        ),
                                      )
                                    ],
                                  ),
                                  Row(
                                    children: [
                                      Text("Active Products : ",
                                          style: TextStyle(
                                              fontSize: 18,
                                              fontWeight: FontWeight.bold)),
                                      Text(
                                        state.myStores[counter].activeProducts
                                            .toString(),
                                        style: TextStyle(
                                          fontSize: 18,
                                        ),
                                      )
                                    ],
                                  ),
                                  Row(
                                    children: [
                                      Text("Active Contracts : ",
                                          style: TextStyle(
                                              fontSize: 18,
                                              fontWeight: FontWeight.bold)),
                                      Text(
                                        state.myStores[counter].activeContracts
                                            .toString(),
                                        style: TextStyle(
                                          fontSize: 18,
                                        ),
                                      )
                                    ],
                                  ),
                                  AddressView(state.myStores[counter].address)
                                ],
                              ),
                            ),
                          );
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
