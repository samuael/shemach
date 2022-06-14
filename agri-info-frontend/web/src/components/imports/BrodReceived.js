import React, { Component } from "react";
import ProductService from "../../services/prodcutService";
import { Link } from "react-router-dom";
import useWebSocket, { ReadyState } from 'react-use-websocket';
// import WebSocket, { WebSocketServer } from 'ws';
// import WebSocket from 'websocket';



import "./brodReceived.css";

export default class BrodReceived extends Component {
  constructor(props) {
    super(props);
    this.onChangeSearchMessage = this.onChangeSearchMessage.bind(this);
    this.retrieveMessage = this.retrieveMessage.bind(this);
    this.refreshList = this.refreshList.bind(this);
    this.setActiveMessage = this.setActiveMessage.bind(this);
    this.removeAllMessage = this.removeAllMessage.bind(this);
    this.searchMessage = this.searchMessage.bind(this);
    this.onChangeSelectedProduct = this.onChangeSelectedProduct.bind(this);
    this.onChangeShowToAdmins = this.onChangeShowToAdmins.bind(this);
    this.onChangeLang = this.onChangeLang.bind(this);
    this.onChangeData = this.onChangeData.bind(this);
    this.connect = this.connect.bind(this);
// 
    this.state = {
      wsoc: null,
      messages: [],
      messagebody: {  
        targets: [],
        lang: "",
        data: "",
        // created_at: 0
      },
      type: 0,
      currentMessage: null,
      currentIndex: -1,
      searchMessage: "",
    };
  }

  componentDidMount() {
    // var conn;
    // var msg = document.getElementById("msg");
    // var log = document.getElementById("log");
  
    // function appendLog(item) {
    //     var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
    //     log.appendChild(item);
    //     if (doScroll) {
    //         log.scrollTop = log.scrollHeight - log.clientHeight;
    //     }
    // }

    // document.getElementById("form").onsubmit = function () {
    //     if (!conn) {
    //         return false;
    //     }
    //     if (!msg.value) {
    //         return false;
    //     }
    //     conn.send(msg.value);
    //     msg.value = "";
    //     return false;
    // };

    // if (window["WebSocket"]) {
    //     var id  = window.id;
    //     conn = new WebSocket(`ws://localhost:8080/api/connection/admins/${id}`);
    //     // const conn = new WebSocket("wss://ws.bitstamp.net");
    //     conn.onclose = function (evt) {
    //         var item = document.createElement("div");
    //         item.innerHTML = "<b>Connection closed.</b>";
    //         appendLog(item);
    //     };
    //     conn.onmessage = function (evt) {
    //         var messages = evt.data.split('\n');
    //         for (var i = 0; i < messages.length; i++) {
    //             var item = document.createElement("div");
    //             item.innerText = messages[i];
    //             appendLog(item);
    //         }
    //     };
    // } else {
    //     var item = document.createElement("div");
    //     item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
    //     appendLog(item);
    // }

    // this.retrieveMessage();
    this.connect();
  }

  

  // componentDidMount(){
  //   // this is an "echo" websocket service
  // //   // this.connection = new WebSocket('ws://localhost:3000/api/connection/admins');
  //   this.connection = new WebSocket('wss://echo.websocket.org');

  //   // listen to onmessage event
  //   this.connection.onmessage = evt => {
  //     // add the new message to state
  //       this.setState({
  //       messages : this.state.messages.concat([ evt.data ])
  //     })
  //   };

  //   // for testing purposes: sending to the echo service which will send it back back
  //   // setInterval( _ =>{
  //   //     this.connection.send( Math.random() )
  //   // }, 2000 )
  // }

  timeout = 250; // Initial timeout duration as a class variable

  connect = () => {
    var token = window.token;
    var id = window.id;
    console.log(token);
    var ws = new WebSocket(`ws://localhost:8080/api/connection/admins/${id}`);

    // {
    //   "Authorization": "Bearer" + token
    // }

    console.log(ws);

    let that = this; // cache the this
    var connectInterval;

    // websocket onopen event listener
    ws.onopen = () => {
      console.log("connected websocket main component");

      this.setState({ wsoc: ws });
      ws.send("Message from the Server")

      that.timeout = 250; // reset timer to 250 on open of websocket connection
    //  clearTimeout(connectInterval); // clear Interval on on open of websocket connection
    };

    // websocket onclose event listener
    ws.onclose = (e) => {
      console.log(
        `Socket is closed. Reconnect will be attempted in ${Math.min(
          10000 / 1000,
          (that.timeout + that.timeout) / 1000
        )} second.`,
        e.reason
      );

      that.timeout = that.timeout + that.timeout; //increment retry interval
    //  connectInterval = setTimeout(this.check, Math.min(10000, that.timeout)); //call check function after timeout
    };

    // websocket onerror event listener
    ws.onerror = (err) => {
      console.error(
        "Socket encountered error: ",
        err.message,
        "Closing socket"
      );

      ws.close();
    };
  };

  check = () => {
    const { ws } = this.state;
    if (!ws || ws.readyState == WebSocket.CLOSED) this.connect(); //check if websocket instance is closed, if so call `connect` function.
  };

  onChangeData(e) {
    this.setState({
      data: e.target.value,
    });
  }

  onChangeLang(e) {
    this.setState({
      lang: e.target.value,
    });
  }

  onChangeShowToAdmins(e) {
    this.setState({
      targets: e.target.checked,
    });
  }

  onChangeSelectedProduct(e) {
    this.setState({
      targets: e.target.value,
    });
  }

  onChangeSearchMessage(e) {
    const searchMessage = e.target.value;

    this.setState({
      searchMessage: searchMessage,
    });
  }

  retrieveMessage() {
    ProductService.getAll()
      .then((response) => {
        this.setState({
          products: response.data,
        });
        console.log(response.data);
      })
      .catch((e) => {
        console.log(e);
      });
  }

  refreshList() {
    this.retrieveMessage();
    this.setState({
      currentMessage: null,
      currentIndex: -1,
    });
  }

  setActiveMessage(tutorial, index) {
    this.setState({
      currentMessage: tutorial,
      currentIndex: index,
    });
  }

  removeAllMessage() {
    ProductService.deleteAll()
      .then((response) => {
        console.log(response.data);
        this.refreshList();
      })
      .catch((e) => {
        console.log(e);
      });
  }

  searchMessage() {
    this.setState({
      currentMessage: null,
      currentIndex: -1,
    });

    ProductService.findByTitle(this.state.searchMessage)
      .then((response) => {
        this.setState({
          products: response.data,
        });
        console.log(response.data);
      })
      .catch((e) => {
        console.log(e);
      });
  }

  render() {
    const { searchMessage, ws, products, currentMessage, currentIndex } =
      this.state;
    console.log(ws);
    return (
      <>
        {/* <ul>{ this.state.messages.slice(-5).map( (msg, idx) => <li key={'msg-' + idx }>{ msg }</li> )}</ul>; */}
        <div>{ws}</div>
        <div id="brodreceived">
          <div className="list row">
            <div className="col-md-8">
              <div className="input-group mb-3">
                <input
                  type="text"
                  className="form-control"
                  placeholder="Search messages here ..."
                  value={searchMessage}
                  onChange={this.onChangeSearchMessage}
                />
                <div className="input-group-append mt-4">
                  <button
                    className="btn btn-outline-secondary searchbuttom"
                    type="button"
                    onClick={this.searchMessage}
                  >
                    Search
                  </button>
                </div>
              </div>
            </div>
          </div>

          <div className="row">
            <div className="col-sm-4 mt-4">
              <h4>Inboxes</h4>

              <ul className="list-group">
                {
                  // products &&
                  //   products.map((product, index) => (
                  //     <div className={
                  //       "list-group-item " +
                  //       (index === currentIndex ? "active" : "")
                  //     }
                  //     onClick={() => this.setActiveMessage(product, index)}
                  //     key={index}
                  //     >
                  //         <div className="Name">
                  //         {product.title}
                  //         </div>
                  //       <p className="Time">{product.updatedAt}</p>
                  //       <p className="message">Eu sit labore minim adipisicing eu Lorem et fugiat non magna cupidatat. Ad consequat nisi qui aute consectetur nulla nulla nisi duis mollit elit laboris nostrud. Tempor sit commodo tempor pariatur excepteur sint culpa dolore. Nisi occaecat est amet deserunt. Anim dolore proident aute ullamco commodo non. </p>
                  //     </div>
                  //   ))
                }
              </ul>

              {/* <button
            className="m-3 btn btn-sm btn-danger"
            onClick={this.removeAllMessage}
          >
            Remove All
          </button> */}
            </div>
            <div className="col-sm-4 description">
              {currentMessage ? (
                <div>
                  <h4>Message</h4>
                  <div>
                    <label>
                      <strong>From:</strong>
                    </label>{" "}
                    {currentMessage.title}
                  </div>
                  <div>
                    <label>
                      <strong>Time:</strong>
                    </label>{" "}
                    {currentMessage.updatedAt}
                  </div>
                  <div>
                    <label>
                      <strong>Description:</strong>
                    </label>{" "}
                    Eu sit labore minim adipisicing eu Lorem et fugiat non magna
                    cupidatat. Ad consequat nisi qui aute consectetur nulla
                    nulla nisi duis mollit elit laboris nostrud. Tempor sit
                    commodo tempor pariatur excepteur sint culpa dolore. Nisi
                    occaecat est amet deserunt. Anim dolore proident aute
                    ullamco commodo non. {currentMessage.description}
                  </div>

                  {/* <div>
                <label>
                  <strong>Description:</strong>
                </label>{" "}
                {currentMessage.description}
              </div>
              <div>
                <label>
                  <strong>Measurement:</strong>
                </label>{" "}
                {currentMessage.measurement}
              </div>
              <div>
                <label>
                  <strong>Prev Price:</strong>
                </label>{" "}
                {currentMessage.prevprice}
              </div>
              <div>
                <label>
                  <strong>Curr Price</strong>
                </label>{" "}
                {currentMessage.currentprice}
              </div>
              <div>
                <label>
                  <strong>Status:</strong>
                </label>{" "}
                {currentMessage.published ? "Published" : "Pending"}
              </div>

              <Link
                to={"/tutorials/" + currentMessage.id}
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

            <div className="col-sm-2 createform">
              {/* <h4>Inboxes</h4> */}
              <div id="log">{this.state.ws}</div>
              <form id="form">
                <div className="">
                  <div className="form-group">
                    <label data-cy="showadminlabel" htmlFor="name" className="chkbox">
                      Show to Admins :
                      <input
                        className="chkboxiput"
                        type="checkbox"
                        id="name"
                        name="name"
                        value={this.state.targets}
                        onChange={this.onChangeShowToAdmins}
                      />
                    </label>
                  </div>
                  <div className="form-group">
                    {/* <label htmlFor="email">
                      Your Email
                      <input
                        type="email"
                        id="email"
                        name="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                      />
                    </label> */}
                    <label data-cy="showachooseprod" htmlFor="prod">Choose a Product :</label>

                    <select
                      name="prod"
                      id="prod"
                      required
                      // className="form-control"
                      value={this.state.targets}
                      onChange={this.onChangeSelectedProduct}
                      type="number"
                    >
                      <option value="1">Teff</option>
                      <option value="2">Nashila</option>
                      <option value="3">Wheat</option>
                      <option value="4">Sorghum</option>
                    </select>
                  </div>

                  {/* <div className="row">
                    <div className="col-sm-1">
                        <div className="form-check">
                        <input
                          className="form-check-input"
                          type="checkbox"
                          name="languages"
                          value="Javascript"
                          id="flexCheckDefault"
                        // onChange={handleChange}
                        />
                        <label
                          className="form-check-label"
                          htmlFor="flexCheckDefault"
                        >
                            Javascript
                        </label>
                      </div>
                      <div className="form-check">
                        <input
                          className="form-check-input"
                          type="checkbox"
                          name="languages"
                          value="Javascript"
                          id="flexCheckDefault"
                        // onChange={handleChange}
                        />
                        <label
                          className="form-check-label"
                          htmlFor="flexCheckDefault"
                        >
                            Javascript
                        </label>
                      </div>
                      <div className="form-check">
                        <input
                          className="form-check-input"
                          type="checkbox"
                          name="languages"
                          value="Javascript"
                          id="flexCheckDefault"
                        // onChange={handleChange}
                        />
                        <label
                          className="form-check-label"
                          htmlFor="flexCheckDefault"
                        >
                            Javascript
                        </label>
                      </div>

                    </div>


                    <div className="col-sm-1">
                        <div className="form-check">
                        <input
                          className="form-check-input"
                          type="checkbox"
                          name="languages"
                          value="Javascript"
                          id="flexCheckDefault"
                        // onChange={handleChange}
                        />
                        <label
                          className="form-check-label"
                          htmlFor="flexCheckDefault"
                        >
                            Javascript
                        </label>
                      </div>
                      <div className="form-check">
                        <input
                          className="form-check-input"
                          type="checkbox"
                          name="languages"
                          value="Javascript"
                          id="flexCheckDefault"
                        // onChange={handleChange}
                        />
                        <label
                          className="form-check-label"
                          htmlFor="flexCheckDefault"
                        >
                            Javascript
                        </label>
                      </div>
                      <div className="form-check">
                        <input
                          className="form-check-input"
                          type="checkbox"
                          name="languages"
                          value="Javascript"
                          id="flexCheckDefault"
                        // onChange={handleChange}
                        />
                        <label
                          className="form-check-label"
                          htmlFor="flexCheckDefault"
                        >
                            Javascript
                        </label>
                      </div>

                    </div>
                  </div> */}

                  <div className="form-group">
                    <label data-cy="showchooselang" htmlFor="lang">Choose a Language :</label>

                    <select
                      name="lang"
                      id="lang"
                      required
                      // className="form-control"
                      value={this.state.lang}
                      onChange={this.onChangeLang}
                      type="text"
                    >
                      <option value="all">All</option>
                      <option value="amh">Amharic</option>
                      <option value="oro">Oromic</option>
                      <option value="som">Somalic</option>
                      <option value="sid">Sidex</option>
                    </select>
                  </div>
                </div>

                <div className="col-sm-4">
                  <div className="form-group">
                    <label htmlFor="msg" className="infocreate">
                      Information :
                      <textarea
                        type="text"
                        id="msg"
                        name="msg"
                        onChange={this.onChangeData}
                        //  value={message}
                        data-cy="showinfo"
                        placeholder="Information about the new product"
                        // onChange={(e) => setMessage(e.target.value)}
                      />
                    </label>
                  </div>
                  <button
                    className="btn btn-primary"
                    type="submit"
                    value="Send"
                    data-cy="sendmessagebtn"
                  >
                    Send
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </>
    );
  }
}
