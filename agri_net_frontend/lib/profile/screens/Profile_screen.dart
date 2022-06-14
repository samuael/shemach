import '../../libs.dart';

class UserProfileScreen extends StatefulWidget {
  static String RouteName = "profile";
  User requestedUser;
  UserProfileScreen({required this.requestedUser});

  @override
  State<UserProfileScreen> createState() => _UserProfileScreenState();
}

class _UserProfileScreenState extends State<UserProfileScreen> {
  final _controller = Completer<GoogleMapController>();

  MapPickerController mapPickerController = MapPickerController();

  final _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    final UserBloc _userBloc = BlocProvider.of<UserBloc>(context);
    User loggedInUser = (_userBloc.state as Authenticated).user;
    return BlocProvider(
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
                                  ? adress(
                                      (widget.requestedUser as Admin).address,
                                      context)
                                  : (widget.requestedUser is Merchant)
                                      ? Expanded(
                                          child: Column(children: [
                                            myStores(
                                                widget.requestedUser
                                                    as Merchant,
                                                context),
                                            adress(
                                                (widget.requestedUser
                                                        as Merchant)
                                                    .address,
                                                context),
                                          ]),
                                        )
                                      : (widget.requestedUser is Agent)
                                          ? adress(
                                              (widget.requestedUser as Agent)
                                                  .address,
                                              context)
                                          : Container()
                            ],
                          )),
                        ),
                      ),
                    ),
                  ),
                )),
          ),
        ));
  }

  Widget myStores(Merchant merchant, context) {
    final UserBloc _userBloc = BlocProvider.of<UserBloc>(context);
    User loggedInUser = (_userBloc.state as Authenticated).user;
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceAround,
      children: [
        Row(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: [
            Text(
              "Stores :",
              style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
            ),
            SizedBox(
              width: 5,
            ),
            Text(merchant.storeCount.toString(),
                style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold)),
          ],
        ),
        Container(
            width: MediaQuery.of(context).size.width / 4 - 19,
            child: Center(
                child: Row(
              children: [
                (widget.requestedUser.id != loggedInUser.id &&
                        loggedInUser is Admin)
                    ? ElevatedButton(
                        onPressed: () {
                          Navigator.push(
                              context,
                              new MaterialPageRoute(
                                  builder: (context) => new MerchantStoreForm(
                                        owner: widget.requestedUser,
                                      )));
                        },
                        child: Text("New",
                            style: TextStyle(
                                fontSize: 16, fontWeight: FontWeight.bold)),
                      )
                    : Container(),
              ],
            )))
      ],
    );
  }

  Widget adress(Address address, BuildContext context) {
    double lat = address.Latitude;
    double lon = address.Longitude;
    CameraPosition cameraPosition =
        CameraPosition(target: LatLng(lat, lon), zoom: 0.0);

    var textController = TextEditingController();
    return Flexible(
      child: Column(mainAxisAlignment: MainAxisAlignment.center, children: [
        Padding(
          padding: const EdgeInsets.all(15.0),
          child: Material(
            elevation: 15,
            child: Container(
              height: MediaQuery.of(context).size.height * 0.2,
              child: Column(
                mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                children: [
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      Container(
                        width: MediaQuery.of(context).size.width / 3 + 30,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            Text(
                              "City : ",
                              style: TextStyle(
                                  fontSize: 20, fontWeight: FontWeight.normal),
                            ),
                            Expanded(
                              child: new SingleChildScrollView(
                                scrollDirection: Axis.horizontal,
                                child: new Text(
                                  address.City,
                                  style: TextStyle(
                                      fontSize: 20,
                                      fontWeight: FontWeight.bold),
                                ),
                              ),
                            )
                          ],
                        ),
                      ),
                      Container(
                        width: MediaQuery.of(context).size.width / 3 + 30,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            Text("Woreda : ",
                                style: TextStyle(
                                    fontSize: 20,
                                    fontWeight: FontWeight.normal)),
                            Expanded(
                              child: new SingleChildScrollView(
                                scrollDirection: Axis.horizontal,
                                child: new Text(
                                  address.Woreda,
                                  style: TextStyle(
                                      fontSize: 20,
                                      fontWeight: FontWeight.bold),
                                ),
                              ),
                            )
                          ],
                        ),
                      ),
                    ],
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      Container(
                        width: MediaQuery.of(context).size.width / 3 + 30,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            Text(
                              "Kebele : ",
                              style: TextStyle(
                                  fontSize: 20, fontWeight: FontWeight.normal),
                            ),
                            Expanded(
                              child: new SingleChildScrollView(
                                scrollDirection: Axis.horizontal,
                                child: new Text(
                                  address.Kebele,
                                  style: new TextStyle(
                                      fontSize: 20,
                                      fontWeight: FontWeight.bold),
                                ),
                              ),
                            )
                          ],
                        ),
                      ),
                      Container(
                        width: MediaQuery.of(context).size.width / 3 + 30,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            Text(
                              "Unique : ",
                              style: TextStyle(
                                  fontSize: 20, fontWeight: FontWeight.normal),
                            ),
                            Expanded(
                              child: new SingleChildScrollView(
                                scrollDirection: Axis.horizontal,
                                child: new Text(
                                  address.UniqueAddressName,
                                  style: new TextStyle(
                                      fontSize: 20,
                                      fontWeight: FontWeight.bold),
                                ),
                              ),
                            )
                          ],
                        ),
                      ),
                    ],
                  ),
                  Row(
                    mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                    children: [
                      Container(
                        width: MediaQuery.of(context).size.width / 3 + 30,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            Text(
                              "Zone : ",
                              style: TextStyle(
                                  fontSize: 20, fontWeight: FontWeight.normal),
                            ),
                            Expanded(
                              child: new SingleChildScrollView(
                                scrollDirection: Axis.horizontal,
                                child: new Text(
                                  address.Zone,
                                  style: new TextStyle(
                                      fontSize: 16,
                                      fontWeight: FontWeight.bold),
                                ),
                              ),
                            )
                          ],
                        ),
                      ),
                      Container(
                        width: MediaQuery.of(context).size.width / 3 + 30,
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.start,
                          children: [
                            Text(
                              "Region : ",
                              style: TextStyle(
                                  fontSize: 20, fontWeight: FontWeight.normal),
                            ),
                            Expanded(
                              child: new SingleChildScrollView(
                                scrollDirection: Axis.horizontal,
                                child: new Text(
                                  address.Region,
                                  style: new TextStyle(
                                      fontSize: 20,
                                      fontWeight: FontWeight.bold),
                                ),
                              ),
                            )
                          ],
                        ),
                      )
                    ],
                  ),
                ],
              ),
            ),
          ),
        ),
        Expanded(
          child: Container(
            height: MediaQuery.of(context).size.height * 0.4,
            child: MapPicker(
              // pass icon widget
              iconWidget: Icon(
                Icons.location_pin,
                color: Colors.red,
                size: 30,
              ),
              //add map picker controller
              mapPickerController: mapPickerController,
              child: GoogleMap(
                myLocationEnabled: true,
                zoomControlsEnabled: false,
                // hide location button
                myLocationButtonEnabled: false,
                mapType: MapType.hybrid,
                //  camera position
                initialCameraPosition: cameraPosition,
                onMapCreated: (GoogleMapController controller) {
                  _controller.complete(controller);
                },
                onCameraMoveStarted: () {
                  // notify map is moving
                  mapPickerController.mapMoving!();
                  textController.text = "checking ...";
                },
                // onCameraMove: (cameraPosition) {
                //   cameraPosition
                // },
                onCameraIdle: () async {
                  // notify map stopped moving
                  mapPickerController.mapFinishedMoving!();
                  //get address name from camera position
                  List<Placemark> placemarks = await placemarkFromCoordinates(
                    cameraPosition.target.latitude,
                    cameraPosition.target.longitude,
                  );

                  // update the ui with the address
                  textController.text =
                      '${placemarks.first.name}, ${placemarks.first.administrativeArea}, ${placemarks.first.country}';
                },
              ),
            ),
          ),
        ),
      ]),
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
              leading: Icon(Icons.person),
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
              leading: Icon(Icons.email),
              title: Text(
                state.user.email,
                style: TextStyle(fontSize: 16, fontWeight: FontWeight.bold),
              ),
            )));
  }

  Widget _phoneTile() {
    return BlocBuilder<ProfileBLoc, ProfileState>(
        builder: ((context, state) => ListTile(
              leading: Icon(Icons.phone),
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
        color: Color.fromARGB(255, 175, 168, 168),
        child: child,
      ));
}
