import React, { Component } from "react";
import ProductDataService from "../../services/prodcutService";
import { Link } from "react-router-dom";
import './superAdminAddPro.css'
import './headerInfo.css'    //For Header



export default class SuperAdminAddPr extends Component {
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
      name: "",
      production_area: "", 
      unit_id: 0,
      current_price: 0,
      created_by: 0,
      created_at: 0,
      last_update_time:0
    };
  }

  onChangeTitle(e) {
    this.setState({
      name: e.target.value
    });
  }

  onChangeDescription(e) {
    this.setState({
      unit_id: e.target.value
    });
  }

  onChangeProdArea(e) {
    this.setState({
      production_area: e.target.value
    });
  }
  onChangeMeasurement(e) {
    this.setState({
      measurement: e.target.value
    });
  }
  onChangeCurrPrice(e) {
    this.setState({
      current_price: e.target.value
    });
  }

  saveProduct() {
    var data = {
      name: this.state.name,
      unit_id: this.state.unit_id,
      production_area: this.state.production_area,
      current_price:this.state.current_price
      // prevprice: this.state.prevprice,
      // currentprice: this.state.currentprice
    };

    ProductDataService.create(data)
      .then(response => {
        this.setState({
          id: response.data.id,
          name: response.data.name,
          unit_id: response.data.unit_id,
          production_area: response.data.production_area,
          current_price: response.data.current_price,
          created_by: response.data.created_by,
          created_at: response.data.created_at,
          last_update_time: response.data.last_update_time

         // submitted: true
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
      name: "",
      production_area: "", 
      unit_id: 0,
      current_price: 0,
      created_by: 0,
      created_at: 0,
      last_update_time:0
    });
  }

  render() {
    return (
        <>
         <header id="navitem">
                <nav class="navbar navbar-expand-lg">
                <div class="container">
                <Link class="navbar-brand text-white" to='/super-admin/products'><i class="fa-solid fa-angles-left"></i> Home </Link>
                <div class="collapse navbar-collapse" id="nvbCollapse">
                    <ul class="navbar-nav ml-auto nav-menu">
                        <p class="nav-item pl-1 nav-link"> Super Admin Page To Add New Product</p>
                    </ul>
                </div>
                </div>
                </nav>
	      </header>

      <div className="submit-form">
        {this.state.submitted ? (
          <div>
            <h2>You added product successfully!</h2>
            <button className="btn btn-success" onClick={this.newProduct}>
              Add
            </button>
          </div>
        ) : (
          <div className="addform">
            <div className="form-group">
              <label htmlFor="title">Name</label>
              <input
                type="text"
                className="form-control"
                id="title"
                required
                value={this.state.name}
                onChange={this.onChangeTitle}
                name="title"
              />
            </div>

            <div className="form-group">
              <label htmlFor="description">Unit ID</label>
              <input
                type="text"
                className="form-control"
                id="description"
                required
                value={this.state.unit_id}
                onChange={this.onChangeDescription}
                name="description"
              />
            </div>

            <div className="form-group">
              <label htmlFor="prodarea">Prod Area</label>
              <input
                type="text"
                className="form-control"
                id="prodarea"
                required
                value={this.state.production_area}
                onChange={this.onChangeProdArea}
                name="prodarea"
              />
            </div>

            {/* <div className="form-group">
              <label htmlFor="description">Measurement</label>
              <input
                type="text"
                className="form-control"
                id="measurement"
                required
                value={this.state.measurement}
                onChange={this.onChangeMeasurement}
                name="measurement"
              />
            </div> */}
            <div className="form-group">
              <label htmlFor="currprice">Price</label>
              <input
                type="number"
                className="form-control"
                id="currprice"
                required
                value={this.state.current_price}
                onChange={this.onChangeCurrPrice}
                name="currprice"
              />
            </div>


            <button onClick={this.saveProduct} className="btn btn-success">
              Submit
            </button>
          </div>
        )}
      </div>
    </>
    );
  }
}

