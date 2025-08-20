import 'package:flutter/services.dart';
import "package:flutter_bloc/flutter_bloc.dart";
import 'package:mobile/auth/auth.dart';
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
  TextEditingController phoneController = TextEditingController();

  bool logging = false;
  String message = "";
  Color messageColor = Colors.white;

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
            // inputFormatters: [WhitelistingTextInputFormatter.digitsOnly],
            keyboardType: TextInputType.number,
            cursorColor: Theme.of(context).primaryColorLight,
            controller: phoneController,
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
                phoneController.text = text.substring(0, 9);
              }
            },
          ),
        ),
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
              if (phoneController.text.length == 9) {
                setState(() {
                  logging = true;
                });
                final response = await context
                    .read<AuthBloc>()
                    .loginSubscriber("+251${phoneController.text}");
                if (response.statusCode == 200 || response.statusCode == 201) {
                  setState(() {
                    logging = false;
                  });
                  Navigator.of(context)
                      .pushNamed(ConfirmationScreen.RouteName, arguments: {
                    "fullname": "",
                    "phone": phoneController.text,
                    "islogin": true,
                  });
                } else {
                  setState(() {
                    this.message = response.msg;
                    this.messageColor = Colors.red;
                    logging = false;
                  });
                }
              }else {
                setState(() {
                    this.message = translate(lang, "please provide valid phone number");
                    this.messageColor = Colors.red;
                    logging = false;
                  });
              }
            },
            icon: Icon(Icons.login),
            label: Text(
              translate(lang, " Login "),
              style: TextStyle(
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
        ),
        Container(
          padding: EdgeInsets.all(5),
          decoration: BoxDecoration(
            border: Border.all(
              color: Colors.black26,
            ),
            borderRadius: BorderRadius.circular(10),
          ),
          child: logging
              ? CircularProgressIndicator(
                  color: Theme.of(context).primaryColor,
                  strokeWidth: 5,
                )
              : Text(
                  translate(lang, message),
                  style: TextStyle(
                    fontStyle: FontStyle.italic,
                    color: messageColor,
                    fontWeight: FontWeight.bold,
                  ),
                ),
        ),
        GestureDetector(
          onTap: () {
            Navigator.of(context).pushNamed(RegistrationScreen.RouteName);
          },
          child: Container(
            padding: EdgeInsets.all(10),
            child: Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                Text(
                  translate( lang , "Registration"),
                  style: TextStyle(
                    fontWeight: FontWeight.bold,
                    color: Theme.of(context).primaryColor,
                    
                  ),
                  
                ),
                Icon(
                  Icons.arrow_right,
                  color: Theme.of(context).primaryColor,
                ),
              ],
            ),
          ),
        )
      ],
    );
  }
}
