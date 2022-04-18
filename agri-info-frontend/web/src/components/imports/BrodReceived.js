import React, { Component } from "react";
import ProductService from '../../services/prodcutService'
import { Link } from "react-router-dom";
import './brodReceived.css'

export default class BrodReceived extends Component {
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
    //  //   id: 1,
    //     title: 'Product 1',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //  //   id: 2,
    //     title: 'Product 2',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //  //   id: 3,
    //     title: 'Product 3',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //  //   id: 4,
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
       <div id="brodreceived">
      <div className="list row">
        <div className="col-md-8">
          <div className="input-group mb-3">
            <input
              type="text"
              className="form-control"
              placeholder="Search messages here ..."
              value={searchProduct}
              onChange={this.onChangesearchProduct}
            />
            <div className="input-group-append mt-4">
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
        <div className="col-md-6 mt-4">
          <h4>Inboxes</h4>

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
                    <div className="Name">
                      
                    {product.title}
                    </div>
                  <p className="Time">{product.updatedAt}</p>
                  <p className="message">Eu sit labore minim adipisicing eu Lorem et fugiat non magna cupidatat. Ad consequat nisi qui aute consectetur nulla nulla nisi duis mollit elit laboris nostrud. Tempor sit commodo tempor pariatur excepteur sint culpa dolore. Nisi occaecat est amet deserunt. Anim dolore proident aute ullamco commodo non. </p>
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
              <h4>Message</h4>
              <div>
                <label>
                  <strong>From:</strong>
                </label>{" "}
                {currentProduct.title}
              </div>
              <div>
                <label>
                  <strong>Time:</strong>
                </label>{" "}
                {currentProduct.updatedAt}
              </div>
              <div>
                <label>
                  <strong>Description:</strong>
                </label>{" "}
                Eu sit labore minim adipisicing eu Lorem et fugiat non magna cupidatat. Ad consequat nisi qui aute consectetur nulla nulla nisi duis mollit elit laboris nostrud. Tempor sit commodo tempor pariatur excepteur sint culpa dolore. Nisi occaecat est amet deserunt. Anim dolore proident aute ullamco commodo non.                {currentProduct.description}
              </div>
         
              {/* <div>
                <label>
                  <strong>Description:</strong>
                </label>{" "}
                {currentProduct.description}
              </div>
              <div>
                <label>
                  <strong>Measurement:</strong>
                </label>{" "}
                {currentProduct.measurement}
              </div>
              <div>
                <label>
                  <strong>Prev Price:</strong>
                </label>{" "}
                {currentProduct.prevprice}
              </div>
              <div>
                <label>
                  <strong>Curr Price</strong>
                </label>{" "}
                {currentProduct.currentprice}
              </div>
              <div>
                <label>
                  <strong>Status:</strong>
                </label>{" "}
                {currentProduct.published ? "Published" : "Pending"}
              </div>

              <Link
                to={"/tutorials/" + currentProduct.id}
                className="badge badge-warning"
              >
                Edit
              </Link> */}
            </div>
          ) : (
            <div className="clickon">
              <br />
              <p>Please click on message to read ...</p>
            </div>
          )}
        </div>
      </div>
      </div>
    
    );
  }
}


