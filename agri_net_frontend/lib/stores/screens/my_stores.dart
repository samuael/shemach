import 'package:agri_net_frontend/profile/widgets/address.dart';

import '../../libs.dart';

class MyStoresScreen extends StatefulWidget {
  static String RouteName = "/stores";
  Merchant mercahnt;
  MyStoresScreen(this.mercahnt);

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
              if (state is LoadingMyStoresState) {
                return Center(
                  child: Column(
                    children: [
                      CircularProgressIndicator(
                        strokeWidth: 3,
                      ),
                      Text(
                        translate(lang, "loading ..."),
                        style: TextStyle(
                          fontWeight: FontWeight.bold,
                          fontStyle: FontStyle.italic,
                          color: Theme.of(context).primaryColor,
                        ),
                      )
                    ],
                  ),
                );
              }
              if (state is MyStoresLoadedState) {
                if ((state.myStores[widget.mercahnt.id]) != null &&
                    (state.myStores[widget.mercahnt.id])!.length > 0) {
                  final stores = state.myStores[widget.mercahnt.id];
                  return Flexible(
                      child: ListView.builder(
                          itemCount: stores!.length,
                          itemBuilder: (context, counter) {
                            return Text('${stores[counter]}');
                          }));
                }
                return Flexible(
                    child: (state.myStores[widget.mercahnt.id]) != null
                        ? ListView.builder(
                            itemCount:
                                state.myStores[widget.mercahnt.id]!.length,
                            itemBuilder: (context, counter) {
                              return StoreView(
                                  state.myStores[widget.mercahnt.id]![counter]);
                            })
                        : Center());
              }

              return Center(
                child: Column(
                  children: [
                    Text("No Store Instance found"),
                    IconButton(
                      icon: Icon(
                        Icons.replay,
                        color: Colors.blue,
                      ),
                      onPressed: () {
                        context.read<StoreBloc>().add(
                            LoadMyStoresEvent(ownerId: widget.mercahnt.id));
                      },
                    )
                  ],
                ),
              );
            },
          ),
        ),
      ),
    );
  }
}
