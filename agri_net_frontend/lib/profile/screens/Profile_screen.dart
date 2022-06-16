import 'package:agri_net_frontend/profile/widgets/address.dart';

import '../../libs.dart';

class UserProfileScreen extends StatefulWidget {
  static String RouteName = "profile";
  User requestedUser;
  UserProfileScreen({required this.requestedUser});

  @override
  State<UserProfileScreen> createState() => _UserProfileScreenState();
}

class _UserProfileScreenState extends State<UserProfileScreen> {
  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    final UserBloc _userBloc = BlocProvider.of<UserBloc>(context);
    User loggedInUser = (_userBloc.state as Authenticated).user;
    return Scaffold(
      appBar: AppBar(
        title: Text("User Profile"),
        backgroundColor: Theme.of(context).primaryColor,
        elevation: 0,
      ),
      body: SafeArea(
        child: BlocProvider(
            create: (Context) {
              return ProfileBLoc(
                  user: widget.requestedUser,
                  isCurrentUser: widget.requestedUser.id == loggedInUser.id,
                  profileRepository:
                      ProfileRepository(profileProvider: ProfileProvider()));
            },
            child: BlocListener<ProfileBLoc, ProfileState>(
              listener: ((context, state) {
                if (state.imageSourceActionSheetIsVisible!) {
                  showImageSource(context);
                }
                if (state.avatorPath != '') {
                  setState(() {
                    print(loggedInUser.imgurl);
                    loggedInUser.imgurl = state.avatorPath!;
                    print(loggedInUser.imgurl);
                  });
                }
              }),
              child: BlocBuilder<ProfileBLoc, ProfileState>(
                builder: ((context, state) => Scaffold(
                      body: SafeArea(
                        child: Padding(
                          padding: const EdgeInsets.all(10.0),
                          child: Material(
                            elevation: 5,
                            child: Container(
                              child: Center(
                                  child: Column(
                                children: [
                                  Stack(children: [
                                    buildImage(widget.requestedUser.imgurl),
                                    (widget.requestedUser.id == loggedInUser.id)
                                        ? Positioned(
                                            child: buildEditIcon(context),
                                            // right: -15,
                                            left: 90,
                                            top: 90,
                                          )
                                        : Positioned(
                                            child: Container(),
                                            right: 4,
                                            top: 10,
                                          ),
                                  ]),
                                  _profilePage(),
                                  (widget.requestedUser is Admin)
                                      ? AddressView(
                                          (widget.requestedUser as Admin)
                                              .address)
                                      : (widget.requestedUser is Merchant)
                                          ? Expanded(
                                              child: Column(children: [
                                                myStores(
                                                    widget.requestedUser
                                                        as Merchant,
                                                    context),
                                                AddressView(
                                                    (widget.requestedUser
                                                            as Merchant)
                                                        .address),
                                              ]),
                                            )
                                          : (widget.requestedUser is Agent)
                                              ? AddressView(
                                                  (widget.requestedUser
                                                          as Agent)
                                                      .address,
                                                )
                                              : Container()
                                ],
                              )),
                            ),
                          ),
                        ),
                      ),
                    )),
              ),
            )),
      ),
    );
  }

  Widget myStores(Merchant merchant, context) {
    final UserBloc _userBloc = BlocProvider.of<UserBloc>(context);
    User loggedInUser = (_userBloc.state as Authenticated).user;
    // BlocProvider.of<StoreBloc>(context)
    //     .add(MyStoresEvent(ownerId: merchant.id));
    return BlocBuilder<StoreBloc, StoreState>(
      builder: (context, state) {
        if (state is MyStoresState) {
          return Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              Expanded(
                  child: InkWell(
                onTap: () {
                  Navigator.push(
                      context,
                      new MaterialPageRoute(
                          builder: (context) => new MyStoresScreen(
                                user: widget.requestedUser,
                              )));
                },
                child: ListTile(
                  leading: Icon(
                    Icons.store,
                    color: Theme.of(context).primaryColor,
                  ),
                  title: Text(
                    "Stores : ${state.myStores.length}",
                    style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
                  ),
                ),
              )),
              Container(
                  width: MediaQuery.of(context).size.width / 4 - 19,
                  child: Row(
                    children: [
                      (widget.requestedUser.id != loggedInUser.id &&
                              loggedInUser is Admin)
                          ? ElevatedButton(
                              // style: ButtonStyle(backgroundColor: Theme.of(context).),
                              onPressed: () {
                                Navigator.push(
                                    context,
                                    new MaterialPageRoute(
                                        builder: (context) =>
                                            new MerchantStoreForm(
                                              owner: widget.requestedUser,
                                            )));
                              },
                              child: Text("Add",
                                  style: TextStyle(
                                      fontSize: 16,
                                      fontWeight: FontWeight.bold)),
                            )
                          : Container(),
                    ],
                  ))
            ],
          );
        }
        return Container();
      },
    );
  }

  void showImageSource(BuildContext context) {
    final UserBloc _userBloc = BlocProvider.of<UserBloc>(context);
    User theUser = (_userBloc.state as Authenticated).user;
    Function(ImageSource) selectedImageSourceMode = (imageSource) => context
        .read<ProfileBLoc>()
        .add(OpenImagePicker(imageSource: imageSource, user: theUser));
    showModalBottomSheet(
        context: context,
        builder: (context) => Wrap(
              children: [
                ListTile(
                  leading: Icon(Icons.camera_alt),
                  title: Text("Camera"),
                  onTap: () {
                    Navigator.pop(context);
                    selectedImageSourceMode(ImageSource.camera);
                  },
                ),
                ListTile(
                  leading: Icon(Icons.photo_album),
                  title: Text("Gallery"),
                  onTap: () {
                    Navigator.pop(context);
                    selectedImageSourceMode(ImageSource.gallery);
                  },
                )
              ],
            ));
  }

  Widget _profilePage() {
    return SafeArea(
        child: Center(
      child: Column(
        children: [
          _nameTile(),
          _emailTile(),
          _phoneTile(),
        ],
      ),
    ));
  }

  Widget _nameTile() {
    return BlocBuilder<ProfileBLoc, ProfileState>(
        builder: ((context, state) => ListTile(
              leading: Icon(
                Icons.person,
                color: Theme.of(context).primaryColor,
              ),
              title: Row(
                children: [
                  Text(
                    state.user.firstname,
                    style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
                  ),
                  SizedBox(
                    width: 10,
                  ),
                  Text(
                    state.user.lastname,
                    style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
                  )
                ],
              ),
            )));
  }

  Widget _emailTile() {
    return BlocBuilder<ProfileBLoc, ProfileState>(
        builder: ((context, state) => ListTile(
              leading: Icon(Icons.email, color: Theme.of(context).primaryColor),
              title: Text(
                state.user.email,
                style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
            )));
  }

  Widget _phoneTile() {
    return BlocBuilder<ProfileBLoc, ProfileState>(
        builder: ((context, state) => ListTile(
              leading: Icon(Icons.phone, color: Theme.of(context).primaryColor),
              title: Text(
                state.user.phone,
                style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
            )));
  }

  // Builds Profile Image
  Widget buildImage(String imgSrc) {
    final backGroundImage = Image.network('${StaticDataStore.URI}${imgSrc}');

    return CircleAvatar(
      radius: 70,
      child: ClipOval(
        child: (imgSrc != '')
            ? Image.network(
                '${StaticDataStore.URI}${imgSrc}',
                fit: BoxFit.cover,
              )
            : Icon(
                Icons.person,
              ),
      ),
    );
  }

  // Builds Edit Icon on Profile Picture
  Widget buildEditIcon(BuildContext context) => buildCircle(
      all: 0,
      child: IconButton(
        icon: Icon(
          Icons.edit,
          size: 20,
        ),
        color: Colors.black,
        onPressed: () {
          print("Editing");
          context.read<ProfileBLoc>().add(ChangeAvatarRequest());
        },
      ));

  // Builds/Makes Circle for Edit Icon on Profile Picture
  Widget buildCircle({
    required Widget child,
    required double all,
  }) =>
      ClipOval(
          child: Container(
        padding: EdgeInsets.all(all),
        color: Theme.of(context).primaryColor,
        child: child,
      ));
}
