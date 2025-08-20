import React, { Component } from "react";
import ProductService from "../../services/prodcutService";
import './productComponent.css'
import './headerInfo.css'   //For Header
import {Switch, Route, Link} from 'react-router-dom'

export default class ProductComponent extends Component {
  constructor(props) {
    super(props);
    this.onChangeCurrPrice = this.onChangeCurrPrice.bind(this);
    this.getProduct = this.getProduct.bind(this);
    this.convertToDate = this.convertToDate.bind(this);
    this.updateProduct = this.updateProduct.bind(this);
    this.deleteProduct = this.deleteProduct.bind(this);

    this.state = {
      tkValue: this.props.location.state,
      status_code: 0,
      currentProduct: { 
          id: null,
          name: "",
          production_area: "",
          unit_id: 0,
          current_price: 0,
          created_by:0,
          created_at: 0,
          last_update_time: 0
        },
        msg: ""
    };
  }

  componentDidMount() {
    // this.getProduct(this.props.match.params.id);
    console.log(this.tkvalue);
    console.log(this.props.location.state);

    var tokenn = window.token;
     if(!tokenn){
      this.props.history.push('/')
     }else{
      this.getProduct(this.props.match.params.id);
     }
  }

  // onChangeTitle(e) {
  //   const name = e.target.value;

  //   this.setState(function(prevState) {
  //     return {
  //       res_data: {
  //         ...prevState.res_data,
  //         product: {
  //           ...prevState.product,
  //           name: name
  //         }
  //       }
  //     };
  //   });
  // }

  convertToDate(id){
    var myDate = new Date(id*1000);
    var convertedDate = myDate.toLocaleDateString();
    console.log(convertedDate);
    return convertedDate;
  }

  // convertToDate(number){
  //   var myDate = new Date( number *1000);
  //   var convertedDate = myDate.toLocaleString();
  //   console.log(convertedDate);
  //   return convertedDate;
  //   // return myDate.toGMTString()+ "<br>" + myDate.toLocaleString();
  //   // document.write(myDate.toGMTString()+"<br>"+myDate.toLocaleString());
  
  // }

  onChangeCurrPrice(e) {
    const current_price = e.target.value;

    this.setState(function(prevState) {
      return {
        currentProduct: {
          ...prevState.currentProduct,
          current_price: current_price
        }
      };
    });
  }


  // onChangeCurrPrice(e) {
  //   this.setState({
  //     current_price: e.target.value
  //   });
  // }



  getProduct(id) {
    ProductService.get(id)
      .then(response => {
        this.setState({
        //  currentProduct: response.data
          status_code: response.data.status_code,
          currentProduct: { 
              id: response.data.product.id,
              name: response.data.product.name,
              production_area: response.data.product.production_area,
              unit_id: response.data.product.unit_id,
              current_price: response.data.product.current_price,
              created_by:response.data.product.created_by,
              created_at: response.data.product.created_at,
              last_update_time: response.data.product.last_update_time
            },
            msg:""
        });
        console.log(response.data);
        console.log(this.state);
        console.log(this.state.currentProduct);
       // console.log(response.data.product);
        
      })
      .catch(e => {
        console.log(e);
      });
  }

  updateProduct() {
    var data = {
      id:parseInt(this.state.currentProduct.id),
      cost:parseInt(this.state.currentProduct.current_price)

    }

  //  var token = this.state.tkValue;
  var token = window.token

    
    console.log(data);
    console.log(token);
    ProductService.update(data, token
      // this.state.currentProduct.id,
      // this.state.currentProduct.current_price
      // parseInt(this.state.currentProduct.id),
      // parseInt(this.state.currentProduct.current_price)
    )
      .then(response => {
        console.log(response.data);
        this.setState({
           msg: "The product was updated successfully!"
        });
        // this.setState({

        //   msg: "The product was updated successfully!"
        // });
      })
      .catch(e => {
        console.log(e);
      });
  }

  deleteProduct() {    
    ProductService.delete(this.state.product.id)
      .then(response => {
        console.log(response.data);
        this.props.history.push('/products')
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
   const {currentProduct} = this.state;
   console.log(this.state.currentProduct);
  //  console.log(res_data.product);
  //  console.log(res_data.msg);
  // console.log(res_data.product.name);

    return (
      <>
                <header id="navitem">
                <nav class="navbar navbar-expand-lg">
                <div class="container">
                {/* <Link to='/' className='navbar-logoo'><i class="fa-solid fa-list"></i>  shemach <i class="fa-solid fa-building-wheat"></i> </Link> */}
                <Link class="returnback navbar-brand text-white" to='/products'><i class="fa-solid fa-angles-left"></i> Products </Link>
                {/* <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#nvbCollapse" aria-controls="nvbCollapse">
                    <span class="navbar-toggler-icon"></span>
                </button> */}
                <div class="collapse navbar-collapse" id="nvbCollapse">
                    <ul class="navbar-nav ml-auto nav-menu">
                        {/* <li class="nav-item active pl-1">
                            <a class="nav-link" href="#"><i class="fa-brands fa-product-hunt"></i>Products <i class="fa-solid fa-circle-dollar"></i></a>
                        </li>
                        <li class="nav-item pl-1">
                            <a class="nav-link" href="#"><i class="fa-solid fa-hand-holding-dollar"></i>Product Price <i class="fa-solid fa-circle-dollar"></i></a>
                        </li>
                        <li class="nav-item pl-1">
                            <a class="nav-link" href="#"><i class="fa-solid fa-bullhorn"></i>Broadcast Info </a><span class="circle">5</span>
                        </li> */}
                        <p class="nav-item pl-1 nav-link">Info Admin Price Update Page</p>
                    </ul>
                </div>
                </div>
                </nav>
	</header>

      <div id="procomponent">
        {currentProduct ? (
          <div className="edit-form mt-4">
            <h4>Product</h4>
            <form className="row infoeditprice">
              <div className="col-sm-6">
              <div className="form-group">
                <label htmlFor="title">Name</label>
                <input
                  type="text"
                  className="form-control"
                  disabled="disabled"
                  id="title"
                  value={currentProduct.name}
                  onChange={this.onChangeTitle}
                />
              </div>
              <div className="form-group">
                <label htmlFor="description">Unit ID</label>
                <input
                  type="text"
                  className="form-control"
                  disabled="disabled"
                  id="description"
                  value={currentProduct.unit_id}
                  // onChange={this.onChangeDescription}
                />
              </div>
              <div className="form-group">
                <label htmlFor="productionarea">Prod Area</label>
                <input
                  type="text"
                  className="form-control"
                  disabled="disabled"
                  id="productionarea"
                  value={currentProduct.production_area}
                  // onChange={this.onChangeProductionArea}
                />
              </div>
              <div className="form-group mb-3">
                <label htmlFor="measurement">Created By</label>
                <input
                  type="text"
                  className="form-control"
                  disabled="disabled"
                  id="measurement"
                  value={currentProduct.created_by}
                  // onChange={this.onChangeMeasurement}
                />
              </div>
              </div>

              <div className="col-sm-6">
            
              <div className="form-group">
                <label htmlFor="prevprice">Created At</label>
                <input
                  type="text"
                  className="form-control"
                  disabled="disabled"
                  id="prevprice"
                  value={this.convertToDate(currentProduct.created_at)}
                />
              </div>
              <div className="form-group">
                <label htmlFor="currentprice">Curr Price</label>
                <input
                  type="number"
                  className="form-control"
                  id="currentprice"
                  value={this.state.currentProduct.current_price}
                  data-cy="currentproprice"
                 onChange={this.onChangeCurrPrice}
                />
              </div>

              {/* <div className="form-group">
                <label>
                  <strong>Status:</strong>
                </label>
                {product.published ? "Published" : "Pending"}
              </div> */}
              </div>
            </form>

            {/* {product.published ? (
              <button
                className="badge badge-primary mr-2"
                onClick={() => this.updatePublished(false)}
              >
                UnPublish
              </button>
            ) : (
              <button
                className="badge badge-primary mr-2"
                onClick={() => this.updatePublished(true)}
              >
                Publish
              </button>
            )} */}

            {/* <button
              className="badge badge-danger mr-2"
              onClick={this.deleteProduct}
            >
              Delete
            </button> */}
            {/* <p>{this.state.message}</p> */}

            <button
              type="submit"
              className="updatebutton badge badge-success btn btn-primary mt-3"
              data-cy="updatepricebutton"
              onClick={this.updateProduct}
            >
              Update
            </button>
            <p>{this.state.msg}</p>
          </div>
        ) : (
          <div>
            <br />
            <p>Please click on a Product...</p>
          </div>
        )}
      </div>
      </>
    );
  }
}
