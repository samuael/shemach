import '../../libs.dart';

class MerchantStoreForm extends StatefulWidget {
  const MerchantStoreForm({Key? key}) : super(key: key);

  @override
  State<MerchantStoreForm> createState() => _MerchantStoreFormState();
}

class _MerchantStoreFormState extends State<MerchantStoreForm> {
  TextEditingController storeNameController = TextEditingController();

  TextEditingController kebeleController = TextEditingController();
  TextEditingController woredaController = TextEditingController();
  TextEditingController cityControler = TextEditingController();
  TextEditingController zoneController = TextEditingController();
  TextEditingController regionController = TextEditingController();
  TextEditingController uniqueAddressController = TextEditingController();
  TextEditingController latitudeControler = TextEditingController();
  TextEditingController longitudeController = TextEditingController();

  final _controller = Completer<GoogleMapController>();
  MapPickerController mapPickerController = MapPickerController();

  final _formKey = GlobalKey<FormState>();

  CameraPosition cameraPosition = const CameraPosition(
      target: LatLng(30.653690770268437, 30.653690770268437), zoom: 0.0);

  var textController = TextEditingController();

  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Container(
          height: MediaQuery.of(context).size.height / 20,
          width: MediaQuery.of(context).size.width / 2 - 20,
          child: TextFormField(
            validator: (value) {
              if (value == null || value.trim().isEmpty) {
                return 'Please enter store name';
              }
              return null;
            },
            cursorColor: Theme.of(context).primaryColorLight,
            controller: storeNameController,
            decoration: InputDecoration(
              labelText: "Region",
              fillColor: Colors.lightBlue,
              hoverColor: Colors.lightBlue,
              border: OutlineInputBorder(
                borderSide: BorderSide(
                  color: Colors.lightBlue,
                  style: BorderStyle.none,
                ),
              ),
            ),
          ),
        ),
        Flexible(
          child: Container(
            height: MediaQuery.of(context).size.height * 0.4,
            child: Stack(
              children: [
                Expanded(
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
                      onCameraMove: (cameraPosition) {
                        this.cameraPosition = cameraPosition;
                      },
                      onCameraIdle: () async {
                        // notify map stopped moving
                        mapPickerController.mapFinishedMoving!();
                        //get address name from camera position
                        List<Placemark> placemarks =
                            await placemarkFromCoordinates(
                          cameraPosition.target.latitude,
                          cameraPosition.target.longitude,
                        );

                        // update the ui with the address
                        uniqueAddressController.text =
                            '${placemarks.first.name}';
                        regionController.text =
                            '${placemarks.first.administrativeArea}';
                        zoneController.text =
                            '${placemarks.first.subAdministrativeArea}';
                        cityControler.text = '${placemarks.first.locality}';
                        woredaController.text =
                            '${placemarks.first.subLocality}';
                        kebeleController.text = '${placemarks.first.street}';
                        textController.text =
                            '${placemarks.first.name}, ${placemarks.first.administrativeArea}, ${placemarks.first.country}';
                      },
                    ),
                  ),
                ),
                Positioned(
                  top: MediaQuery.of(context).size.height * 0.01,
                  width: MediaQuery.of(context).size.width - 20,
                  height: 50,
                  child: TextFormField(
                    maxLines: 3,
                    textAlign: TextAlign.center,
                    readOnly: true,
                    decoration: const InputDecoration(
                        contentPadding: EdgeInsets.zero,
                        border: InputBorder.none),
                    controller: textController,
                  ),
                ),
                Positioned(
                  top: MediaQuery.of(context).size.height * 0.3 + 35,
                  left: MediaQuery.of(context).size.width * 0.305,
                  child: Row(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      Padding(
                        padding: const EdgeInsets.all(2.0),
                        child: TextButton(
                          child: const Text(
                            "Track Location",
                            style: TextStyle(
                              fontWeight: FontWeight.bold,
                              fontStyle: FontStyle.normal,
                              color: Colors.black,
                              fontSize: 19,
                              // height: 19/19,
                            ),
                          ),
                          onPressed: () {
                            latitudeControler.text =
                                '${cameraPosition.target.latitude}';
                            longitudeController.text =
                                '${cameraPosition.target.longitude}';
                            print(
                                "Location ${cameraPosition.target.latitude} ${cameraPosition.target.longitude}");
                            print("Address: ${textController.text}");
                          },
                          style: ButtonStyle(
                            backgroundColor: MaterialStateProperty.all<Color>(
                                Color.fromARGB(255, 228, 223, 223)),
                            shape: MaterialStateProperty.all<
                                RoundedRectangleBorder>(
                              RoundedRectangleBorder(
                                borderRadius: BorderRadius.circular(10.0),
                              ),
                            ),
                          ),
                        ),
                      ),
                    ],
                  ),
                )
              ],
            ),
          ),
        ),
      ],
    );
  }

  bool _trySubmitForm() {
    final bool? isValid = _formKey.currentState?.validate();
    return isValid!;
  }
}
