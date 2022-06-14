import "../../libs.dart";
import "dart:convert";

class ProductPostSingleImageItem extends StatefulWidget {
  int imageID;
  ProductPostSingleImageItem(this.imageID, {Key? key}) : super(key: key);

  @override
  _ProductPostImageItemState createState() => _ProductPostImageItemState();
}

class _ProductPostImageItemState extends State<ProductPostSingleImageItem> {
  bool loadFullImage = false;

  @override
  Widget build(BuildContext context) {
    print(
      "${StaticDataStore.SCHEME}://${StaticDataStore.HOST}:${StaticDataStore.PORT}/post/image/${widget.imageID}/blurred/",
    );
    print(jsonEncode(
        {"authorization": StaticDataStore.HEADERS["authorization"]!}));
    return ClipRRect(
      borderRadius: BorderRadius.circular(5),
      child: Container(
        decoration: BoxDecoration(
          borderRadius: BorderRadius.circular(10),
          border: Border.all(color: Theme.of(context).primaryColorLight),
        ),
        child: Image.network(
          "${StaticDataStore.SCHEME}://${StaticDataStore.HOST}:${StaticDataStore.PORT}/post/image/${widget.imageID}/blurred/",
          headers: {"authorization": StaticDataStore.HEADERS["authorization"]!},
          errorBuilder: (context, _, er) {
            return Stack(
              children: [
                AnimatedOpacity(
                  opacity: 0.4,
                  duration: Duration(
                    seconds: 2,
                  ),
                  child: Container(
                    color: Color.fromARGB(40, 49, 48, 48),
                    width: double.infinity,
                    height: MediaQuery.of(context).size.width * 0.5,
                    child: Image.asset(
                      'assets/images/logo.jpg',
                      // color: Color.fromARGB(5, 0, 0, 0),
                    ),
                  ),
                ),
                Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  crossAxisAlignment: CrossAxisAlignment.center,
                  children: [
                    Positioned(
                      top: MediaQuery.of(context).size.width * 0.25,
                      child: Container(
                        // color: Color.fromARGB(40, 26, 25, 25),
                        child: Icon(
                          widget.imageID ==0 ? Icons.close :Icons.refresh,
                          color: widget.imageID ==0 ?  Colors.red:Theme.of(context).primaryColorDark,
                          size: 40,
                        ),
                      ),
                    ),
                  ],
                )
              ],
            );
          },
        ),
      ),
    );
  }
}
