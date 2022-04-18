import React, { Component } from "react";
import ProductDataService from "../../services/prodcutService";
import { Link } from "react-router-dom";
import './superAdminRegisterAdmin.css'
import './headerInfo.css'    //For Header



export default class SuperAdminRegisterAdmin extends Component {
  constructor(props) {
    super(props);
    this.onChangeTitle = this.onChangeTitle.bind(this);
    this.onChangeDescription = this.onChangeDescription.bind(this);
    this.onChangeProdArea = this.onChangeProdArea.bind(this);
    this.onChangeMeasurement = this.onChangeMeasurement.bind(this);
    this.onChangeCurrPrice = this.onChangeCurrPrice.bind(this);
    this.saveProduct = this.saveProduct.bind(this);
    this.newProduct = this.newProduct.bind(this);

    this.state = {
      id: null,
      title: "",
      description: "", 
      productionarea: "",
      measurement: "",
      prevprice: 0,
      currentprice:0,
      published: false,

      submitted: false
    };
  }

  onChangeTitle(e) {
    this.setState({
      title: e.target.value
    });
  }

  onChangeDescription(e) {
    this.setState({
      description: e.target.value
    });
  }

  onChangeProdArea(e) {
    this.setState({
      productionarea: e.target.value
    });
  }
  onChangeMeasurement(e) {
    this.setState({
      measurement: e.target.value
    });
  }
  onChangeCurrPrice(e) {
    this.setState({
      currentprice: e.target.value
    });
  }

  saveProduct() {
    var data = {
      title: this.state.title,
      description: this.state.description,
      productionarea: this.state.productionarea,
      measurement:this.state.measurement,
      prevprice: this.state.prevprice,
      currentprice: this.state.currentprice
    };

    ProductDataService.create(data)
      .then(response => {
        this.setState({
          id: response.data.id,
          title: response.data.title,
          description: response.data.description,
          productionarea: response.data.productionarea,
          measurement: response.data.measurement,
          prevprice: response.data.prevprice,
          currentprice: response.data.currentprice,
          published: response.data.published,

          submitted: true
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  newProduct() {
    this.setState({
      id: null,
      title: "",
      description: "",
      productionarea: "",
      measurement: "",
      prevprice: 0,
      currentprice: 0,
      published: false,

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
            <button className="btn btn-success" onClick={this.newProduct}>
              Register
            </button>
          </div>
        ) : (
          <div className="addform">
            <div className="form-group">
              <label htmlFor="title">First Name</label>
              <input
                type="text"
                className="form-control"
                id="title"
                required
                value={this.state.title}
                onChange={this.onChangeTitle}
                name="title"
              />
            </div>

            <div className="form-group">
              <label htmlFor="description">Last Name</label>
              <input
                type="text"
                className="form-control"
                id="description"
                required
                value={this.state.description}
                onChange={this.onChangeDescription}
                name="description"
              />
            </div>

            <div className="form-group">
              <label htmlFor="prodarea">Surname</label>
              <input
                type="text"
                className="form-control"
                id="prodarea"
                required
                value={this.state.productionarea}
                onChange={this.onChangeProdArea}
                name="prodarea"
              />
            </div>

            <div className="form-group">
              <label htmlFor="description">Email</label>
              <input
                type="text"
                className="form-control"
                id="measurement"
                required
                value={this.state.measurement}
                onChange={this.onChangeMeasurement}
                name="measurement"
              />
            </div>
            <div className="form-group">
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
            </div>


            <button onClick={this.saveProduct} className="btn btn-success">
              Register
            </button>
          </div>
        )}
      </div>
    </>
    );
  }
}


