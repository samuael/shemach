import React, { Component } from 'react'
import { Link } from "react-router-dom";
import './headerSuper.css'


class HeaderSuper extends Component {

    constructor(props) {
        super(props);
        this.onLogout = this.onLogout.bind(this);

    }

    onLogout(){
        window.token= "";
    }


    render() {
        return (
            <header id="navitemsuper">
            <nav className="navbar navbar-expand-lg">
            <div className="container">
            {/* <Link to='/' classNameName='navbar-logoo'><i className="fa-solid fa-list"></i>  Agri-Net <i className="fa-solid fa-building-wheat"></i> </Link> */}
            <Link className="navbar-brand text-white" to='/'><i className="fa-solid fa-list"></i> Agri Net </Link>
            <div className="collapse navbar-collapse" id="nvbCollapse">
                <ul className="navbar-nav ml-auto nav-menu">
                    <li className="nav-item  pl-1">
                        <Link className="nav-link" to="/super-admin/products"><i className="fa-brands fa-product-hunt"></i>Products <i className="fa-solid fa-circle-dollar"></i></Link>
                    </li>
                    <li className="nav-item pl-1">
                        <Link  data-cy="adminlink" className="nav-link" to="/super-admin/control-admins"><i className="fa-solid fa-user"></i>Admins</Link>
                    </li>
                    <li className="nav-item pl-1">
                        <Link data-cy="messagelink" className="nav-link" to="/super-admin/broadcast/received"><i className="fa-solid fa-bullhorn"></i>Broadcast </Link>
                    </li>
                    <li className="nav-item pl-1">
                        <Link data-cy="dictlink" className="nav-link" to="/super-admin/dic"><i className="fa-solid fa-spell-check"></i>Dictionary </Link>
                    </li>

                    <li className="nav-item pl-1 logout">
                            <Link className="nav-link" to="/" onClick={this.onLogout}><i className="fa fa-sign-out me-1" aria-hidden="true"></i>Logout </Link>
                        </li>
                </ul>
            </div>
            </div>
            </nav>
</header>

        )
    }
}



export default HeaderSuper
