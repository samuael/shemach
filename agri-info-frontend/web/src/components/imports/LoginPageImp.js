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
        token: ""
  //    },
     
    };
  }




    // // handle Submit handler function
    // handleSubmit = (event) =>{
    //     event.preventDefault();

    //     const userCredential = {
    //         email,
    //         password
    //     }
    //     // const userdata = { email: "fantahunfekadu1@gmail.com", password : "admin123"}
    //     const login = dispatch(loginAction(userCredential))
    //     login
    //         .then(data => {
    //             if (data.role == "infoadmin"){
    //                 history.push('/products')
    //             }
    //             else if (data.role == "superadmin"){
    //                 history.push('/super-admin/products')
    //             }
    //             else if (data.role == "admin"){
    //                 history.push('/super-admin/products')
    //             }
    //             else if (data.role == "agent"){
    //                 history.push('/super-admin/products')
    //             }
    //             else if (data.role == "merchant"){
    //                 history.push('/super-admin/products')
    //             }
    //             else {
    //                 history.push('/super-admin/products')
    //             }
    //             console.log(data.role);
                
    //         }
            
    //         )
    //         .catch(error => { 
    //             console.log(error.err)
    //             setError(error.err)
    //         })
    // }







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
    //   firstname: this.state.firstname,
    //   lastname: this.state.lastname,
      email: this.state.email,
      password: this.state.password
      // measurement:this.state.measurement,
      // prevprice: this.state.prevprice,
      // currentprice: this.state.currentprice
    };

    console.log(data);
   /// const history = useHistory()

    UserService.userlogin(data)
      .then(response => {
       // this.props.history.push('/products')
        this.setState({
          // id: response.data.id,
          // firstname: response.data.firstname,
          // lastname: response.data.lastname,
          // email: response.data.email,

          msg: response.data.msg,
          errors: response.data.errors,
          status_code: response.data.status_code,
          user: {
            id: response.data.id,
            firstname: response.data.firstname,
            lastname: response.data.lastname,
            email: response.data.email,
            created_at: response.data.created_at,
            lang: response.data.lang,
            password: response.data.password,
            broadcasted_messages: response.data.broadcasted_messages,
            Createdby: response.data.created_at
          },
          role: response.data.role,
          token: response.data.token
          // measurement: response.data.measurement,
          // prevprice: response.data.prevprice,
          // currentprice: response.data.currentprice,
          // published: response.data.published,

        //   submitted: true
        });

        const data = this.state;
        console.log(data);

       // this.props.history.push('/products')
        if (response.data.role === "infoadmin"){
          this.props.history.push({pathname: '/products', state: data})
        }
        else if (response.data.role === "superadmin"){
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
          this.props.history.push('/super-admin/products')
        }
        console.log(response.data);
        console.log(this.state);
        console.log(this.state.token);
      })
      .catch(e => {
        console.log(e);
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
                                        <Link className="nav-brand" to="/">Agri-Net</Link>
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
                                                <input type="text" id="email" onChange={this.onChangeEmail} className="input-text" placeholder="Email Address" />
                                                <i className="icon email"></i>
                                            </div>

                                            <div className="form-group form-box">
                                                <input type="text" id="password" onChange={this.onChangePassword} className="input-text" placeholder="Password" />
                                                <i className="icon lock"></i>
                                            </div>

                                            {/* {
                                                errorMessage && <ErrorAlter errorMessage={errorMessage} clearError={() => setError(undefined) }></ErrorAlter>
                                            } */}
                                             <p>{this.state.msg}</p>

                                            <div className="form-group">
                                            <button className="btn primary-btn">Login</button>

                                           

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



