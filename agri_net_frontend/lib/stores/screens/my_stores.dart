import 'dart:developer';

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
  void initState() {
    super.initState();
    SchedulerBinding.instance?.addPostFrameCallback((_) {
      BlocProvider.of<StoreBloc>(context)
          .add(MyStoresEvent(ownerId: widget.user.id));
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
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
                            child: Column(
                              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                              children: [
                                Row(
                                  children: [
                                    Text("Store : "),
                                    Text(
                                      state.myStores[counter].storeName,
                                      style: TextStyle(
                                          fontSize: 18,
                                          fontWeight: FontWeight.bold),
                                    )
                                  ],
                                ),
                                Row(
                                  children: [
                                    Text("Owner : "),
                                    Text(
                                      widget.user.firstname +
                                          " " +
                                          widget.user.lastname,
                                      style: TextStyle(
                                          fontSize: 10,
                                          fontWeight: FontWeight.bold),
                                    )
                                  ],
                                ),
                                ExpansionTile(
                                  title: Text("Address"),
                                  children: [
                                    Column(
                                      crossAxisAlignment:
                                          CrossAxisAlignment.start,
                                      children: [
                                        (state.myStores[counter].address
                                                    .UniqueAddressName !=
                                                '')
                                            ? Text(
                                                "Unique : ${state.myStores[counter].address.UniqueAddressName}:")
                                            : SizedBox(),
                                        (state.myStores[counter].address
                                                    .Kebele !=
                                                '')
                                            ? Text(
                                                "Kebele : ${state.myStores[counter].address.Kebele}:")
                                            : SizedBox(),
                                        (state.myStores[counter].address
                                                    .Woreda !=
                                                '')
                                            ? Text(
                                                "Woreda : ${state.myStores[counter].address.Woreda}:")
                                            : SizedBox(),
                                        (state.myStores[counter].address.City !=
                                                '')
                                            ? Text(
                                                "City : ${state.myStores[counter].address.City}:")
                                            : SizedBox(),
                                        (state.myStores[counter].address.Zone !=
                                                '')
                                            ? Text(
                                                "Zone : ${state.myStores[counter].address.Zone}:")
                                            : SizedBox(),
                                        (state.myStores[counter].address
                                                    .Region !=
                                                '')
                                            ? Text(
                                                "Region : ${state.myStores[counter].address.Region}:")
                                            : SizedBox(),
                                        IconButton(
                                            onPressed: () {},
                                            icon: Icon(
                                                Icons.maps_home_work_sharp))
                                      ],
                                    ),
                                  ],
                                )
                              ],
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
