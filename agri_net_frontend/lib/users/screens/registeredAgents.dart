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
          backgroundColor: Theme.of(context).primaryColor,
          elevation: 5,
          toolbarHeight: MediaQuery.of(context).size.height / 13,
          title: Text("Agents",
              style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
          leading: IconButton(
              color: Colors.black,
              onPressed: () {
                Navigator.pop(context);
              },
              icon: BackButton()),
        ),
        body: Stack(
          children: [
            BlocBuilder<AgentsBloc, AgentsState>(builder: (context, state) {
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
                        context.read<AgentsBloc>().add(
                            LoadMyAgentsEvent(adminID: StaticDataStore.ID));
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
                  color: Colors.black,
                  size: 35,
                ),
              ),
            )
          ],
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
                elevation: 5,
                child: Container(
                  decoration: BoxDecoration(),
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
                                      ? Image.asset(
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
                              Text(agents[counter].firstname),
                              SizedBox(
                                width: 7.5,
                              ),
                              Text(agents[counter].firstname)
                            ],
                          ),
                        ),
                        IconButton(onPressed: () {}, icon: Icon(Icons.delete))
                      ],
                    ),
                  ),
                ),
              ),
            );
          }),
    );
  }
}
