import 'package:agri_net_frontend/profile/widgets/address.dart';
import 'package:agri_net_frontend/profile/widgets/widgets.dart';

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
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    final userProvider = BlocProvider.of<UserBloc>(context);
    final loggedInUser = (userProvider.state as Authenticated).user;
    final storeProvider = BlocProvider.of<StoreBloc>(context);
    if (widget.requestedUser is Merchant) {
      storeProvider.add(LoadMyStoresEvent(ownerId: widget.requestedUser.id));
    }

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
          " Profile ",
          style: TextStyle(
              fontSize: 18, fontWeight: FontWeight.bold, color: Colors.black),
        ),
        centerTitle: true,
      ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.fromLTRB(10, 0, 10, 10),
          child: ClipRRect(
            borderRadius: BorderRadius.circular(10),
            child: Container(
              child: BlocProvider(
                  create: (Context) {
                    return ProfileBLoc(
                        user: widget.requestedUser,
                        isCurrentUser:
                            widget.requestedUser.id == loggedInUser.id,
                        profileRepository: ProfileRepository(
                            profileProvider: ProfileProvider()));
                  },
                  child: BlocListener<ProfileBLoc, ProfileState>(
                    listener: ((context, state) {
                      if (state.imageSourceActionSheetIsVisible!) {
                        showImageSource(context);
                      }
                      if (state.avatorPath != '') {
                        setState(() {
                          loggedInUser.imgurl = state.avatorPath!;
                        });
                      }
                    }),
                    child: BlocBuilder<ProfileBLoc, ProfileState>(
                      builder: ((context, state) => Scaffold(
                            body: Material(
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
                                child: Center(
                                    child: Column(
                                  children: [
                                    Stack(children: [
                                      buildImage(widget.requestedUser.imgurl),
                                      (widget.requestedUser.id ==
                                              loggedInUser.id)
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
                                    Column(
                                      children: [
                                        (widget.requestedUser.firstname != '')
                                            ? _nameTile()
                                            : SizedBox(),
                                        (widget.requestedUser.email != '')
                                            ? _emailTile()
                                            : SizedBox(),
                                        (widget.requestedUser.phone != '')
                                            ? _phoneTile()
                                            : SizedBox(),
                                      ],
                                    ),
                                    (widget.requestedUser is Admin)
                                        ? AddressView(
                                            (widget.requestedUser as Admin)
                                                .address)
                                        : (widget.requestedUser is Merchant)
                                            ? Expanded(
                                                child: Column(children: [
                                                  AddressView(
                                                      (widget.requestedUser
                                                              as Merchant)
                                                          .address),
                                                  ExpansionTile(
                                                    textColor: Theme.of(context)
                                                        .primaryColor,
                                                    iconColor: Theme.of(context)
                                                        .primaryColor,
                                                    title: Text(
                                                      "Stores",
                                                      style: TextStyle(
                                                          color:
                                                              Theme.of(context)
                                                                  .primaryColor,
                                                          fontSize: 18,
                                                          fontWeight:
                                                              FontWeight.bold),
                                                    ),
                                                    children: [
                                                      myStores(context)
                                                    ],
                                                  ),
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
                          )),
                    ),
                  )),
            ),
          ),
        ),
      ),
    );
  }

  Widget myStores(BuildContext context) {
    final storesState = context.watch<StoreBloc>().state;
    if (storesState is LoadingMyStoresState) {
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
    } else if (storesState is MyStoresLoadedState &&
        storesState.myStores[widget.requestedUser.id] != null) {
      final myStores = storesState.myStores[widget.requestedUser.id];
      return Container(
        height: MediaQuery.of(context).size.height * 0.5 - 65,
        child: Expanded(
          child: SingleChildScrollView(
            child: ListView.builder(
                shrinkWrap: true,
                physics: ClampingScrollPhysics(),
                itemCount: myStores!.length,
                itemBuilder: (context, counter) {
                  return StoreView(myStores[counter]);
                }),
          ),
        ),
      );
      // Expanded(
      //     child: Column(
      //   children: myStores!.map<StoreView>((e) => StoreView(e)).toList(),
      // ));
    } else if (storesState is LoadingStoresFailedState) {
      return Center(
        child: Column(
          children: [
            Text("${storesState.msg}"),
            IconButton(
              icon: Icon(
                Icons.replay,
                color: Colors.blue,
              ),
              onPressed: () {
                context
                    .read<StoreBloc>()
                    .add(LoadMyStoresEvent(ownerId: widget.requestedUser.id));
              },
            )
          ],
        ),
      );
    } else {
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
                context
                    .read<StoreBloc>()
                    .add(LoadMyStoresEvent(ownerId: widget.requestedUser.id));
              },
            )
          ],
        ),
      );
    }
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
            : Image.asset('assets/images/Avatar_icon_green.svg.png'),
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
