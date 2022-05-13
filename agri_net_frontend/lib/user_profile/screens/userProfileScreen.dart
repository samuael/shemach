import '../../libs.dart';

// void main() {
//   runApp(ProfileScreen());

//   // Handles Status and Nav bar styling/theme
//   // SystemChrome.setSystemUIOverlayStyle(SystemUiOverlayStyle(
//   //   systemNavigationBarColor: Colors.transparent,
//   //   statusBarColor: Colors.transparent,
//   //   statusBarIconBrightness: Brightness.dark,
//   // ));
// }

class ProfileScreen extends StatefulWidget {
  static const String RouteName = "profile";
  @override
  _ProfileScreenState createState() => _ProfileScreenState();
}

class _ProfileScreenState extends State<ProfileScreen> {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'User Profile',
      debugShowCheckedModeBanner: false,
      theme: ThemeData(
          primaryColor: Colors.black,
          fontFamily: 'Roboto',
          elevatedButtonTheme: ElevatedButtonThemeData(
              style: ElevatedButton.styleFrom(
                  primary: Colors.black,
                  shadowColor: Colors.grey,
                  elevation: 20,
                  shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.all(Radius.circular(0.0))))),
          inputDecorationTheme: InputDecorationTheme(
              border:
                  OutlineInputBorder(borderRadius: BorderRadius.circular(0.0))),
          textButtonTheme: TextButtonThemeData(
            style: TextButton.styleFrom(
              alignment: Alignment.centerLeft,
              primary: Colors.black,
            ),
          )),
      home: ProfilePage(),
    );
  }
}
