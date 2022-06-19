import '../../libs.dart';

class RegisteredAgentsScreen extends StatefulWidget {
  static String RouteName = "agents";
  const RegisteredAgentsScreen({Key? key}) : super(key: key);

  @override
  State<RegisteredAgentsScreen> createState() => _RegisteredAgentsScreenState();
}

class _RegisteredAgentsScreenState extends State<RegisteredAgentsScreen> {
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
            " Agents ",
            style: TextStyle(
                fontSize: 18, fontWeight: FontWeight.bold, color: Colors.black),
          ),
          centerTitle: true,
        ),
        body: Padding(
          padding: const EdgeInsets.fromLTRB(10, 0, 10, 10),
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
                  BlocBuilder<AgentsBloc, AgentsState>(
                      builder: (context, state) {
                    if (state is AgentsLoadingState) {
                      print("AgentsLoadingState");
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
                    if (state is AgentsLoadedState) {
                      print("AgentsLoadedState");
                      return Column(
                        children: [
                          Flexible(
                            child: Column(
                              children: [
                                agentRow(state.agentsList),
                              ],
                            ),
                          ),
                        ],
                      );
                    }
                    print("No Agents Instance found");
                    return Center(
                      child: Column(
                        children: [
                          Text("No Agents Instance found"),
                          IconButton(
                            icon: Icon(
                              Icons.replay,
                              color: Colors.blue,
                            ),
                            onPressed: () {
                              context.read<AgentsBloc>().add(LoadMyAgentsEvent(
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
                          MaterialPageRoute(
                              builder: (context) => RegisterAgentForm()),
                        );
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

  Widget agentRow(List<Agent> agents) {
    return Flexible(
      child: ListView.builder(
          itemCount: agents.length,
          itemBuilder: (context, counter) {
            return Padding(
              padding: const EdgeInsets.fromLTRB(5, 3, 5, 3),
              child: Material(
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
                  child: Padding(
                    padding: const EdgeInsets.all(5),
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        InkWell(
                          onTap: (() => Navigator.push(
                                context,
                                MaterialPageRoute(
                                  builder: (context) => UserProfileScreen(
                                    requestedUser: agents[counter],
                                  ),
                                ),
                              )),
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                            children: [
                              CircleAvatar(
                                child: ClipOval(
                                  child: (agents[counter].imgurl != '')
                                      ? Image.network(
                                          agents[counter].imgurl,
                                          width: 70,
                                          height: 70,
                                          fit: BoxFit.cover,
                                        )
                                      : Icon(
                                          Icons.person,
                                        ),
                                ),
                              ),
                              SizedBox(
                                width: 7.5,
                              ),
                              Text(
                                  agents[counter].firstname +
                                      " " +
                                      agents[counter].firstname,
                                  style: TextStyle(
                                    fontSize: 16,
                                  )),
                            ],
                          ),
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
                                      child:
                                          deleteAgentDialog(agents[counter].id),
                                    ),
                                  );
                                },
                              );
                            },
                            icon: Icon(
                              Icons.delete_forever,
                              size: 30,
                              color: Theme.of(context).primaryColor,
                            ))
                      ],
                    ),
                  ),
                ),
              ),
            );
          }),
    );
  }

  Widget deleteAgentDialog(int agentId) {
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
                "Are You Sure To Delete This Agent?",
                style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
              onTap: () {},
            ),
            Row(
              mainAxisAlignment: MainAxisAlignment.spaceEvenly,
              children: [
                ElevatedButton(
                    style: ButtonStyle(
                      backgroundColor: MaterialStateProperty.all<Color>(
                          Theme.of(context).primaryColor),
                      animationDuration: Duration(seconds: 1),
                      padding: MaterialStateProperty.all<EdgeInsets>(
                        EdgeInsets.symmetric(
                          horizontal: 20,
                          vertical: 10,
                        ),
                      ),
                      elevation: MaterialStateProperty.all<double>(0),
                    ),
                    onPressed: () {
                      Navigator.pop(context, true);
                    },
                    child: Text("Cancel")),
                ElevatedButton(
                    style: ButtonStyle(
                      backgroundColor:
                          MaterialStateProperty.all<Color>(Colors.red),
                      animationDuration: Duration(seconds: 1),
                      padding: MaterialStateProperty.all<EdgeInsets>(
                        EdgeInsets.symmetric(
                          horizontal: 20,
                          vertical: 10,
                        ),
                      ),
                      elevation: MaterialStateProperty.all<double>(0),
                    ),
                    onPressed: () {},
                    child: Text("Delete"))
              ],
            )
          ],
        ),
      ),
    );
  }
}
