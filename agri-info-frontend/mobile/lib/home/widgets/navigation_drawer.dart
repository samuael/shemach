import "../../libs.dart";
// import 'package:flutter_svg/svg.dart';

class NavigationDrawer extends StatelessWidget {
  NavigationDrawer({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Drawer(
      child: Container(
        height: MediaQuery.of(context).size.height * 0.9,
        color: Color(0xff00000),
        child: SingleChildScrollView(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: [
              Stack(
                children: [
                  Container(
                    height: 200,
                    color: Theme.of(context).primaryColor,
                    child: Column(
                      children: [
                        Container(
                          color: Colors.white,
                          child: ClipRRect(
                            borderRadius: BorderRadius.only(
                              bottomRight: Radius.circular(100),
                            ),
                            child: Container(
                              color: Theme.of(context).primaryColor,
                              height: 100,
                            ),
                          ),
                        ),
                        ClipRRect(
                          borderRadius: BorderRadius.only(
                            topLeft: Radius.circular(100),
                          ),
                          child: Container(
                            height: 100,
                            color: Colors.white,
                          ),
                        ),
                      ],
                    ),
                  ),
                  Positioned(
                    top: 40,
                    left: 100,
                    child: ClipRRect(
                      borderRadius: BorderRadius.circular(40),
                      child: Container(
                        decoration: BoxDecoration(
                          borderRadius: BorderRadius.circular(40),
                          border: Border.all(
                            color: Colors.black,
                          ),
                          color: Colors.white,
                        ),
                        child: Image.asset(
                          "assets/image_assets/agri_net_final_temporary_logo.png",
                          width: 80,
                          height: 80,
                        ),
                      ),
                    ),
                  ),
                  // backgroundImage:
                  //     AssetImage('assets/image_assets/user1.jpg')),
                  Container(
                      margin: EdgeInsets.only(top: 135, left: 90),
                      width: MediaQuery.of(context).size.width,
                      child: Text(
                        "nathyalem@aait.com",
                        style: TextStyle(
                            fontFamily: "Roboto", fontStyle: FontStyle.italic),
                      )),
                ],
              ),
              Container(
                color: Colors.white,
                // height: 530,
                height: MediaQuery.of(context).size.height * 0.65,
                child: Container(
                  margin: EdgeInsets.symmetric(vertical: 40),
                  child: Column(
                    crossAxisAlignment: CrossAxisAlignment.center,
                    children: [
                      GestureDetector(
                        onTap: () {
                          // Navigator.of(context).pushNamed(SubScreen1.RouteName);
                        },
                        child: Container(
                            padding: EdgeInsets.symmetric(
                                horizontal: 15, vertical: 20),
                            child: Row(
                              children: [
                                Icon(Icons.subscriptions,
                                    color: Theme.of(context).primaryColor),
                                Container(
                                  margin: EdgeInsets.symmetric(
                                    horizontal: 10,
                                  ),
                                  padding: EdgeInsets.symmetric(horizontal: 10),
                                  child: Text(
                                    "Subscription",
                                    style: TextStyle(
                                        fontSize: 18,
                                        color: Theme.of(context).primaryColor,
                                        fontFamily: "Roboto",
                                        fontWeight: FontWeight.bold),
                                  ),
                                )
                              ],
                            )),
                      ),
                      Container(
                        padding:
                            EdgeInsets.symmetric(horizontal: 15, vertical: 20),
                        child: Row(
                          children: [
                            // SvgPicture.asset(
                            //   "assets/icons/settings_black_24dp (1).svg",
                            //   color: Colors.white,
                            //   width: 20,
                            //   height: 20,
                            // ),
                            Icon(Icons.settings,
                                color: Theme.of(context).primaryColor),
                            Container(
                              margin: EdgeInsets.symmetric(horizontal: 10),
                              padding: EdgeInsets.symmetric(horizontal: 10),
                              child: Text(
                                "Settings",
                                style: TextStyle(
                                    fontSize: 18,
                                    color: Theme.of(context).primaryColor,
                                    fontFamily: "Roboto",
                                    fontWeight: FontWeight.bold),
                              ),
                            )
                          ],
                        ),
                      ),
                    ],
                  ),
                ),
              ),
              Container(
                height: 60,
                margin: EdgeInsets.symmetric(horizontal: 7),
                padding: EdgeInsets.symmetric(horizontal: 10, vertical: 20),
                child: Row(
                  children: [
                    // SvgPicture.asset(
                    //   'assets/icons/logout_black_24dp (1).svg',
                    //   color: Colors.white,
                    //   width: 20,
                    //   height: 20,
                    // ),
                    Icon(Icons.logout, color: Theme.of(context).primaryColor),
                    Container(
                      margin: EdgeInsets.symmetric(horizontal: 10),
                      padding: EdgeInsets.symmetric(horizontal: 10),
                      child: Text(
                        "Logout",
                        style: TextStyle(
                            fontSize: 18,
                            color: Theme.of(context).primaryColor,
                            fontFamily: "Roboto",
                            fontWeight: FontWeight.bold), //"Times New Roman"),
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
