import '../../libs.dart';

class RegisteredMerchantsScreen extends StatefulWidget {
  static String RouteName = "merchants";
  const RegisteredMerchantsScreen({Key? key}) : super(key: key);

  @override
  State<RegisteredMerchantsScreen> createState() =>
      _RegisteredMerchantsScreenState();
}

class _RegisteredMerchantsScreenState extends State<RegisteredMerchantsScreen> {
  TextEditingController searchController = TextEditingController();

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final storeProvider = BlocProvider.of<StoreBloc>(context);
    return Scaffold(
        appBar: AppBar(
          elevation: 0,
          backgroundColor: Theme.of(context).primaryColor,
          toolbarHeight: MediaQuery.of(context).size.height / 13,
          title: Text("Merchants",
              style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
          leading: IconButton(
              color: Colors.black,
              onPressed: () {
                Navigator.pop(context);
              },
              icon: BackButton()),
        ),
        body: Padding(
          padding: const EdgeInsets.fromLTRB(10, 0, 10, 10),
          child: Stack(
            children: [
              BlocBuilder<MercahntsBloc, MerchantsState>(
                  builder: (context, state) {
                if (state is MerchantsLoadingState) {
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
                if (state is MerchantsLoadedState) {
                  return Column(
                    children: [
                      // topBarOfMerchantsList(context),
                      Expanded(child: merchantRow(state.merchants)),
                    ],
                  );
                }
                return Center(
                  child: Column(
                    children: [
                      Text("No Merchant Instance found"),
                      IconButton(
                        icon: Icon(
                          Icons.replay,
                          color: Colors.blue,
                        ),
                        onPressed: () {
                          context.read<MercahntsBloc>().add(
                              LoadMyMerchantsEvent(
                                  adminID: StaticDataStore.ID));
                        },
                      )
                    ],
                  ),
                );
              }),
              Align(
                alignment: Alignment.bottomRight,
                child: FloatingActionButton(
                  backgroundColor: Theme.of(context).canvasColor,
                  elevation: 5.0,
                  onPressed: () {
                    Navigator.push(
                        context,
                        new MaterialPageRoute(
                            builder: (context) => new RegisterMerchantForm()));
                  },
                  child: Icon(
                    Icons.add,
                    color: Theme.of(context).primaryColor,
                    size: 50,
                  ),
                ),
              )
            ],
          ),
        ));
  }

  Widget searchBar() {
    return Container(
      padding: EdgeInsets.symmetric(
        horizontal: 25,
      ),
      decoration: BoxDecoration(
        border: Border.all(color: Colors.black),
        borderRadius: BorderRadius.circular(10),
      ),
      child: TextField(
        onChanged: (t) {
          setState(() {
            // widget.text = t;
          });
        },
        autofocus: true,
        style: TextStyle(
          fontWeight: FontWeight.bold,
        ),
        // controller: searchController,
        decoration: InputDecoration(
          border: InputBorder.none,
          suffixIcon: Icon(
            Icons.search,
            color: Colors.black,
          ),
        ),
      ),
    );
  }

  Widget topBarOfMerchantsList(BuildContext context) {
    return Padding(
      padding: EdgeInsets.fromLTRB(5, 5, 5, 3),
      child: Material(
        elevation: 10,
        child: Container(
          child: Padding(
            padding: EdgeInsets.all(5),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                Container(
                  width: MediaQuery.of(context).size.width / 3,
                  child: Center(
                    child: Text(
                      "Merchant",
                      style:
                          TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
                    ),
                  ),
                ),
                Container(
                  width: MediaQuery.of(context).size.width / 4 - 19,
                  child: Center(
                    child: Row(
                      children: [
                        Text("Stores",
                            style: TextStyle(
                                fontSize: 20, fontWeight: FontWeight.bold)),
                      ],
                    ),
                  ),
                ),
                Container(
                  width: MediaQuery.of(context).size.width / 4 - 19,
                  child: Center(
                    child: Text("Posts",
                        style: TextStyle(
                            fontSize: 20, fontWeight: FontWeight.bold)),
                  ),
                ),
                Container(
                  width: MediaQuery.of(context).size.width / 4 - 19,
                  child: Center(
                    child: Text("Remove",
                        style: TextStyle(
                            fontSize: 20, fontWeight: FontWeight.bold)),
                  ),
                )
              ],
            ),
          ),
        ),
      ),
    );
  }

  Widget merchantRow(List<Merchant> merchants) {
    return ListView.builder(
        itemCount: merchants.length,
        itemBuilder: (context, counter) {
          return Material(
            elevation: 2.5,
            child: Container(
              decoration: BoxDecoration(),
              child: Padding(
                padding: const EdgeInsets.all(5),
                child: InkWell(
                    onTap: (() => Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) => UserProfileScreen(
                              requestedUser: merchants[counter],
                            ),
                          ),
                        )),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        Row(
                          children: [
                            CircleAvatar(
                              child: ClipOval(
                                child: (merchants[counter].imgurl != '')
                                    ? Image.asset(
                                        merchants[counter].imgurl,
                                        width: 70,
                                        height: 70,
                                        fit: BoxFit.cover,
                                      )
                                    : Icon(
                                        Icons.person,
                                      ),
                              ),
                            ),
                            SizedBox(width: 5),
                            Text(
                                merchants[counter].firstname +
                                    " " +
                                    merchants[counter].firstname,
                                style: TextStyle(
                                    fontSize: 16, fontWeight: FontWeight.bold)),
                          ],
                        ),
                        IconButton(
                            onPressed: () {
                              showDialog(
                                barrierDismissible: false,
                                context: context,
                                builder: (BuildContext cxt) {
                                  return Padding(
                                    padding: EdgeInsets.fromLTRB(0, 50, 0, 10),
                                    child: Dialog(
                                      shape: RoundedRectangleBorder(
                                          borderRadius: BorderRadius.only(
                                              bottomLeft: Radius.circular(10),
                                              bottomRight: Radius.circular(10),
                                              topLeft: Radius.circular(10),
                                              topRight: Radius.circular(10))),
                                      child: editMerchantDialog(
                                          merchants[counter].id),
                                    ),
                                  );
                                },
                              );
                            },
                            icon: Icon(
                              Icons.edit,
                              size: 30,
                              color: Theme.of(context).primaryColor,
                            ))
                      ],
                    )),
              ),
            ),
          );
        });
  }

  Widget editMerchantDialog(int merchantId) {
    return Container(
      width: MediaQuery.of(context).size.width * 0.15,
      height: MediaQuery.of(context).size.height * 0.15,
      child: Padding(
        padding: const EdgeInsets.fromLTRB(5, 2.5, 5, 2.5),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            ListTile(
              title: Text(
                "Add Store",
                style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
              onTap: () {
                Navigator.push(context, MaterialPageRoute(builder: (context) {
                  return NewStoreForm(ownerID: merchantId);
                }));
              },
            ),
            ListTile(
              title: Text(
                "Remove",
                style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
              onTap: () {},
            ),
          ],
        ),
      ),
    );
  }
}
