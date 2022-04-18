import React from 'react'
import { useLocation, useHistory, Link } from 'react-router-dom';
import './baseloginn.css'
import '../pages/loginn.css'



const Baseloginn = () => {

    return (
        <div className="col-sm-5 bg-img align-self-center">
             <div className="info">
                <div className="logo clearfix">
                    <Link className="nav-brand" to="/">Agri-Net</Link>
                </div>
            </div>
            
            
        </div>
    )
}

export default Baseloginn

