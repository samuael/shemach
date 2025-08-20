import './headerInfo.css'
import {Switch, Route, Link} from 'react-router-dom'

import React, { Component } from 'react'

class HeaderInfo extends Component {
    constructor(props) {
        super(props);
        this.onLogout = this.onLogout.bind(this);

    }

    onLogout(){
        window.token= "";
    }


    render() {
        return (
            <div>
                <header id="navitem">
                <nav class="navbar navbar-expand-lg">
                <div class="container">
                {/* <Link to='/' className='navbar-logoo'><i class="fa-solid fa-list"></i>  shemach <i class="fa-solid fa-building-wheat"></i> </Link> */}
                <Link class="navbar-brand text-white" to='/products' ><i class="fa-solid fa-list"></i> Agri Net </Link>
                {/* <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#nvbCollapse" aria-controls="nvbCollapse">
                    <span class="navbar-toggler-icon"></span>
                </button> */}
                <div class="collapse navbar-collapse" id="nvbCollapse">
                    <ul class="navbar-nav ml-auto nav-menu">
                        {/* <li class="nav-item active pl-1">
                            <Link class="nav-link" to="/marketing"><i class="fa-brands fa-product-hunt"></i>Products <i class="fa-solid fa-circle-dollar"></i></Link>
                        </li> */}
                        <li className="nav-item pl-1">
                            <Link class="nav-link" to="/products"><i class="fa-solid fa-hand-holding-dollar"></i>Products & Price <i class="fa-solid fa-circle-dollar"></i></Link>
                        </li>
                        <li className="nav-item pl-1 ms-4">
                            <Link className="nav-link" to="/info-admin/broadcast/received"><i class="fa-solid fa-message"></i>Messages </Link>
                            {/* <span class="circle">5</span> */}
                        </li>
                        <li className="nav-item pl-1 logout">
                            <Link className="nav-link" to="/" onClick={this.onLogout}><i className="fa fa-sign-out me-1" aria-hidden="true"></i>Logout </Link>
                        </li>
                    </ul>
                </div>
                </div>
                </nav>
      	</header>
            </div>
        )
    }
}



export default HeaderInfo
