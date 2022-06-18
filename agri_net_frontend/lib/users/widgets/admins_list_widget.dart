import '../../libs.dart';

class AdminsListView extends StatelessWidget {
  final List<Admin> admins;
  AdminsListView(this.admins);

  @override
  Widget build(BuildContext context) {
    return Expanded(
      child: ListView.builder(
          itemCount: admins.length,
          itemBuilder: (context, counter) {
            return Material(
              elevation: 10,
              child: Container(
                decoration: BoxDecoration(),
                child: Padding(
                  padding: const EdgeInsets.all(5),
                  child: Card(
                    elevation: 0.0,
                    child: Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        InkWell(
                          onTap: (() => Navigator.push(
                                context,
                                MaterialPageRoute(
                                  builder: (context) => UserProfileScreen(
                                    requestedUser: admins[counter],
                                  ),
                                ),
                              )),
                          child: Row(
                            mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                            children: [
                              CircleAvatar(
                                child: ClipOval(
                                  child: (admins[counter].imgurl != '')
                                      ? Image.asset(
                                          admins[counter].imgurl,
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
                              Text(admins[counter].firstname),
                              SizedBox(
                                width: 7.5,
                              ),
                              Text(admins[counter].lastname)
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
