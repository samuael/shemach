import React, { Component } from "react";
import ProductDataService from "../../services/prodcutService";
import { Link } from "react-router-dom";
import './superAdminAddPro.css'
import './headerInfo.css'    //For Header



export default class SuperAdminAddPr extends Component {
  constructor(props) {
    super(props);
    this.onChangeProdName = this.onChangeProdName.bind(this);
    this.onChangeProdUnitId = this.onChangeProdUnitId.bind(this);
    this.onChangeProdArea = this.onChangeProdArea.bind(this);
   // this.onChangeMeasurement = this.onChangeMeasurement.bind(this);
    this.onChangeCurrPrice = this.onChangeCurrPrice.bind(this);
    this.saveProduct = this.saveProduct.bind(this);
    this.newProduct = this.newProduct.bind(this);
  //  this.response_data = this.response_data.bind(this);

    this.state = {
  //    tkValue: this.props.location.state,
    //  response_data: {
        msg: "",
        product: {
          id: null,
          name: "",
          production_area: "",
          unit_id: 0,
          current_price: 0,
          created_by:0,
          created_at: 0,
          last_update_time: 0
        },
        status_code: 0,
    //  },
      submitted: false
    };
  }

  onChangeProdName(e) {
    this.setState({
      name: e.target.value
      });
  }



  // onChangeProdName(e) {
  //   const name = e.target.value;

  //   this.setState(function(prevState) {
  //     return {
  //       response_data: {
  //         ...prevState.response_data,
  //         product: {
  //           ...prevState.product,
  //           name: name
  //         }
  //       }
  //     };
  //   });
  // }

  // onChangeProdName(e) {
  //   const name = e.target.value;
    
  //   this.setState(prevState => ({
  //     response_data: {
  //       ...prevState.response_data,
  //       product: {
  //         ...prevState.product,
  //         name: name
  //       }
  //     }
  //   }));
  // }



  // onChangeProdUnitId(e) {
  //   this.setState({
  //     unit_id: e.target.value
  //   });
  // }

  // onChangeProdUnitId(e) {
  //   const unit_id = e.target.value;

  //   this.setState(function(prevState) {
  //     return {
  //       response_data: {
  //         ...prevState.response_data,
  //         product: {
  //           ...prevState.product,
  //           unit_id: unit_id
  //         }
  //       }
  //     };
  //   });
  // }

  // onChangeProdUnitId(e) {
  //   const unit_id = e.target.value;
    
  //   this.setState(prevState => ({
  //     response_data: {
  //       ...prevState.response_data,
  //       product: {
  //         ...prevState.product,
  //         unit_id: unit_id
  //       }
  //     }
  //   }));
  // }


  // onChangeProdUnitId(e) {
  //  // const unit_id = e.target.value;
    
  //   this.setState({
  //       product: {
  //         unit_id : e.target.value
  //       }
  //     }
  //   );
  // }

  onChangeProdUnitId(e) {
    this.setState({
      unit_id: e.target.value
      });
  }


  

  // onChangeProdArea(e) {
  //   this.setState({
  //     production_area: e.target.value
  //   });
  // }

  // onChangeProdArea(e) {
  //   const production_area = e.target.value;

  //   this.setState(function(prevState) {
  //     return {
  //       response_data: {
  //         ...prevState.response_data,
  //         product: {
  //           ...prevState.product,
  //           production_area: production_area
  //         }
  //       }
  //     };
  //   });
  // }

  // onChangeProdArea(e) {
  //   const production_area = e.target.value;
    
  //   this.setState(prevState => ({
  //     response_data: {
  //       ...prevState.response_data,
  //       product: {
  //         ...prevState.product,
  //         production_area: production_area
  //       }
  //     }
  //   }));
  // }


  // onChangeProdArea(e) {
  //   //const unit_id = e.target.value;
    
  //   this.setState({
  //       product: {
  //         production_area: e.target.value
  //       }
  //     }
  //   );
  // }
  

  onChangeProdArea(e) {
    this.setState({
      production_area: e.target.value
      });
  }

  

  // onChangeCurrPrice(e) {
  //   this.setState({
  //     current_price: e.target.value
  //   });
  // }

  // onChangeCurrPrice(e) {
  //   const current_price = e.target.value;

  //   this.setState(function(prevState) {
  //     return {
  //       response_data: {
  //         ...prevState.response_data,
  //         product: {
  //           ...prevState.product,
  //           current_price: current_price
  //         }
  //       }
  //     };
  //   });
  // }

  // onChangeCurrPrice(e) {
  //   const current_price = e.target.value;
    
  //   this.setState(prevState => ({
  //     response_data: {
  //       ...prevState.response_data,
  //       product: {
  //         ...prevState.product,
  //         current_price: current_price
  //       }
  //     }
  //   }));
  // }

  // onChangeCurrPrice(e) {
  //   //const unit_id = e.target.value;
    
  //   this.setState({
  //       product: {
  //         current_price: e.target.value
  //       }
  //     }
  //   );
  // }


  onChangeCurrPrice(e) {
    this.setState({
      current_price: e.target.value
      });
  }
  

 


  saveProduct() {
    var data = {
      unit_id: parseInt(this.state.unit_id),
      name: this.state.name,
      production_area: this.state.production_area,
      current_price:parseInt(this.state.current_price)

      // unit_id: 6,
      // name: "Faniman",
      // production_area: "AA",
      // current_price:2000
    };

    // var token = this.state.tkValue;
    var token = window.token;
    console.log(data);
    console.log(token);

    ProductDataService.create(data, token)
      .then(response => {
        console.log(response.data)
        this.setState({

       // response_data: {parseInt(
          msg: response.data.msg,
          product: {
            id: response.data.product.id,
            name: response.data.product.name,
            production_area: response.data.product.production_area,
            unit_id: response.data.product.unit_id,
            current_price: response.data.product.current_price,
            created_by:response.data.product.current_price,
            created_at: response.data.product.created_at,
            last_update_time: response.data.product.last_update_time
          },
          status_code: response.data.status_code,
       // }
          submitted: true
        




          // id: response.data.response_data.id,
          // name: response.data.response_data.name,
          // unit_id: response.data.response_data.unit_id,
          // production_area: response.data.response_data.production_area,
          // current_price: response.data.response_data.current_price,
          // created_by: response.data.response_data.created_by,
          // created_at: response.data.response_data.created_at,
          // last_update_time: response.data.response_data.last_update_time

         // submitted: true
        });
        console.log(response.data);
      //  console.log(response_data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  newProduct() {


    this.setState({
            msg: "",
            product: {
              id: null,
              name: "",
              production_area: "",
              unit_id: 0,
              current_price: 0,
              created_by:0,
              created_at: 0,
              last_update_time: 0
            },
            status_code: 0,
        //  },
          submitted: false
        });
  }

  render() {
  //  const {product} = this.state;
  //  console.log(product);
//console.log(response_data.product);
    return (
        <>
         <header id="navitem">
                <nav class="navbar navbar-expand-lg">
                <div class="container">
                <Link data-cy="returnback" class="navbar-brand text-white" to='/super-admin/products'><i class="fa-solid fa-angles-left"></i> Home </Link>
                <div class="collapse navbar-collapse" id="nvbCollapse">
                    <ul class="navbar-nav ml-auto nav-menu">
                        <p data-cy="supernav" class="nav-item pl-1 nav-link"> Super Admin Page To Add New Product</p>
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
                onChange={this.onChangeProdName}
                name="title"
                data-cy="prodname"
              />
            </div>

            <div className="form-group">
              <label htmlFor="unitId">Unit ID</label>
              <select data-cy="produnit" name="unitId" id="unitId" required className="form-control" value={this.state.unit_id} onChange={this.onChangeProdUnitId} type="number">
                  <optgroup label="Mass">
                  <option value="" selected disabled hidden>Choose here</option>
                      <option value="1">Kg</option>
                      <option value="2">g</option>
                      <option value="3">Kun</option>
                      <option value="4">Ton</option>
                  </optgroup>

                  <optgroup label="Volume">
                      <option value="5">L</option>
                      <option value="6">M3</option>
                      <option value="7">Gal</option>
                 </optgroup>
                 <optgroup label="Item">
                      <option value="8">SIT</option>
                      <option value="9">DZ</option>
                      <option value="10">HDZ</option>
                      <option value="11">QDZ</option>
                 </optgroup>
                 <optgroup label="Size">
                      <option value="12">SM</option>
                      <option value="13">LG</option>
                      <option value="14">MD</option>
                 </optgroup>

                 <optgroup label="Length">
                      <option value="15">M</option>
                      <option value="16">KM</option>
                      <option value="17">CM</option>
                 </optgroup>
             </select>

              {/* <input
                type="number"
                className="form-control"
                id="unitId"
                required
                value={this.state.unit_id}
                onChange={this.onChangeProdUnitId}
                name="unitId"
              /> */}
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
                data-cy="prodarea"
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
                onChange={this.onChangeProdUnitId}
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
                data-cy="prodprice"
              />
            </div>


            <button  data-cy="submitprod" onClick={this.saveProduct} className="btn btn-success">
              Submit
            </button>
          </div>
        )}
      </div>
    </>
    );
  }
}

