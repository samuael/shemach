import '../../libs.dart';

class UsersScreen extends StatefulWidget {
  static String RouteName = "users";
  const UsersScreen({Key? key}) : super(key: key);

  @override
  State<UsersScreen> createState() => _UsersScreenState();
}

class _UsersScreenState extends State<UsersScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          backgroundColor: Theme.of(context).canvasColor,
          elevation: 5,
          toolbarHeight: MediaQuery.of(context).size.height / 13,
          leading: IconButton(
              color: Colors.black,
              onPressed: () {
                Navigator.pop(context);
              },
              icon: BackButton()),
          title: Text(
            "Registered Admins",
            style: TextStyle(
                fontSize: 18, fontWeight: FontWeight.bold, color: Colors.black),
          ),
        ),
        body: Column(
          children: [
            Container(
              decoration:
                  BoxDecoration(borderRadius: BorderRadius.circular(10)),
              child: Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  IconButton(
                      onPressed: () {
                        context.read<AdminsBloc>().add(CreateNewUserEvent());
                      },
                      icon: Icon(Icons.add)),
                  Text(
                    "New",
                    style: TextStyle(
                        fontSize: 20,
                        fontWeight: FontWeight.bold,
                        color: Colors.black),
                  ),
                ],
              ),
            ),
            BlocBuilder<AdminsBloc, UsersState>(builder: (context, state) {
              if (state is GetAllUsersEvent) {
                return Center(
                  child: CircularProgressIndicator(),
                );
              }
              if (state is AllUsersRetrievedState) {
                return Positioned(
                  child: userRow(state.usersList),
                );
              }
              if (state is NoUserFoundState) {
                return Center(child: Text("No User is registered yet!"));
              }
              return Center(
                child: Text("Some thing went wrong"),
              );
            }),
          ],
        ));
  }

  Widget userRow(List<User> users) {
    return Flexible(
      child: ListView.builder(
          itemCount: users.length,
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
                        Row(
                          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                          children: [
                            CircleAvatar(
                              child: ClipOval(
                                child: Image.asset(
                                  users[counter].imgurl,
                                  width: 70,
                                  height: 70,
                                  fit: BoxFit.cover,
                                ),
                              ),
                            ),
                            SizedBox(
                              width: 7.5,
                            ),
                            Text(users[counter].firstname),
                          ],
                        ),
                        Text(StaticDataStore.ROLE),
                        IconButton(onPressed: () {}, icon: Icon(Icons.edit))
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
