import React, { Component } from "react";
import UserService from "../../services/prodcutService";
import { Link} from "react-router-dom";
import { withRouter } from "react-router";
import './loginpage.css'
import './headerInfo.css'    //For Header



class LoginPageImp extends Component {
  constructor(props) {
    super(props);
    this.onChangePassword = this.onChangePassword.bind(this);
    this.onChangeEmail = this.onChangeEmail.bind(this);
    this.onHandleSubmit = this.onHandleSubmit.bind(this);
    this.newLogin = this.newLogin.bind(this);

    

    // this.state = {
    //   id: null,
    //   firstname: "",
    //   lastname: "", 
    //   email: "",
    //   measurement: "",
    //   prevprice: 0,
    //   currentprice:0,
    //   published: false,

    //   submitted: false
    // };

    this.state = {
   //   response_data: {
        msg: "",
        errors: {},
        status_code: 0,
        user: {
          id: null,
          firstname: "",
          lastname: "",
          email: "",
          created_at: 0,
          lang: "",
          broadcasted_messages: 0,
          Createdby: 0
        },
        role: "",
        token: "",
        errormsg:""
  //    },
     
    };
  }

  onChangePassword(e) {
    this.setState({
        password: e.target.value
    });
  }


  onChangeEmail(e) {
    this.setState({
      email: e.target.value
    });
  }

  onHandleSubmit(event) {
    event.preventDefault();
    var data = {
      email: this.state.email,
      password: this.state.password
    };

    console.log(data);
   /// const history = useHistory()

    UserService.userlogin(data)
      .then(response => {
        console.log(response.data);
       // this.props.history.push('/products')
       if (response.data.status_code === 200){
        this.setState({
          // id: response.data.id,
          // firstname: response.data.firstname,
          // lastname: response.data.lastname,
          // email: response.data.email,

          msg: response.data.msg,
          errors: response.data.errors,
          status_code: response.data.status_code,
          user: {
            id: response.data.user.id,
            firstname: response.data.user.firstname,
            lastname: response.data.user.lastname,
            email: response.data.user.email,
            created_at: response.data.user.created_at,
            lang: response.data.user.lang,
            password: response.data.user.password,
            broadcasted_messages: response.data.user.broadcasted_messages,
            Createdby: response.data.user.created_at
          },
          role: response.data.role,
          token: response.data.token
        });

        var data = this.state;
        console.log(this.state.user);
        console.log(data);
        console.log(data.token);
        window.token = data.token;
        window.id = data.user.id;
        console.log(data.user.id);
        console.log(window.id);

       // this.props.history.push('/products')
        if (response.data.role === "infoadmin"){
          // this.props.history.push({pathname: '/products', state: data.token})
          this.props.history.push('/products')
        }
        else if (response.data.role === "superadmin"){
         // this.props.history.push({pathname: '/super-admin/dic', state: data.token})
         this.props.history.push('/super-admin/products')

        }
        else if (response.data.role === "admin"){
          this.props.history.push('/super-admin/products')
        }
        else if (response.data.role === "agent"){
          this.props.history.push('/super-admin/products')
        }
        else if (response.data.role === "merchant"){
          this.props.history.push('/super-admin/products')
        }
        else {
          this.props.history.push('/products')
        }
        console.log(response.data);
        console.log(this.state);
        console.log(this.state.token);

      }else {
        this.setState({
          msg: response.data.msg,
          errors: response.data.errors,
          status_code: response.data.status_code
        })
        console.log(response.data);
        console.log(this.state);
      }
      })
      .catch(e => {
        console.log(e);
        console.log(e.data);
        this.setState({
          errormsg: "Incorrect User Credential. Please try again!!",
          // msg:e.msg,
          // errors: e.errors
        })
        console.log(e);
        console.log(this.state)
      });
  }

  newLogin() {
    this.setState({
      //   response_data: {
        msg: "",
        errors: {},
        status_code: 0,
        user: {
          id: null,
          firstname: "",
          lastname: "",
          email: "",
          created_at: 0,
          lang: "",
          password: "",
          broadcasted_messages: 0,
          Createdby: 0
        },
        role: "",
        token: ""
        
  //    },
      //submitted: false
    });
  }

  render() {
  const  {errormsg} = this.state;
  console.log(this.state);
   console.log(this.state.msg)
    return (
    
                <div id="login"> 
                        <div className="container">
                            <div className="row login-box">
 
                            {/* Baselogin */}
                            <div className="col-sm-5 bg-img align-self-left baselogin">
                                <div className="info">
                                    <div className="logo clearfix">
                                        <Link className="nav-brand" to="/">shemach</Link>
                                    </div>
                                </div>
                            </div>



                            {/* LoginForm */}

                            <div className="col-sm-7  bg-color align-self-center">
                                <div className="form-section">
                                    <div className="title">
                                    <h3>Sign into your account</h3>
                                    </div>
                                    <div className="login-inner-form">
                                        {/* <form> */}
                                        <form method="POST" onSubmit={this.onHandleSubmit}>
                                            <div className="form-group form-box">
                                                <input type="text" id="email" onChange={this.onChangeEmail} className="input-text" placeholder="Email Address" data-cy="email" />
                                                <i className="icon email"></i>
                                            </div>

                                            <div className="form-group form-box">
                                                <input type="text" id="password" onChange={this.onChangePassword} className="input-text" placeholder="Password" data-cy="password"/>
                                                <i className="icon lock"></i>
                                            </div>

                                            {/* {
                                                errorMessage && <ErrorAlter errorMessage={errorMessage} clearError={() => setError(undefined) }></ErrorAlter>
                                            } */}
                                             <p>{errormsg}</p>

                                            <div className="form-group">
                                            <button 
                                            className="loginbutton btn primary-btn" 
                                            data-cy="submit"
                                            >
                                              Login
                                            </button>
                                                {/* <button className="btn primary-btn" onClick={this.onHandleSubmit}>Login</button> */}
                                            </div>

                                        </form>
                                    </div>
                                </div>
                          </div>  
                         </div>
                        </div>
                </div>
    );
  }
}

export default withRouter(LoginPageImp);



