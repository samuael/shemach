import "../../libs.dart";

Future<List<File>> SelectFilesPopup(BuildContext context) async {
  List<File> files = [];
  showDialog(
    context: context,
    builder: (contex) {
      return AlertDialog(
        title: Text(
          translate(lang, "Select images"),
          style: TextStyle(
            fontWeight: FontWeight.bold,
            fontStyle: FontStyle.italic,
          ),
        ),
        content: Container(
            height: 100,
            width: 100,
            color: Theme.of(context).primaryColor,
            child: Column(children: [
              Text(
                translate(
                  lang,
                  "Select images of the product",
                ),
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                  fontStyle: FontStyle.italic,
                ),
              ),
            ])),
      );
    },
  );
  return files;
}
