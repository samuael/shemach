import '../../libs.dart';
import '../../theme.dart';

class CollapsingSideBarDrawer extends StatefulWidget {
  const CollapsingSideBarDrawer({Key? key}) : super(key: key);

  @override
  State<CollapsingSideBarDrawer> createState() =>
      _CollapsingSideBarDrawerState();
}

class NavigationModel {
  String title;
  IconData icon;

  NavigationModel(this.title, this.icon) {}
}

List<NavigationModel> sideBarItems = [
  NavigationModel("Products", Icons.home),
  NavigationModel("My Stores", Icons.store),
  NavigationModel("Contracts", Icons.person),
  NavigationModel("Search", Icons.search),
  NavigationModel("Notifications", Icons.notifications),
  NavigationModel("Sttings", Icons.settings),
  NavigationModel("Log Out", Icons.logout),
];

class _CollapsingSideBarDrawerState extends State<CollapsingSideBarDrawer>
    with SingleTickerProviderStateMixin {
  double maxWidth = 140;
  double minWidth = 40;
  bool isSideBarExpanded = false;
  late AnimationController animationController;
  late Animation<double> widthAnimation;
  int currentSelectedIndex = -1;

  @override
  void initState() {
    super.initState();
    animationController =
        AnimationController(vsync: this, duration: Duration(milliseconds: 200));
    widthAnimation =
        Tween(begin: maxWidth, end: minWidth).animate(animationController);
  }

  @override
  Widget build(BuildContext context) {
    return Material(
      elevation: 5,
      child: AnimatedBuilder(
          animation: animationController,
          builder: (context, widget) {
            return getWidget(context, widget);
          }),
    );
  }

  Widget getWidget(context, widget) {
    var we = MediaQuery.of(context).size.width;
    var he = MediaQuery.of(context).size.height;
    Widget divider;
    return Container(
      width: widthAnimation.value,
      color: sideBarBackgroundColor,
      child: Column(
        children: [
          Container(
            child: InkWell(
              onTap: () {
                setState(() {
                  isSideBarExpanded = !isSideBarExpanded;
                  isSideBarExpanded
                      ? animationController.forward()
                      : animationController.reverse();
                });
              },
              child: Padding(
                padding: const EdgeInsets.fromLTRB(0, 0, 10, 0),
                child: Row(
                  mainAxisAlignment: MainAxisAlignment.end,
                  children: [
                    AnimatedIcon(
                      icon: AnimatedIcons.close_menu,
                      progress: animationController,
                      color: Colors.white,
                      size: 25,
                    ),
                  ],
                ),
              ),
            ),
          ),
          Expanded(
              child: ListView.builder(
            itemBuilder: (context, counter) {
              if (counter == sideBarItems.length - 3) {
                divider = Divider(
                  height: 150,
                );
              } else {
                divider = Container();
              }
              return Column(
                children: [
                  CollapsingNavTile(
                      sideBarItems[counter].title,
                      sideBarItems[counter].icon,
                      animationController,
                      currentSelectedIndex == counter, () {
                    setState(() {
                      currentSelectedIndex = counter;
                    });
                  }),
                  divider
                ],
              );
            },
            itemCount: sideBarItems.length,
          )),
        ],
      ),
    );
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(DiagnosticsProperty<AnimationController>(
        'animationController', animationController));
  }
}
