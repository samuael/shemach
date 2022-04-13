import 'package:flutter/services.dart';

import "../../libs.dart";

class LoginWidget extends StatefulWidget {
  Function forgotFunction;
  BuildContext screenContext;
  LoginWidget(this.forgotFunction, this.screenContext, {Key? key})
      : super(key: key);

  @override
  _LoginWidgetState createState() => _LoginWidgetState();
}

class _LoginWidgetState extends State<LoginWidget> {
  TextEditingController emailController = TextEditingController();
  TextEditingController passwordController = TextEditingController();
  bool logging = false;

  @override
  Widget build(BuildContext context) {
    return Column(
      crossAxisAlignment: CrossAxisAlignment.center,
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        Container(
          padding: EdgeInsets.all(5),
          child: Text(
            " Login ",
            style: TextStyle(
              color: Theme.of(context).primaryColor,
              fontWeight: FontWeight.bold,
              fontFamily: "Moms TypeWriter",
              fontSize: 22,
              // fontSize: 18
            ),
          ),
        ),
        Container(
          padding: EdgeInsets.symmetric(
            horizontal: 20,
            vertical: 10,
          ),
          child: Text(
            "  ",
            style: TextStyle(
              color: Colors.green,
              fontWeight: FontWeight.bold,
              fontFamily: "Elegant TypeWriter",
              // fontSize: 18
            ),
          ),
        ),
        Container(
          padding: EdgeInsets.symmetric(
            horizontal: 20,
          ),
          child: TextField(
            inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
            keyboardType: TextInputType.number,
            cursorColor: Theme.of(context).primaryColorLight,
            controller: emailController,
            style: TextStyle(
              fontWeight: FontWeight.bold,
            ),
            decoration: InputDecoration(
              labelText: "Phone",
              fillColor: Colors.lightBlue,
              hoverColor: Colors.lightBlue,
              prefix: Container(
                padding: EdgeInsets.symmetric(horizontal: 5),
                child: Text(
                  "+251",
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                  ),
                ),
              ),
              suffixIcon: Icon(Icons.phone),
              border: OutlineInputBorder(
                borderSide: BorderSide(
                  color: Colors.lightBlue,
                  style: BorderStyle.none,
                ),
              ),
            ),
            onChanged: (text) {
              if (text.length > 9) {
                emailController.text = text.substring(1, 9);
              }
            },
          ),
        ),
        // Container(
        //   padding: EdgeInsets.symmetric(
        //     horizontal: 20,
        //     vertical: 10,
        //   ),
        //   child: TextField(
        //     cursorColor: Theme.of(context).primaryColorLight,
        //     obscureText: true,
        //     controller: passwordController,
        //     decoration: InputDecoration(
        //         labelText: "Password",
        //         border: OutlineInputBorder(),
        //         suffixIcon: Icon(
        //           Icons.remove_red_eye,
        //         )),
        //   ),
        // ),
        Container(
            padding: EdgeInsets.symmetric(
              horizontal: 30,
              vertical: 10,
            ),
            child: ElevatedButton.icon(
              style: ButtonStyle(
                padding: MaterialStateProperty.all<EdgeInsets>(
                  EdgeInsets.symmetric(
                    vertical: 10,
                    horizontal: 40,
                  ),
                ),
              ),
              onPressed: () async {
                // checking the validity of input values
              },
              icon: Icon(Icons.login),
              label: Text(
                " Login ",
                style: TextStyle(
                  fontWeight: FontWeight.bold,
                ),
              ),
            )),
        GestureDetector(
          onTap: () {
            Navigator.of(context).pushNamed(RegistrationScreen.RouteName);
          },
          child: Container(
            padding: EdgeInsets.all(10),
            child: Text(
              "Registration",
              style: TextStyle(
                fontWeight: FontWeight.bold,
                color: Theme.of(context).primaryColor,
              ),
            ),
          ),
        )
      ],
    );
  }
}
