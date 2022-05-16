import React, { Component } from "react";
import ProductDataService from "../../services/prodcutService";
import { Link } from "react-router-dom";
import './superAdminRegisterAdmin.css'
import './headerInfo.css'    //For Header



export default class SuperAdminRegisterAdmin extends Component {
  constructor(props) {
    super(props);
    this.onChangeFirstName = this.onChangeFirstName.bind(this);
    this.onChangeLastName = this.onChangeLastName.bind(this);
   // this.onChangeProdArea = this.onChangeProdArea.bind(this);
    this.onChangeEmail = this.onChangeEmail.bind(this);
  //  this.onChangeCurrPrice = this.onChangeCurrPrice.bind(this);
    this.saveAdmin = this.saveAdmin.bind(this);
    this.newAdmin = this.newAdmin.bind(this);

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
        status_code: 0,
        errors: {},
        info_admin: {
          id: null,
          firstname: "",
          lastname: "",
          email: "",
          created_at: 0,
          lang: "",
          password: "",
          broadcasted_messages: "",
          Createdby: 0
        },
        
  //    },
      submitted: false
    };
  }

  onChangeFirstName(e) {
    this.setState({
      firstname: e.target.value
    });
  }

  onChangeLastName(e) {
    this.setState({
      lastname: e.target.value
    });
  }

  // onChangeProdArea(e) {
  //   this.setState({
  //     email: e.target.value
  //   });
  // }


  onChangeEmail(e) {
    this.setState({
      email: e.target.value
    });
  }


  // onChangeCurrPrice(e) {
  //   this.setState({
  //     currentprice: e.target.value
  //   });
  // }

  saveAdmin() {
    var data = {
      firstname: this.state.firstname,
      lastname: this.state.lastname,
      email: this.state.email,
      // measurement:this.state.measurement,
      // prevprice: this.state.prevprice,
      // currentprice: this.state.currentprice
    };

    console.log(data);

    ProductDataService.registerAdmin(data)
      .then(response => {
        this.setState({
          // id: response.data.id,
          // firstname: response.data.firstname,
          // lastname: response.data.lastname,
          // email: response.data.email,

          msg: response.data.msg,
          status_code: response.data.status_code,
          errors: response.data.errors,
          info_admin: {
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
          // measurement: response.data.measurement,
          // prevprice: response.data.prevprice,
          // currentprice: response.data.currentprice,
          // published: response.data.published,

          submitted: true
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  newAdmin() {
    this.setState({
      // id: null,
      // firstname: "",
      // lastname: "",
      // email: "",
      // measurement: "",
      // prevprice: 0,
      // currentprice: 0,
      // published: false,

      // submitted: false


      //   response_data: {
        msg: "",
        status_code: 0,
        errors: {},
        info_admin: {
          id: null,
          firstname: "",
          lastname: "",
          email: "",
          created_at: 0,
          lang: "",
          password: "",
          broadcasted_messages: "",
          Createdby: 0
        },
        
  //    },
      submitted: false
    });
  }

  render() {
    return (
        <>
         <header id="navitem">
                <nav class="navbar navbar-expand-lg">
                <div class="container">
                <Link class="navbar-brand text-white" to='/super-admin/control-admins'><i class="fa-solid fa-angles-left"></i> Home </Link>
                <div class="collapse navbar-collapse" id="nvbCollapse">
                    <ul class="navbar-nav ml-auto nav-menu">
                        <p class="nav-item pl-1 nav-link"> Super Admin Page To Register New Admin</p>
                    </ul>
                </div>
                </div>
                </nav>
	      </header>

      <div className="submit-form">
        {this.state.submitted ? (
          <div>
            <h2>You registered admin successfully!</h2>
            <button className="btn btn-success" onClick={this.newAdmin}>
              Register
            </button>
          </div>
        ) : (
          <div className="addform">
            <div className="form-group">
              <label htmlFor="firstname">First Name</label>
              <input
                type="text"
                className="form-control"
                id="firstname"
                required
                value={this.state.firstname}
                onChange={this.onChangeFirstName}
                name="firstname"
              />
            </div>

            <div className="form-group">
              <label htmlFor="lastname">Last Name</label>
              <input
                type="text"
                className="form-control"
                id="lastname"
                required
                value={this.state.lastname}
                onChange={this.onChangeLastName}
                name="lastname"
              />
            </div>
{/* 
            <div className="form-group">
              <label htmlFor="prodarea">Surname</label>
              <input
                type="text"
                className="form-control"
                id="prodarea"
                required
                value={this.state.email}
                onChange={this.onChangeProdArea}
                name="prodarea"
              />
            </div> */}

            <div className="form-group">
              <label htmlFor="lastname">Email</label>
              <input
                type="text"
                className="form-control"
                id="measurement"
                required
                value={this.state.measurement}
                onChange={this.onChangeEmail}
                name="measurement"
              />
            </div>
            {/* <div className="form-group">
              <label htmlFor="currprice">Phone</label>
              <input
                type="number"
                className="form-control"
                id="currprice"
                required
                value={this.state.currentprice}
                onChange={this.onChangeCurrPrice}
                name="currprice"
              />
            </div> */}


            <button onClick={this.saveAdmin} className="btn btn-success">
              Register
            </button>

            <div>
              <strong><h1>{this.state.msg}</h1></strong>
            </div>
          </div>
        )}
      </div>
    </>
    );
  }
}


