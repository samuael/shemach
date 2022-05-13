import React, { Component } from "react";
import ProductService from "../../services/prodcutService";
import './productComponent.css'
import './headerInfo.css'   //For Header
import {Switch, Route, Link} from 'react-router-dom'

export default class ProductComponent extends Component {
  constructor(props) {
    super(props);
    this.onChangeTitle = this.onChangeTitle.bind(this);
    this.onChangeDescription = this.onChangeDescription.bind(this);
    this.onChangeProductionArea = this.onChangeProductionArea.bind(this);
    this.onChangeMeasurement = this.onChangeMeasurement.bind(this);
    this.onChangePrevPrice = this.onChangePrevPrice.bind(this);
    this.onChangeCurrPrice = this.onChangeCurrPrice.bind(this);
    this.getProduct = this.getProduct.bind(this);
    // this.updatePublished = this.updatePublished.bind(this);
    this.updateProduct = this.updateProduct.bind(this);
    this.deleteProduct = this.deleteProduct.bind(this);

    this.state = {
      res_data: {
        status_code: 0,
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
        msg: ""
      }
      // status_code:0,
      // product: {
      //   id: null,
      //   name: "",
      //   production_area: "",
      //   unit_id: 0,
      //   current_price: 0,
      //   created_by:0,
      //   created_at: 0,
      //   last_update_time: 0
      // },
      // msg: ""
    };
  }

  componentDidMount() {
    this.getProduct(this.props.match.params.id);
  }

  onChangeTitle(e) {
    const name = e.target.value;

    this.setState(function(prevState) {
      return {
        product: {
          ...prevState.product,
          name: name
        }
      };
    });
  }

  onChangeDescription(e) {
    const production_area = e.target.value;
    
    this.setState(prevState => ({
      product: {
        ...prevState.product,
        description: production_area
      }
    }));
  }
  onChangeProductionArea(e) {
    const productionarea = e.target.value;
    
    this.setState(prevState => ({
      product: {
        ...prevState.product,
        productionarea: productionarea
      }
    }));
  }

  onChangeMeasurement(e) {
    const measurement = e.target.value;
    
    this.setState(prevState => ({
      product: {
        ...prevState.product,
        measurement: measurement
      }
    }));
  }

  onChangePrevPrice(e) {
    const prevprice = e.target.value;
    
    this.setState(prevState => ({
      product: {
        ...prevState.product,
        prevprice: prevprice
      }
    }));
  }

  onChangeCurrPrice(e) {
    const current_price = e.target.value;
    
    this.setState(prevState => ({
      product: {
        ...prevState.product,
        current_price: current_price
      }
    }));
  }


  getProduct(id) {
    ProductService.get(id)
      .then(response => {
        this.setState({
          res_data: response.data
        });
        console.log(response.data);
        console.log(response.data.product);
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  // updatePublished(status) {
  //   var data = {
  //     id: this.state.product.id,
  //     title: this.state.product.title,
  //     description: this.state.product.description,
  //     productionarea: this.state.product.productionarea,
  //     prevprice: this.state.product.prevprice,
  //     currentprice: this.state.product.currentprice,
  //     published: status
  //   };

  //   ProductService.update(this.state.product.id, data)
  //     .then(response => {
  //       this.setState(prevState => ({
  //         product: {
  //           ...prevState.product,
  //           published: status
  //         }
  //       }));
  //       console.log(response.data);
  //     })
  //     .catch(e => {
  //       console.log(e);
  //     });
  // }

  updateProduct() {
    ProductService.update(
      this.state.product.id,
      this.state.product
    )
      .then(response => {
        console.log(response.data);
        this.setState({
          message: "The product was updated successfully!"
        });
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
    const { res_data } = this.state;
   // console.log(res_data);
   // console.log(res_data.product);

    return (
      <>
                <header id="navitem">
                <nav class="navbar navbar-expand-lg">
                <div class="container">
                {/* <Link to='/' className='navbar-logoo'><i class="fa-solid fa-list"></i>  Agri-Net <i class="fa-solid fa-building-wheat"></i> </Link> */}
                <Link class="navbar-brand text-white" to='/products'><i class="fa-solid fa-angles-left"></i> Products </Link>
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
        {res_data.product ? (
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
                  value={res_data.product.name}
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
                  value={res_data.product.unit_id}
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
                  value={res_data.product.production_area}
                  // onChange={this.onChangeProductionArea}
                />
              </div>
              <div className="form-group">
                <label htmlFor="measurement">Created By</label>
                <input
                  type="text"
                  className="form-control"
                  disabled="disabled"
                  id="measurement"
                  value={res_data.product.created_by}
                  // onChange={this.onChangeMeasurement}
                />
              </div>
              </div>

              <div className="col-sm-6">
            
              <div className="form-group">
                <label htmlFor="prevprice">Created At</label>
                <input
                  type="number"
                  className="form-control"
                  disabled="disabled"
                  id="prevprice"
                  value={res_data.product.created_at}
                 // onChange={this.onChangePrevPrice}
                />
              </div>
              <div className="form-group">
                <label htmlFor="currentprice">Curr Price</label>
                <input
                  type="number"
                  className="form-control"
                  id="currentprice"
                  value={res_data.product.current_price}
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
              className="badge badge-success btn btn-primary"
            
              onClick={this.updateProduct}
            >
              Update
            </button>
            <p>{this.state.message}</p>
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
