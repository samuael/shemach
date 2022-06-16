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
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        elevation: 0,
        backgroundColor: Theme.of(context).canvasColor,
        toolbarHeight: MediaQuery.of(context).size.height / 13,
        leading: IconButton(
            color: Colors.black,
            onPressed: () {
              Navigator.pop(context);
            },
            icon: BackButton()),
        title: searchBar(),
        actions: [
          Row(
            children: [
              IconButton(
                onPressed: () {
                  Navigator.push(
                      context,
                      new MaterialPageRoute(
                          builder: (context) => new RegisterMerchantForm()));
                },
                icon: Icon(Icons.add),
                color: Colors.black,
              ),
              Padding(
                padding: const EdgeInsets.fromLTRB(0, 0, 20, 0),
                child: InkWell(
                  onTap: () {
                    Navigator.push(
                        context,
                        new MaterialPageRoute(
                            builder: (context) => new RegisterMerchantForm()));
                  },
                  child: Text(
                    "New",
                    style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.bold,
                        color: Colors.black),
                  ),
                ),
              ),
            ],
          ),
        ],
      ),
      body: BlocBuilder<AdminsBloc, AdminsState>(builder: (context, state) {
        if (state is GetAllMerchantsState) {
          print("GetAllMerchantsState");
          return Center(
            child: CircularProgressIndicator(),
          );
        }
        if (state is NoMerchantsFoundState) {
          print("NoMerchantsFoundState");
          return Center(child: Text(state.msg));
        }
        if (state is FailedToFechMerchntsState) {
          print("FailedToFechMerchntsState");
          return Center(
            child: Text(state.msg),
          );
        }
        if (state is AllMerchantsFechedState) {
          print("AllMerchantsFechedState");
          return Column(
            children: [
              topBarOfMerchantsList(context),
              Expanded(child: merchantRow(state.merchants)),
            ],
          );
        }
        return Container(
          child: Center(
            child: CircularProgressIndicator(),
          ),
        );
      }),
    );
  }

  Widget searchBar() {
    return Container(
      padding: EdgeInsets.symmetric(
        horizontal: 25,
      ),
      decoration: BoxDecoration(
        border: Border.all(color: Theme.of(context).primaryColor),
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
            color: Theme.of(context).primaryColor,
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
          BlocProvider.of<StoreBloc>(context)
              .add(MyStoresEvent(ownerId: merchants[counter].id));
          return Padding(
            padding: const EdgeInsets.fromLTRB(5, 3, 5, 3),
            child: Material(
              elevation: 5,
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
                          Container(
                            width: MediaQuery.of(context).size.width / 3,
                            child: Row(
                              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
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
                                Text(merchants[counter].firstname,
                                    style: TextStyle(
                                        fontSize: 16,
                                        fontWeight: FontWeight.bold)),
                              ],
                            ),
                          ),
                          Container(
                              width: MediaQuery.of(context).size.width / 4 - 19,
                              child: Center(
                                  child: Text(merchants[counter]
                                      .storeCount
                                      .toString()))),
                          Container(
                              width: MediaQuery.of(context).size.width / 4 - 19,
                              child: Center(
                                  child: Text(merchants[counter]
                                      .postsCounte
                                      .toString()))),
                          Container(
                            width: MediaQuery.of(context).size.width / 4 - 19,
                            child: Center(
                              child: IconButton(
                                  onPressed: () {},
                                  icon: Icon(Icons.delete_forever)),
                            ),
                          )
                        ],
                      )),
                ),
              ),
            ),
          );
        });
  }
}
