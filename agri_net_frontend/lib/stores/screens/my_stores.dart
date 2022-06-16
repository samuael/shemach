import '../../libs.dart';

class MyStoresScreen extends StatefulWidget {
  static String RouteName = "stores";
  User user;
  MyStoresScreen({required this.user});

  @override
  State<MyStoresScreen> createState() => _MyStoresScreenState();
}

class _MyStoresScreenState extends State<MyStoresScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        backgroundColor: Theme.of(context).primaryColor,
        elevation: 0,
        title: Text(
          "Stores",
          style: TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
        ),
      ),
      body: SafeArea(
        child: Padding(
          padding: EdgeInsets.all(10),
          child: BlocBuilder<StoreBloc, StoreState>(
            builder: (context, state) {
              if (state is MyStoresState) {
                return Flexible(
                    child: ListView.builder(
                        itemCount: state.myStores.length,
                        itemBuilder: (context, counter) {
                          return Card(
                            child: Padding(
                              padding: const EdgeInsets.all(10.0),
                              child: Column(
                                mainAxisAlignment:
                                    MainAxisAlignment.spaceEvenly,
                                children: [
                                  Row(
                                    children: [
                                      Text("Store : ",
                                          style: TextStyle(
                                              fontSize: 20,
                                              fontWeight: FontWeight.bold)),
                                      Text(
                                        state.myStores[counter].storeName,
                                        style: TextStyle(
                                          fontSize: 20,
                                        ),
                                      )
                                    ],
                                  ),
                                  Row(
                                    children: [
                                      Text("Owner : ",
                                          style: TextStyle(
                                              fontSize: 18,
                                              fontWeight: FontWeight.bold)),
                                      Text(
                                        widget.user.firstname +
                                            " " +
                                            widget.user.lastname,
                                        style: TextStyle(
                                          fontSize: 18,
                                        ),
                                      )
                                    ],
                                  ),
                                  Row(
                                    children: [
                                      Text("Active Products : ",
                                          style: TextStyle(
                                              fontSize: 18,
                                              fontWeight: FontWeight.bold)),
                                      Text(
                                        state.myStores[counter].activeProducts
                                            .toString(),
                                        style: TextStyle(
                                          fontSize: 18,
                                        ),
                                      )
                                    ],
                                  ),
                                  Row(
                                    children: [
                                      Text("Active Contracts : ",
                                          style: TextStyle(
                                              fontSize: 18,
                                              fontWeight: FontWeight.bold)),
                                      Text(
                                        state.myStores[counter].activeContracts
                                            .toString(),
                                        style: TextStyle(
                                          fontSize: 18,
                                        ),
                                      )
                                    ],
                                  ),
                                  ExpansionTile(
                                    textColor: Theme.of(context).primaryColor,
                                    iconColor: Theme.of(context).primaryColor,
                                    title: Text(
                                      "Address",
                                      style: TextStyle(
                                          color: Theme.of(context).primaryColor,
                                          fontSize: 18,
                                          fontWeight: FontWeight.bold),
                                    ),
                                    children: [
                                      Column(
                                        crossAxisAlignment:
                                            CrossAxisAlignment.start,
                                        children: [
                                          (state.myStores[counter].address
                                                      .UniqueAddressName !=
                                                  '')
                                              ? Row(
                                                  children: [
                                                    Text("Unique : ",
                                                        style: TextStyle(
                                                            fontSize: 16,
                                                            fontWeight:
                                                                FontWeight
                                                                    .bold)),
                                                    Text(
                                                        '${state.myStores[counter].address.UniqueAddressName}',
                                                        style: TextStyle(
                                                          fontSize: 16,
                                                        )),
                                                  ],
                                                )
                                              : SizedBox(),
                                          (state.myStores[counter].address
                                                      .Kebele !=
                                                  '')
                                              ? Row(
                                                  children: [
                                                    Text("Kebele : ",
                                                        style: TextStyle(
                                                            fontSize: 16,
                                                            fontWeight:
                                                                FontWeight
                                                                    .bold)),
                                                    Text(
                                                        '${state.myStores[counter].address.Kebele}',
                                                        style: TextStyle(
                                                          fontSize: 16,
                                                        )),
                                                  ],
                                                )
                                              : SizedBox(),
                                          (state.myStores[counter].address
                                                      .Woreda !=
                                                  '')
                                              ? Row(
                                                  children: [
                                                    Text("Woreda : ",
                                                        style: TextStyle(
                                                            fontSize: 16,
                                                            fontWeight:
                                                                FontWeight
                                                                    .bold)),
                                                    Text(
                                                        '${state.myStores[counter].address.Woreda}',
                                                        style: TextStyle(
                                                          fontSize: 16,
                                                        )),
                                                  ],
                                                )
                                              : SizedBox(),
                                          (state.myStores[counter].address
                                                      .City !=
                                                  '')
                                              ? Row(
                                                  children: [
                                                    Text("City : ",
                                                        style: TextStyle(
                                                            fontSize: 16,
                                                            fontWeight:
                                                                FontWeight
                                                                    .bold)),
                                                    Text(
                                                        '${state.myStores[counter].address.City}',
                                                        style: TextStyle(
                                                          fontSize: 16,
                                                        )),
                                                  ],
                                                )
                                              : SizedBox(),
                                          (state.myStores[counter].address
                                                      .Zone !=
                                                  '')
                                              ? Row(
                                                  children: [
                                                    Text("Zone : ",
                                                        style: TextStyle(
                                                            fontSize: 16,
                                                            fontWeight:
                                                                FontWeight
                                                                    .bold)),
                                                    Text(
                                                        '${state.myStores[counter].address.Zone}',
                                                        style: TextStyle(
                                                          fontSize: 16,
                                                        ))
                                                  ],
                                                )
                                              : SizedBox(),
                                          (state.myStores[counter].address
                                                      .Region !=
                                                  '')
                                              ? Row(
                                                  children: [
                                                    Text("Region : ",
                                                        style: TextStyle(
                                                            fontSize: 16,
                                                            fontWeight:
                                                                FontWeight
                                                                    .bold)),
                                                    Text(
                                                        '${state.myStores[counter].address.Region},',
                                                        style: TextStyle(
                                                          fontSize: 16,
                                                        ))
                                                  ],
                                                )
                                              : SizedBox(),
                                          Row(
                                            mainAxisAlignment:
                                                MainAxisAlignment.spaceBetween,
                                            children: [
                                              Text(
                                                  "Created At ${state.myStores[counter].createdAt}",
                                                  style: TextStyle(
                                                      fontSize: 14,
                                                      fontWeight:
                                                          FontWeight.bold)),
                                              IconButton(
                                                  iconSize: 40,
                                                  onPressed: () {
                                                    showDialog(
                                                      barrierDismissible: false,
                                                      context: context,
                                                      builder:
                                                          (BuildContext cxt) {
                                                        return Padding(
                                                          padding: EdgeInsets
                                                              .fromLTRB(
                                                                  0, 50, 0, 10),
                                                          child: Dialog(
                                                            shape: RoundedRectangleBorder(
                                                                borderRadius: BorderRadius.only(
                                                                    topLeft: Radius
                                                                        .circular(
                                                                            10),
                                                                    topRight: Radius
                                                                        .circular(
                                                                            10))),
                                                            child: Stack(
                                                              children: <
                                                                  Widget>[
                                                                storeLocation(
                                                                    state.myStores[
                                                                        counter]),
                                                                Positioned.fill(
                                                                  right: 0.0,
                                                                  child: Align(
                                                                    alignment:
                                                                        Alignment
                                                                            .topRight,
                                                                    child: InkWell(
                                                                        onTap: () {
                                                                          Navigator.of(context)
                                                                              .pop();
                                                                        },
                                                                        child: Icon(
                                                                          Icons
                                                                              .close_outlined,
                                                                          size:
                                                                              30,
                                                                          color:
                                                                              Colors.red,
                                                                        )),
                                                                  ),
                                                                ),
                                                              ],
                                                            ),
                                                          ),
                                                        );
                                                      },
                                                    );
                                                  },
                                                  color: Colors.redAccent,
                                                  icon: Icon(Icons
                                                      .location_on_rounded)),
                                            ],
                                          ),
                                        ],
                                      ),
                                    ],
                                  )
                                ],
                              ),
                            ),
                          );
                        }));
              }

              return Container();
            },
          ),
        ),
      ),
    );
  }

  Widget storeLocation(Store store) {
    final _controller = Completer<GoogleMapController>();
    MapPickerController mapPickerController = MapPickerController();
    LatLng target = LatLng(store.address.Latitude, store.address.Longitude);
    CameraPosition cameraPosition = CameraPosition(target: target, zoom: 2.0);

    final Set<Marker> markers = new Set();
    markers.add(Marker(
      //add second marker
      markerId: MarkerId(target.toString()),
      position: target, //position of marker
      infoWindow: InfoWindow(
        //popup info
        title: 'Store',
        snippet: '${store.storeName}',
      ),
      icon: BitmapDescriptor.defaultMarker, //Icon for Marker
    ));

    return MapPicker(
      // pass icon widget
      iconWidget: Icon(
        Icons.location_pin,
        color: Colors.red,
        size: 30,
      ),
      //add map picker controller
      mapPickerController: mapPickerController,
      child: GoogleMap(
        markers: markers,
        myLocationEnabled: true,
        zoomControlsEnabled: false,
        // hide location button
        myLocationButtonEnabled: false,
        mapType: MapType.normal,
        //  camera position
        initialCameraPosition: cameraPosition,
        onMapCreated: (GoogleMapController controller) {
          _controller.complete(controller);
        },
        onCameraMoveStarted: () {
          // notify map is moving
          mapPickerController.mapMoving!();
          // textController.text = "checking ...";
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
          // textController.text =
          //     '${placemarks.first.name}, ${placemarks.first.administrativeArea}, ${placemarks.first.country}';
        },
      ),
    );
  }
}
