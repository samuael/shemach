import React, { Component } from "react";
import ProductService from "../../services/prodcutService";
import './superProductEdit.css'
import './headerInfo.css'   //For Header
import {Switch, Route, Link} from 'react-router-dom'

export default class SuperProductEdit extends Component {
  constructor(props) {
    super(props);
    this.onChangeTitle = this.onChangeTitle.bind(this);
    this.onChangeDescription = this.onChangeDescription.bind(this);
    this.onChangeProductionArea = this.onChangeProductionArea.bind(this);
    this.onChangeMeasurement = this.onChangeMeasurement.bind(this);
    this.onChangePrevPrice = this.onChangePrevPrice.bind(this);
    this.onChangeCurrPrice = this.onChangeCurrPrice.bind(this);
    this.getProduct = this.getProduct.bind(this);
    this.updatePublished = this.updatePublished.bind(this);
    this.updateProduct = this.updateProduct.bind(this);
    this.deleteProduct = this.deleteProduct.bind(this);

    this.state = {
      currentProduct: {
        id: null,
        title: "",
        description: "",
        productionarea: "",
        measurement: "",
        prevprice: 0,
        currentprice:0,
        published: false
      },
      message: ""
    };
  }

  componentDidMount() {
    this.getProduct(this.props.match.params.id);
  }

  onChangeTitle(e) {
    const title = e.target.value;

    this.setState(function(prevState) {
      return {
        currentProduct: {
          ...prevState.currentProduct,
          title: title
        }
      };
    });
  }

  onChangeDescription(e) {
    const description = e.target.value;
    
    this.setState(prevState => ({
      currentProduct: {
        ...prevState.currentProduct,
        description: description
      }
    }));
  }
  onChangeProductionArea(e) {
    const productionarea = e.target.value;
    
    this.setState(prevState => ({
      currentProduct: {
        ...prevState.currentProduct,
        productionarea: productionarea
      }
    }));
  }

  onChangeMeasurement(e) {
    const measurement = e.target.value;
    
    this.setState(prevState => ({
      currentProduct: {
        ...prevState.currentProduct,
        measurement: measurement
      }
    }));
  }

  onChangePrevPrice(e) {
    const prevprice = e.target.value;
    
    this.setState(prevState => ({
      currentProduct: {
        ...prevState.currentProduct,
        prevprice: prevprice
      }
    }));
  }

  onChangeCurrPrice(e) {
    const currentprice = e.target.value;
    
    this.setState(prevState => ({
      currentProduct: {
        ...prevState.currentProduct,
        currentprice: currentprice
      }
    }));
  }


  getProduct(id) {
    ProductService.get(id)
      .then(response => {
        this.setState({
          currentProduct: response.data
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  updatePublished(status) {
    var data = {
      id: this.state.currentProduct.id,
      title: this.state.currentProduct.title,
      description: this.state.currentProduct.description,
      productionarea: this.state.currentProduct.productionarea,
      prevprice: this.state.currentProduct.prevprice,
      currentprice: this.state.currentProduct.currentprice,
      published: status
    };

    ProductService.update(this.state.currentProduct.id, data)
      .then(response => {
        this.setState(prevState => ({
          currentProduct: {
            ...prevState.currentProduct,
            published: status
          }
        }));
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  updateProduct() {
    ProductService.update(
      this.state.currentProduct.id,
      this.state.currentProduct
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
    ProductService.delete(this.state.currentProduct.id)
      .then(response => {
        console.log(response.data);
        this.props.history.push('/super-admin/products')
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
    const { currentProduct } = this.state;

    return (
      <>
                <header id="navitem">
                <nav class="navbar navbar-expand-lg">
                <div class="container">
                {/* <Link to='/' className='navbar-logoo'><i class="fa-solid fa-list"></i>  Agri-Net <i class="fa-solid fa-building-wheat"></i> </Link> */}
                <Link class="navbar-brand text-white" to='/super-admin/products'><i class="fa-solid fa-angles-left"></i> Products </Link>
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
                        <p class="nav-item pl-1 nav-link">Super Admin Product Update Page</p>
                    </ul>
                </div>
                </div>
                </nav>
	</header>

      <div id="superproedit">
        {currentProduct ? (
          <div className="edit-form mt-4">
            <h4>Product</h4>
            <form className="row supeditform">
              <div className="col-md-5">
              <div className="form-group">
                <label htmlFor="title">Name</label>
                <input
                  type="text"
                  className="form-control"
                  id="title"
                  value={currentProduct.title}
                  onChange={this.onChangeTitle}
                />
              </div>
              <div className="form-group">
                <label htmlFor="description">Description</label>
                <input
                  type="text"
                  className="form-control"
                  id="description"
                  value={currentProduct.description}
                  onChange={this.onChangeDescription}
                />
              </div>
           </div>

           <div className="col-md-5">
              <div className="form-group">
                <label htmlFor="productionarea">Prod Area</label>
                <input
                  type="text"
                  className="form-control"
                  id="productionarea"
                  value={currentProduct.productionarea}
                  onChange={this.onChangeProductionArea}
                />
              </div>
              <div className="form-group">
                <label htmlFor="measurement">Measurement</label>
                <input
                  type="text"
                  className="form-control"
                  id="measurement"
                  value={currentProduct.measurement}
                  onChange={this.onChangeMeasurement}
                />
              </div>
            </div>

            <div className="col-md-5">
            
              <div className="form-group">
                <label htmlFor="prevprice">Prev Price</label>
                <input
                  type="number"
                  className="form-control"
                  id="prevprice"
                  value={currentProduct.prevprice}
                  onChange={this.onChangePrevPrice}
                />
              </div>
              <div className="form-group">
                <label htmlFor="currentprice">Curr Price</label>
                <input
                  type="number"
                  className="form-control"
                  id="currentprice"
                  value={currentProduct.currentprice}
                  onChange={this.onChangeCurrPrice}
                />
              </div>

              </div>

              <div className="col-md-5">
              <div className="form-group">
                <label>
                  <strong>Status:</strong>
                </label>
                {currentProduct.published ? "Published" : "Pending"}
              </div>
              </div>
            </form>

            {currentProduct.published ? (
              <button
                className="badge badge-primary mr-2 btn btn-primary"
                onClick={() => this.updatePublished(false)}
              >
                UnPublish
              </button>
            ) : (
              <button
                className="badge badge-primary mr-2 btn btn-primary"
                onClick={() => this.updatePublished(true)}
              >
                Publish
              </button>
            )}

            <button
              className="badge badge-danger mr-2 btn btn-primary"
              onClick={this.deleteProduct}
            >
              Delete
            </button>
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
