import React from 'react'
import './headerInfo.css'
import {Switch, Route, Link} from 'react-router-dom'

const header = () => {
    return (
        <div>
            {/* <nav className="navbar navbar-expand navbar-dark bg-dark justify-content-center">
                <Link to={"/tutorials"} className="navbar-brand">
                    Fani Tutorials
                </Link>
            <div className="navbar-nav mr-auto">
                <li className="nav-item">
                <Link to={"/tutorials"} className="nav-link">
                    Product Price
                </Link>
                </li>
                <li className="nav-item">
                <Link to={"/add"} className="nav-link">
                    Broadcast Info
                </Link>
                </li>
            </div>
            </nav> */}
              <header id="navitem">
                <nav class="navbar navbar-expand-lg">
                <div class="container">
                {/* <Link to='/' className='navbar-logoo'><i class="fa-solid fa-list"></i>  Agri-Net <i class="fa-solid fa-building-wheat"></i> </Link> */}
                <Link class="navbar-brand text-white" to='/'><i class="fa-solid fa-list"></i> Agri Net </Link>
                {/* <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#nvbCollapse" aria-controls="nvbCollapse">
                    <span class="navbar-toggler-icon"></span>
                </button> */}
                <div class="collapse navbar-collapse" id="nvbCollapse">
                    <ul class="navbar-nav ml-auto nav-menu">
                        {/* <li class="nav-item active pl-1">
                            <Link class="nav-link" to="/marketing"><i class="fa-brands fa-product-hunt"></i>Products <i class="fa-solid fa-circle-dollar"></i></Link>
                        </li> */}
                        <li class="nav-item active pl-1">
                            <Link class="nav-link" to="/products"><i class="fa-solid fa-hand-holding-dollar"></i>Products & Price <i class="fa-solid fa-circle-dollar"></i></Link>
                        </li>
                        <li class="nav-item pl-1">
                            <Link class="nav-link" to="/broadcast/received"><i class="fa-solid fa-bullhorn"></i>Broadcast Info </Link><span class="circle">5</span>
                        </li>
                        {/* <li class="nav-item pl-1">
                            <a class="nav-link" href="#"><i class="fa fa-info-circle fa-fw mr-1"></i>Hakkımızda</a>
                        </li>
                        <li class="nav-item pl-1">
                            <a class="nav-link" href="#"><i class="fa fa-phone fa-fw fa-rotate-180 mr-1"></i>İletişim</a>
                        </li>
                        <li class="nav-item pl-1">
                            <a class="nav-link" href="#"><i class="fa fa-user-plus fa-fw mr-1"></i>Kayıt Ol</a>
                        </li>
                        <li class="nav-item pl-1">
                            <a class="nav-link" href="#"><i class="fa fa-sign-in fa-fw mr-1"></i>Oturum Aç</a>
                        </li> */}
                    </ul>
                </div>
                </div>
                </nav>
      	</header>

            
        </div>
    )
}

export default header
