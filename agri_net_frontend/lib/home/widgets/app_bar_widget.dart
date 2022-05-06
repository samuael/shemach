import '../../libs.dart';
import '../../theme.dart';

class UserScreenAppBarDrawer extends StatefulWidget
    implements PreferredSizeWidget {
  const UserScreenAppBarDrawer({Key? key}) : super(key: key);

  @override
  State<UserScreenAppBarDrawer> createState() => _UserScreenAppBarDrawerState();

  @override
  Size get preferredSize => Size.fromHeight(60);
}

class _UserScreenAppBarDrawerState extends State<UserScreenAppBarDrawer> {
  @override
  Widget build(BuildContext context) {
    var we = MediaQuery.of(context).size.width;
    var he = MediaQuery.of(context).size.height;
    Widget divider;
    return Material(
      elevation: 5,
      child: Container(
        // color: appBarTheme,
        width: we,
        height: he / 11.5,
        child: Padding(
          padding: const EdgeInsets.fromLTRB(20, 0, 30, 0),
          child: Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              AgriNetLogo(),
              Row(
                mainAxisAlignment: MainAxisAlignment.end,
                children: [
                  Row(
                    children: [
                      Container(child: LanguageDropDown()),
                      SizedBox(
                        width: 30,
                      ),
                      UserAccountePage(),
                    ],
                  ),
                ],
              )
            ],
          ),
        ),
      ),
    );
  }

  Widget LanguageDropDown() {
    List lang = ["Amh", "Eng"];
    return DropdownButton(
        hint: Text("lang"),
        items: lang
            .map((item) => DropdownMenuItem(value: item, child: new Text(item)))
            .toList(),
        onChanged: (value) {
          setState(() {
            value = value;
          });
        });
  }
}
