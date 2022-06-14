import React, { Component } from "react";
import ProductService from '../../services/prodcutService'
import { withRouter } from 'react-router-dom'
import { Link } from "react-router-dom";
import './productsSuper.css'

class ProductsSuper extends Component {
  constructor(props) {
    super(props);
    this.onChangesearchProduct = this.onChangesearchProduct.bind(this);
    this.retrieveProducts = this.retrieveProducts.bind(this);
    this.refreshList = this.refreshList.bind(this);
    this.setActiveProduct = this.setActiveProduct.bind(this);
    this.removeAllProducts = this.removeAllProducts.bind(this);
    this.searchProduct = this.searchProduct.bind(this);
    this.renderUnitId = this.renderUnitId.bind(this);

    this.state = {
      tokenValue: this.props.location.state,
      res_data: {
        status_code: 0,
        products: [],
        msg: ""
      },
      productts: [
    //     {
    //     id: 1,
    //     title: 'Product 1',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //     id: 2,
    //     title: 'Product 2',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //     id: 3,
    //     title: 'Product 3',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //     id: 4,
    //     title: 'Product 4',
    //     description: 'Create tasks',
    //     published : false,
    // },

      ],
      currentProduct: null,
      currentIndex: -1,
      searchProduct: ""
    };
  }

  componentDidMount() {
    console.log(this.props.location.state);
    console.log(this.state.tokenValue);
    var tokenn = window.token;
    this.retrieveProducts();

    //  if(!tokenn){
    //   this.props.history.push('/')
    //  }else{
    //   this.retrieveProducts();
    //  }
  }

  onChangesearchProduct(e) {
    const searchProduct = e.target.value;

    this.setState({
      searchProduct: searchProduct
    });
  }

  retrieveProducts() {
    ProductService.getAll()
      .then(response => {
        this.setState({
          res_data: response.data
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  refreshList() {
    this.retrieveProducts();
    this.setState({
      currentProduct: null,
      currentIndex: -1
    });
  }

  setActiveProduct(tutorial, index) {
    this.setState({
      currentProduct: tutorial,
      currentIndex: index
    });
  }

  removeAllProducts() {
    ProductService.deleteAll()
      .then(response => {
        console.log(response.data);
        this.refreshList();
      })
      .catch(e => {
        console.log(e);
      });
  }

renderUnitId(id) {
    switch (id) {
      case 1:
        return "Kg"
      case 2:
        return "g"
      case 3:
        return "Kun"
      case 4:
        return "Ton"
      case 5:
        return "L"
      case 6:
        return "M3"
      case 7:
        return "Gal"
      case 8:
        return "SIT"
      case 9:
        return "DZ"
      case 10:
        return "HDZ"
      case 11:
        return "ODZ"
      case 12:
        return "SM"
      case 13:
        return "LG"
      case 14:
        return "MD"
      case 15:
        return "KM"
      case 16:
        return "Cm"
      case 17:
        return "Mile"
      default:
        return "Kun"
    }
}

  searchProduct() {
    this.setState({
      currentProduct: null,
      currentIndex: -1
    });

    ProductService.findByTitle(this.state.searchProduct)
      .then(response => {
        this.setState(prevState => ({
          res_data: {
            ...prevState.res_data,
            products: response.data
          } 
        }));
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
    const { searchProduct, res_data, currentProduct, currentIndex } = this.state;
    console.log(currentProduct);
    console.log(res_data);
    return (
    <div id="productsuper">
      <div className="list row">
        <div className="col-md-8">
          <div className="input-group mb-3">
            <input
              type="text"
              className="form-control"
              placeholder="Search by product name"
              value={searchProduct}
              onChange={this.onChangesearchProduct}
            />
            <div className="input-group-append">
              <button
                className="btn btn-outline-secondary searchbuttom"
                type="button"
                onClick={this.searchProduct}
              >
                Search
              </button>
            </div>
          </div>
       
        </div>
        <div className="col-md-4">
              <button className="btn btn-primary add-button">
              <Link  
              className="linkadd" 
                  // to="/super-admin/add-product"
                  to={{
                    pathname: "/super-admin/add-product",
                    state: this.state.tokenValue
                }}
              >
                    <i data-cy="superaddpro" className="fa-solid fa-plus"></i>Add Product
              </Link>
              </button>
          </div>
       
        </div>
    
      <div className="row">
        <div className="col-md-9">
          <h4>Product List</h4>

          <ul className="list-group">
            {
              res_data.products.map((product, index) => (
                <div className={
                  "list-group-item " +
                  (index === currentIndex ? "active" : "")
                }
                onClick={() => this.setActiveProduct(product, index)}
                key={index}
                >                
                    <div className="Name">
                      
                    Crop Name :{product.name}
                    </div>
                  <p className="Place">Price :{product.current_price}</p>
                  <p className="Price"> Before :{product.created_at} min</p>
                </div>

              ))}
          </ul>

          <button
            className="mt-4 btn btn-sm btn-danger removebutton"
            onClick={this.removeAllProducts}
          >
            Remove All
          </button>
        </div>
        <div className="col-md-3 description">
          {currentProduct ? (
            <div>
              <h4>Product</h4>
              <div>
                <label>
                  <strong>Name:</strong>
                </label>{" "}
                {currentProduct.name}
              </div>
              {/* <div>
                <label>
                  <strong>Description:</strong>
                </label>{" "}
                {currentProduct.description}
              </div> */}
              <div>
                <label>
                  <strong>Prod Area:</strong>
                </label>{" "}
                {currentProduct.production_area}
              </div>
              <div>
                <label>
                  <strong>Unit ID:</strong>
                </label>{" "}
                {this.renderUnitId(currentProduct.unit_id)}
              </div>
              <div>
                <label>
                  <strong>Curr Price:</strong>
                </label>{" "}
                {currentProduct.current_price}
              </div>
              <div>
                <label>
                  <strong>Created By :</strong>
                </label>{" "}
                {currentProduct.created_by}
              </div>

              <div>
                <label>
                  <strong>Created At :</strong>
                </label>{" "}
                {currentProduct.created_at}
              </div>

              <div>
                <label>
                  <strong>Last Update Time : </strong>
                </label>{" "}
                {currentProduct.last_update_time}
              </div>

             
              {/* <div>
                <label>
                  <strong>Status:</strong>
                </label>{" "}
                {currentProduct.published ? "Published" : "Pending"}
              </div> */}

              <Link
              to={"/super/products/" + currentProduct.id}
              //   to={{
              //     pathname: "/super/products/" + currentProduct.id,
              //     state: this.state.tokenValue
              // }}
                className="badge badge-warning"
              >
                Edit
              </Link>
            </div>
          ) : (
            <div className="clickon">
              <br />
              <p>Please click on a Product...</p>
            </div>
          )}
        </div>
      </div>
    </div>
    );
  }
}

export default withRouter(ProductsSuper)

