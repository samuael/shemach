import React from 'react'
import { Link } from "react-router-dom";
import './headerSuper.css'


const HeaderSuper = () => {
    return (
        
          <header id="navitemsuper">
            <nav class="navbar navbar-expand-lg">
            <div class="container">
            {/* <Link to='/' className='navbar-logoo'><i class="fa-solid fa-list"></i>  Agri-Net <i class="fa-solid fa-building-wheat"></i> </Link> */}
            <Link class="navbar-brand text-white" to='/'><i class="fa-solid fa-list"></i> Agri Net </Link>
            <div class="collapse navbar-collapse" id="nvbCollapse">
                <ul class="navbar-nav ml-auto nav-menu">
                    <li class="nav-item  pl-1">
                        <Link class="nav-link" to="/super-admin/products"><i class="fa-brands fa-product-hunt"></i>Products <i class="fa-solid fa-circle-dollar"></i></Link>
                    </li>
                    <li class="nav-item pl-1">
                        <Link class="nav-link" to="/super-admin/control-admins"><i class="fa-solid fa-user"></i>Admins</Link>
                    </li>
                    <li class="nav-item pl-1">
                        <Link class="nav-link" to="/super-admin/broadcast"><i class="fa-solid fa-bullhorn"></i>Broadcast </Link>
                    </li>
                    <li class="nav-item pl-1">
                        <Link class="nav-link" to="/super-admin/dic"><i class="fa-solid fa-spell-check"></i>Dictionary </Link>
                    </li>
                    {/* <li class="nav-item pl-1">
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

    
    )
}

export default HeaderSuper
