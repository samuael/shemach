import React, { Component } from "react";
import ProductService from '../../services/prodcutService'
import { Link } from "react-router-dom";
import './superControlAdmin.css'
import photo from '../../assets/18.jpg'; 


export default class SuperControlAdmin extends Component {
  constructor(props) {
    super(props);
    this.onChangesearchProduct = this.onChangesearchProduct.bind(this);
    this.retrieveProducts = this.retrieveProducts.bind(this);
    this.refreshList = this.refreshList.bind(this);
    this.setActiveProduct = this.setActiveProduct.bind(this);
    this.removeAllProducts = this.removeAllProducts.bind(this);
    this.searchProduct = this.searchProduct.bind(this);

    this.state = {
      products: [
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
    this.retrieveProducts();
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
          products: response.data
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

  searchProduct() {
    this.setState({
      currentProduct: null,
      currentIndex: -1
    });

    ProductService.findByTitle(this.state.searchProduct)
      .then(response => {
        this.setState({
          products: response.data
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
    const { searchProduct, products, currentProduct, currentIndex } = this.state;

    return (
    <div id="supercontroladmins">
      <div className="list row">
        <div className="col-md-8">
          <div className="input-group mb-3">
            <input
              type="text"
              className="form-control"
              placeholder="Search admins here"
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
              <Link  className="linkadd" to="/super-admin/reg-admin"><i class="fa-solid fa-plus"></i>Register Admin</Link>
              </button>
          </div>
        <div className="col-md-6 adminlistcontainer">
          <h4>Admin List</h4>

          <ul className="list-group">
            {products &&
              products.map((product, index) => (
            <div className={
                  "list-group-item " +
                  (index === currentIndex ? "active" : "")
                }
                onClick={() => this.setActiveProduct(product, index)}
                key={index}
                >    
                <div className="row eachadmins mt-4"> 
                    <div className="col-sm-4">
                        <img src={photo} alt="photo" className="adminimg"></img>
                    </div>   

                    <div className="col-sm-7">
                    <div className="Name">
                             Name :{product.title}
                        </div>
                    <p className="Place">Phone :{product.productionarea}</p>
                    <p className="Price">Email :{product.currentprice}</p>
                    </div> 

                    <div className="col-sm-1">
                    <Link  className="deleteadmin" to="/super-admin/add-product"><i class="fa-solid fa-trash-can"></i></Link>


                    </div>        
                        
                 </div>
             </div> 

              ))}
          </ul>

          <button
            className="m-3 btn btn-sm btn-danger"
            onClick={this.removeAllProducts}
          >
            Remove All
          </button>
        </div>
        <div className="col-md-6 description">
          {currentProduct ? (
            <div>
              <h4>Admin</h4>
              <div className="row">
                <div className="col-sm-6">
                            <img src={photo} alt="photo" className="adminimglarge"></img>
                </div> 
             <div className="col-sm-6">
                <div>
                    <label>
                    <strong>Name :</strong>
                    </label>{" "}
                    {currentProduct.title}
                </div>
                <div>
                    <label>
                    <strong>Gender :</strong>
                    </label>{" "}
                    {currentProduct.description}
                </div>
              
              <div>
                <label>
                  <strong>Email :</strong>
                </label>{" "}
                {currentProduct.productionarea}
              </div>
              <div>
                <label>
                  <strong>Phone :</strong>
                </label>{" "}
                {currentProduct.measurement}
              </div>
              <div>
                <label>
                  <strong>Age :</strong>
                </label>{" "}
                {currentProduct.prevprice}
              </div>
              </div>
              {/* <div>
                <label>
                  <strong>Curr Price</strong>
                </label>{" "}
                {currentProduct.currentprice}
              </div> */}
              {/* <div>
                <label>
                  <strong>Status:</strong>
                </label>{" "}
                {currentProduct.published ? "Published" : "Pending"}
              </div> */}

              <Link
                to={"/super/products/" + currentProduct.id}
                className="badge badge-warning col-sm-2 btn btn-primary"
              >
                Edit
              </Link>
              </div>
            </div>
          ) : (
            <div className="clickon">
              <br />
              <p>Please click on a Admin to view full profile...</p>
            </div>
          )}
        </div>
      </div>
    </div>
    );
  }
}


