import React, { Component } from "react";
import MessageService from "../../services/prodcutService";
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
    this.onChangeSelectedProduct = this.onChangeSelectedProduct.bind(this);
    this.onChangeShowToAdmins = this.onChangeShowToAdmins.bind(this);
    this.onChangeLang = this.onChangeLang.bind(this);
    this.onChangeData = this.onChangeData.bind(this);
    this.connect = this.connect.bind(this);
    this.sendMessage = this.sendMessage.bind(this);
    this.convertToDate = this.convertToDate.bind(this);
    this.deleteMessage = this.deleteMessage.bind(this);

// 
    this.state = {
      status_code: 0,
      msg: "",
      msgres:"",
      targets:[],
      message: [],
      // targets: [],
      // lang: "",
      // data: "",
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

    this.retrieveMessage();
    // this.connect();


  }

  componentWillUnmount() {
    // console.log("closing websocket is closed...");
    // this.ws.close();
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

  // onChangeSelectedProduct(e) {
  //   this.setState({
  //     targets: e.target.value,
  //   });
  // }

  onChangeSelectedProduct(event) {
  let value = Array.from(
    event.target.selectedOptions,
    (option) => parseInt(option.value)
    // console.log(option.value)
  );
  this.setState({
    targets: value,
  });
}

  onChangeSearchMessage(e) {
    const searchMessage = e.target.value;

    this.setState({
      searchMessage: searchMessage,
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

  retrieveMessage() {
    var token = window.token;
    MessageService.superGetAllMessage(token)
      .then(response => {
        this.setState({
          status_code: response.data.status_code,
          msg: response.data.msg,
          message: response.data.message,
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

  sendMessage(event) {
    event.preventDefault();
    var data = {
     // unit_id: parseInt(this.state.unit_id),
      // targets: [-1,2],
      targets: this.state.targets,
      lang: this.state.lang,
      data:this.state.data
    };

    // var token = this.state.tkValue;
    var token = window.token;
    console.log(data);
    console.log(token);

    MessageService.superSendMessage(data, token)
      .then(response => {
        console.log(response.data)
        this.setState({

       // response_data: {parseInt(
          status_code: response.data.status_code,
          msg: response.data.msg,
          // message: response.data.message,
          msgres:"sent",
          submitted: true
        
        });
        console.log(response.data);
      //  console.log(response_data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  convertToDate(number){
    var myDate = new Date( number *1000);
    var convertedDate = myDate.toLocaleString();
    return convertedDate;
    // return myDate.toGMTString()+ "<br>" + myDate.toLocaleString();
    // document.write(myDate.toGMTString()+"<br>"+myDate.toLocaleString());
  
  }


  searchMessage() {
    this.setState({
      currentMessage: null,
      currentIndex: -1,
    });

    MessageService.findByTitle(this.state.searchMessage)
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

  deleteMessage() {   
    var token  = window.token
    console.log(token); 
    MessageService.deleteMessageSuper(this.state.currentMessage.id, token)
      .then(response => {
        this.setState({
          status_code: response.data.status_code,
          msg: response.data.msg
        });
      //  console.log(response.data);
        console.log(response.data);
        console.log(this.props.history);
        // this.props.history.push('/super-admin/broadcast/received')
      })
      .catch(e => {
        console.log(e);
      });
  }


  render() {
    const { searchMessage, message, msgres, currentMessage, currentIndex } =
      this.state;
    // console.log(ws);
    return (
      <>
        <div id="brodreceived">

          <div className="row">
            <div className="col-sm-4 mt-4">
              <div className="row">
                <h4 className="mt-4 col-sm-4">Inboxes</h4>
                <a className="col-sm-4 mt-4 refreshlink" onClick={this.retrieveMessage}>Refresh</a>
              </div>


              <ul className="list-group ms-4">
              { 
                      message.map((message, index) => (
                        <li className={
                          "list-group-item " +
                          (index === currentIndex ? "active" : "")
                        }
                        onClick={() => this.setActiveMessage(message, index)}
                        key={index}
                        >  
                        <div className="row">              
                            <div className="Name col-sm-10">
                              
                            {message.data}
                            </div>
                          
                          {/* <p className="Time">Time :{message.created_at}</p> */}
                          {/* <p className="Time">Before {this.getDifferenceInDays(message.last_update_time)} days</p> */}

                          <div className="col-sm-2">
                             <button  className="deletemessage" onClick={this.deleteMessage}><i class="fa-solid fa-trash-can"></i></button>
                         </div> 
                         </div>
                        </li>

                      ))

                    //  data = Array.from(products.data);
                    
                    //  res_data.products.map((product, index) => (
                    //     <li
                    //       className={
                    //         "list-group-item " +
                    //         (index === currentIndex ? "active" : "")
                    //       }
                    //       onClick={() => this.setActiveMessage(product, index)}
                    //       key={index}
                    //     >
                    //       {product.name}
                    //     </li>
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
                      <strong>From :</strong>
                    </label>{" "}
                    {currentMessage.id}
                  </div>
                  <div>
                    <label>
                      <strong>Time :</strong>
                    </label>{" "}
                    {this.convertToDate(currentMessage.created_at)}
                  </div>
                  <div>
                    <label>
                      <strong>Description :</strong>
                    </label>{" "}
                    {currentMessage.data}
                  </div>
                  <div>
                    <label>
                      <strong>Language :</strong>
                    </label>{" "}
                    {currentMessage.lang}
                  </div>
{/* 
                  <div> 
                      <button
                      className="badge badge-danger col-sm-2 btn btn-primary"
                      onClick={this.deleteMessage}
                    >
                      Delete
                    </button>
                  </div> */}
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
              {/* <div id="log">{this.state.ws}</div> */}
              <div id="log"></div>
              <form id="form">
              {/* <input type="submit" value="Send" />
              <input type="text" id="msg" size="64" autofocus />    */}
                <div className="">
                  {/* <div className="form-group">
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
                  </div> */}
                  <div className="form-group mt-3">
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
                    <label data-cy="showachooseprod" htmlFor="prod">Product :</label>

                    <select
                      name="prod"
                      id="prod"
                      required
                      // className="form-control"
                      value={this.state.targets}
                      onChange={this.onChangeSelectedProduct}
                      multiple={true}
                      type="number"
                    >
                      <option value="0">Select Crop</option>
                      <option value="-1">All</option>
                      <option value="1">Teff</option>
                      <option value="2">Wheat</option>
                      <option value="3">Sorghum</option>
                      <option value="4">Maize</option>
                      <option value="5">Bean</option>
                      <option value="6">Ater</option>
                    </select>
                  </div>

                  <div className="form-group">
                    <label data-cy="showchooselang" htmlFor="lang">Language :</label>

                    <select
                      name="lang"
                      id="lang"
                      required
                      // className="form-control"
                      value={this.state.lang}
                      onChange={this.onChangeLang}
                      type="text"
                    >
                      <option value="sel">Select Language</option>
                      <option value="all">All</option>
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
                         value={this.state.data}
                        data-cy="showinfo"
                        placeholder="Information about the new product"
                        // onChange={(e) => setMessage(e.target.value)}
                      />
                    </label>
                  </div>
                  <button
                    className="btn btn-primary mb-3"
                    type="submit"
                    value="Send"
                    data-cy="sendmessagebtn"
                    onClick={this.sendMessage} 
                  >
                    Send
                  </button>
                  <p>{msgres}</p>
                </div>
              </form>
            </div>
          </div>
        </div>
      </>
    );
  }
}
