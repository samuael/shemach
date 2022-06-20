import React, { Component } from "react";
import moment from 'moment'
import Moment from 'react-moment';
import { withRouter } from 'react-router-dom'
import ProductService from '../../services/prodcutService'
import { Link } from "react-router-dom";
import './products.css'

class Products extends Component {
  constructor(props) {
    super(props);
    this.onChangesearchProduct = this.onChangesearchProduct.bind(this);
    this.retrieveProducts = this.retrieveProducts.bind(this);
    this.refreshList = this.refreshList.bind(this);
    this.setActiveProduct = this.setActiveProduct.bind(this);
    this.removeAllProducts = this.removeAllProducts.bind(this);
    this.searchProduct = this.searchProduct.bind(this);
    this.renderUnitId = this.renderUnitId.bind(this);
    this.convertToDate = this.convertToDate.bind(this);
    this.getDifferenceInDays = this.getDifferenceInDays.bind(this);
   // this.tokenValue = this.props.location.state;
   
    this.state = {
      tokenValue: this.props.location.state,
      res_data: {
        status_code: 0,
        products: [],
        msg: ""
      },
      prodduct: [
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
    console.log(this.props);
    console.log(this.props.location)
    console.log(this.props.location.state);
    console.log(this.state.tokenValue);
     var dt = this.props.location.state;
     var tokenn = window.token;

    //  this.retrieveProducts();


     if(!tokenn){
      this.props.history.push('/')
     }else{
      this.retrieveProducts();
     }

   
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
          // res_data: JSON.parse(response.data)
          res_data: response.data
        }
        //console.log(response.data);
        );
        console.log(this.props.location.state);
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

  setActiveProduct(product, index) {
    this.setState({
      currentProduct: product,
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

convertToDate(number){
  var myDate = new Date( number *1000);
  var convertedDate = myDate.toLocaleString();
  return convertedDate;
  // return myDate.toGMTString()+ "<br>" + myDate.toLocaleString();
  // document.write(myDate.toGMTString()+"<br>"+myDate.toLocaleString());

}

getDifferenceInDays(date2) {
  var date1 =     new Date().toLocaleDateString();
  var date2 = this.convertToDate(date2);
  var date3 = new Date(date2).toLocaleDateString();
  // var date2new = this.convertToDate(date2).toLocaleDateString();
  // var dateonly = date2new.toLocaleDateString();
  console.log(date1);
  console.log(date3);

  var diffDays =  new Date().getTime() - new Date(date3).getTime() ;    //Current date - Past date  return in millisecond
  var diffDaysNew = Math.floor(diffDays / (1000 * 60 * 60 * 24));
  console.log(diffDaysNew);
  return diffDaysNew;
  // return diffInMs / (1000 * 60 * 60 * 24);

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
    const { searchProduct, products, prodduct, res_data, currentProduct, currentIndex } = this.state;
    console.log(prodduct);
    console.log(res_data);
    console.log(this.state)

    return (
      <div id="productsmain">
          <div className="list row">
              <div className="col-md-8">
              <div className="input-group mb-3">
                <input
                  type="text"
                  className="form-control"
                  placeholder="Search by product name"
                  value={searchProduct}
                  onChange={this.onChangesearchProduct}
                  data-cy="searchprod"
                />
                <div className="input-group-append">
                  <button
                    className="btn btn-outline-secondary searchbuttom"
                    type="button"
                    onClick={this.searchProduct}
                    data-cy = "onclicksearch"
                  >
                    Search
                  </button>
                </div>
              </div>
            </div>
             
          </div>
              <div className="row">
                <div className="col-sm-8 productlist">
                  <h4>Product List</h4>

                  {/* <ul className="list-group">
                    {products &&
                      products.map((product, index) => (
                        <div className="list-group-comp">                
                            <li
                            className={
                              "list-group-item " +
                              (index === currentIndex ? "active" : "")
                            }
                            onClick={() => this.setActiveProduct(product, index)}
                            key={index}
                          >
                            {product.title}
                          </li>
                        </div>

                      ))}
                  </ul> */}
                      <ul className="list-group lstofitems">
                    { 
                      res_data.products.map((product, index) => (
                        <li className={
                          "list-group-item " +
                          (index === currentIndex ? "active" : "")
                        }
                        onClick={() => this.setActiveProduct(product, index)}
                        key={index}
                        >                
                            <div className="Name">
                              
                            Crop :{product.name}
                            </div>
                          
                          <p className="Time">Price :{product.current_price}</p>
                          <p className="Time">Before {this.getDifferenceInDays(product.last_update_time)} days</p>
                        </li>

                      ))

                    //  data = Array.from(products.data);
                    
                    //  res_data.products.map((product, index) => (
                    //     <li
                    //       className={
                    //         "list-group-item " +
                    //         (index === currentIndex ? "active" : "")
                    //       }
                    //       onClick={() => this.setActiveProduct(product, index)}
                    //       key={index}
                    //     >
                    //       {product.name}
                    //     </li>
                    //   ))


                      }
                    </ul>

                  {/* <button
                    className="m-3 btn btn-sm btn-danger"
                    onClick={this.removeAllProducts}
                  >
                    Remove All
                  </button> */}
                </div>
                <div className="col-sm-4 description ">
                  {currentProduct ? (
                    <div>
                      <h4>Product</h4>
                      <div>
                        <label>
                          <strong>Name:</strong>
                        </label>{" "}
                        {currentProduct.name}
                      </div>
                      <div>
                        <label>
                          <strong>Prod Area:</strong>
                        </label>{" "}
                        {currentProduct.production_area}
                      </div>
                      <div>
                        <label>
                          <strong>Unit ID :</strong>
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
                          <strong>Created By : </strong>
                        </label>{" "}
                        {currentProduct.created_by}
                      </div>

                      <div>
                        <label>
                          <strong>Created At : </strong>
                        </label>{" "}
                        {this.convertToDate(currentProduct.created_at)}
                      </div>

                      <div>
                        <label>
                          <strong>Last Update Time : </strong>
                        </label>{" "}
                        {this.convertToDate(currentProduct.last_update_time)}
                      </div>

                      {/* <div>
                        <label>
                          <strong>Status:</strong>
                        </label>{" "}
                        {currentProduct.published ? "Published" : "Pending"}
                      </div> */}

                      {/* <Link
                        to={"/info/product/" + currentProduct.id}
                        state={this.state.tokenValue}
                        className="badge badge-warning"
                      >
                        Edit
                      </Link> */}
                        <Link
                        to={"/info/product/" + currentProduct.id}
                        // state={this.state.tokenValue}
                      //   to={{
                      //     pathname: "/info/product/" + currentProduct.id,
                      //     state: this.state.tokenValue
                      // }}
                        className="editbutton badge badge-warning"
                        data-cy="editpricebutton"
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

export default withRouter(Products)

