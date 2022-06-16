import '../../libs.dart';

class StoreAddress extends StatefulWidget {
  double lat;
  double lon;
  StoreAddress({required this.lat, required this.lon});

  @override
  State<StoreAddress> createState() => _StoreAddressState();
}

class _StoreAddressState extends State<StoreAddress> {
  final _controller = Completer<GoogleMapController>();
  MapPickerController mapPickerController = MapPickerController();

  @override
  Widget build(BuildContext context) {
    LatLng target = LatLng(widget.lat, widget.lon);
    CameraPosition cameraPosition = CameraPosition(target: target, zoom: 2.0);
    final Set<Marker> markers = new Set();
    markers.add(Marker(
      //add second marker
      markerId: MarkerId(target.toString()),
      position: target, //position of marker
      infoWindow: InfoWindow(
        //popup info
        title: 'Store',
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
          mapPickerController.mapMoving!();
        },
        onCameraIdle: () async {
          // notify map stopped moving
          mapPickerController.mapFinishedMoving!();

          List<Placemark> placemarks = await placemarkFromCoordinates(
            cameraPosition.target.latitude,
            cameraPosition.target.longitude,
          );
        },
      ),
    );
  }
}
