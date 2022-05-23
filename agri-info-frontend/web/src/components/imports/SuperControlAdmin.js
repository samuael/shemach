import React, { Component } from "react";
import adminService from '../../services/prodcutService'
import { Link } from "react-router-dom";
import './superControlAdmin.css'
import photo from '../../assets/18.jpg'; 


export default class SuperControlAdmin extends Component {
  constructor(props) {
    super(props);
    this.onChangeSearchAdmin = this.onChangeSearchAdmin.bind(this);
    this.retrieveAdmins = this.retreveAdmins.bind(this);
    this.refreshList = this.refreshList.bind(this);
    this.setActiveAdmin = this.setActiveAdmin.bind(this);
    this.removeAllAdmins = this.removeAllAdmins.bind(this);
    this.searchAdmin = this.searchAdmin.bind(this);
    this.deleteAdmin = this.deleteAdmin.bind(this);

    this.state = {
     
        admins: [],
      prodducts: [
    //     {
    //     id: 1,
    //     title: 'admin 1',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //     id: 2,
    //     title: 'admin 2',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //     id: 3,
    //     title: 'admin 3',
    //     description: 'Create tasks',
    //     published : false,
    // },
    // {
    //     id: 4,
    //     title: 'admin 4',
    //     description: 'Create tasks',
    //     published : false,
    // },

      ],
      currentAdmin: null,
      currentIndex: -1,
      searchAdmin: ""
    };
  }

  componentDidMount() {
    console.log(this.props)
    this.retreveAdmins();
  }

  onChangeSearchAdmin(e) {
    const searchAdmin = e.target.value;

    this.setState({
      searchAdmin: searchAdmin
    });
  }

  retreveAdmins() {
    adminService.getAllAdmins()
      .then(response => {
        this.setState({
          admins: response.data
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  refreshList() {
    this.retreveAdmins();
    this.setState({
      currentAdmin: null,
      currentIndex: -1
    });
  }

  setActiveAdmin(admin, index) {
    this.setState({
      currentAdmin: admin,
      currentIndex: index
    });
  }

  removeAllAdmins() {
    adminService.deleteAll()
      .then(response => {
        console.log(response.data);
        this.refreshList();
      })
      .catch(e => {
        console.log(e);
      });
  }

  searchAdmin() {
    this.setState({
      currentAdmin: null,
      currentIndex: -1
    });

    adminService.findAdminByName(this.state.searchAdmin)
      .then(response => {
        this.setState({
          admins: response.data
        });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  deleteAdmin() {   
    var token  = window.token
    console.log(token); 
    adminService.deleteInfoAdmin(this.state.currentAdmin.id, token)
      .then(response => {
        console.log(response.data);
        this.props.history.push('/super-admin/control-admins')
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
    const { searchAdmin, admins, currentAdmin, currentIndex } = this.state;
    console.log(admins);
 //  console.log(res_data);

    return (
    <div id="supercontroladmins">
      <div className="list row">
        <div className="col-md-10">
          {/* <div className="input-group mb-3">
            <input
              type="text"
              className="form-control"
              placeholder="Search admins here"
              value={searchAdmin}
              onChange={this.onChangeSearchAdmin}
            />
            <div className="input-group-append">
              <button
                className="btn btn-outline-secondary searchbuttom"
                type="button"
                onClick={this.searchAdmin}
              >
                Search
              </button>
            </div>
          </div> */}
       
        </div> 
        <div className="col-md-2">
              <button className="btn btn-primary add-button">
              <Link  className="linkadd" to="/super-admin/reg-admin"><i class="fa-solid fa-plus"></i>Register Admin</Link>
              </button>
          </div>
        <div className="col-md-6 adminlistcontainer">
          <h4>Admin List</h4>

          <ul className="list-group">
            {
            admins.map((admin, index) => (
            <div className={
                  "list-group-item " +
                  (index === currentIndex ? "active" : "")
                }
                onClick={() => this.setActiveAdmin(admin, index)}
                key={index}
                >    
                <div className="row eachadmins mt-4"> 
                    <div className="col-sm-4">
                        <img src={admin.imgurl} alt="photo" className="adminimg"></img>
                    </div>   

                    <div className="col-sm-7">
                    <div className="Name">
                             Name :{admin.firstname}
                        </div>
                    <p className="Place">Phone :{admin.phone}</p>
                    <p className="Price">Email :{admin.email}</p>
                    </div> 

                    <div className="col-sm-1">
                    <button  className="deleteadmin" onClick={this.deleteAdmin}><i class="fa-solid fa-trash-can"></i></button>


                    </div>        
                        
                 </div>
             </div> 

              ))}
          </ul>

          {/* <button
            className="m-3 btn btn-sm btn-danger"
            onClick={this.removeAllAdmins}
          >
            Remove All
          </button> */}
        </div>
        <div className="col-md-6 description">
          {currentAdmin ? (
            <div>
              <h4>Admin</h4>
              <div className="row">
                <div className="col-sm-6">
                            <img src={admins.imgurl} alt="photo" className="adminimglarge"></img>
                </div> 
             <div className="col-sm-6">
                <div>
                    <label>
                    <strong>FirstName :</strong>
                    </label>{" "}
                    {currentAdmin.firstname}
                </div>
                <div>
                    <label>
                    <strong>Lastname :</strong>
                    </label>{" "}
                    {currentAdmin.lastname}
                </div>
              
              <div>
                <label>
                  <strong>Email :</strong>
                </label>{" "}
                {currentAdmin.email}
              </div>
              <div>
                <label>
                  <strong>Phone :</strong>
                </label>{" "}
                {currentAdmin.phone}
              </div>
              <div>
                <label>
                  <strong>Created By :</strong>
                </label>{" "}
                {currentAdmin.Createdby}
              </div>

              <div>
                <label>
                  <strong>Created At :</strong>
                </label>{" "}
                {currentAdmin.created_at}
              </div>

              <div>
                <label>
                  <strong>Language :</strong>
                </label>{" "}
                {currentAdmin.lang}
              </div>


              </div>
              {/* <div>
                <label>
                  <strong>Curr Price</strong>
                </label>{" "}
                {currentAdmin.currentprice}
              </div> */}
              {/* <div>
                <label>
                  <strong>Status:</strong>
                </label>{" "}
                {currentAdmin.published ? "Published" : "Pending"}
              </div> */}

              {/* <Link
                to={"/super-admin/admins/" + currentAdmin.id}
                className="badge badge-warning col-sm-2 btn btn-primary"
              >
                Edit
              </Link> */}

              <button
              className="badge badge-danger col-sm-2 btn btn-primary"
              onClick={this.deleteAdmin}
            >
              Delete
            </button>

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


