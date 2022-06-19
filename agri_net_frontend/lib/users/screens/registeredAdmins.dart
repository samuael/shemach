import 'package:agri_net_frontend/users/widgets/admins_list_widget.dart';

import '../../libs.dart';

class AdminsScreen extends StatefulWidget {
  static String RouteName = "admins";
  const AdminsScreen({Key? key}) : super(key: key);

  @override
  State<AdminsScreen> createState() => _AdminsScreenState();
}

class _AdminsScreenState extends State<AdminsScreen> {
  bool adminsList = true;
  @override
  void initState() {
    super.initState();
  }

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
            icon: BackButton(),
          ),
          title: Text(
            " Admins ",
            style: TextStyle(
                fontSize: 18, fontWeight: FontWeight.bold, color: Colors.black),
          ),
          centerTitle: true,
        ),
        body: Padding(
          padding: EdgeInsets.fromLTRB(10, 0, 10, 10),
          child: ClipRRect(
            borderRadius: BorderRadius.circular(10),
            child: Container(
              margin: EdgeInsets.symmetric(
                vertical: 3,
              ),
              decoration: BoxDecoration(
                border: Border.all(
                  color: Theme.of(context).primaryColorLight,
                ),
                borderRadius: BorderRadius.circular(10),
              ),
              child: Stack(
                children: [
                  BlocBuilder<AdminsBloc, AdminsState>(
                      builder: (context, state) {
                    if (state is AdminsStateInIt) {
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
                    if (state is AdminsLoadedState) {
                      return Column(
                        children: [
                          Flexible(
                            child: Column(
                              children: [
                                AdminsListView(state.adminsList),
                              ],
                            ),
                          ),
                        ],
                      );
                      ;
                    }
                    if (state is AdminsLoadingFailed) {
                      return Center(
                        child: Column(
                          children: [
                            Text("Sorry : ${state.msg}"),
                            IconButton(
                              icon: Icon(
                                Icons.replay,
                                color: Colors.blue,
                              ),
                              onPressed: () {
                                context
                                    .read<AdminsBloc>()
                                    .add(LoadAdminsInIt());
                              },
                            )
                          ],
                        ),
                      );
                    }
                    return Center(
                      child: Column(
                        children: [
                          Text("No Admins Instance found"),
                          IconButton(
                            icon: Icon(
                              Icons.replay,
                              color: Colors.blue,
                            ),
                            onPressed: () {
                              context.read<AdminsBloc>().add(LoadAdminsInIt());
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
                        Navigator.of(context)
                            .pushNamed(RegisterAdminPage.RouteName);
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
            ),
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
}
