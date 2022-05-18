import React from 'react'
import './headerBrod.css'
import {Switch, Route, Link} from 'react-router-dom'

const header = () => {
    return (
        <div>
              <header id="navbrod">
                <nav class="navbar navbar-expand-lg">
                    <div class="container">
                    {/* <Link to='/' className='navbar-logoo'><i class="fa-solid fa-list"></i>  Agri-Net <i class="fa-solid fa-building-wheat"></i> </Link> */}
                    <Link class="navbar-brand text-white" to='/super-admin/control-admins'><i class="fa-solid fa-angles-left"></i> Home </Link>
                    <div class="collapse navbar-collapse" id="nvbCollapse">
                        <ul class="navbar-nav ml-auto nav-menu">
                            <li class="nav-item pl-1">
                                <Link class="nav-link" to="/super-admin/broadcast/received"><i class="fa-solid fa-folder-open"></i>Received</Link>
                            </li>
                            <li class="nav-item pl-1">
                                <Link class="nav-link" to="/super-admin/broadcast/create"><i class="fa-solid fa-calendar-plus"></i>Create</Link>
                            </li>
                            <li class="nav-item pl-1">
                                <Link class="nav-link" to="/broadcast"><i class="fa-solid fa-envelope-circle-check"></i>Sent</Link>
                            </li>
                        </ul>
                     </div>
                    </div>
                </nav>
        	</header>

            
        </div>
    )
}

export default header
