import React from 'react'
import './loginn.css'
import Navbar from '../../components/Navbar';

import BaseLoginn from '../imports/Baseloginn'
import LoginFormm from '../imports/LoginFormm'

const Loginn = () => {
    return (
        <>
         <Navbar />
           <div id="login">
            {/* <Header/> */}
           
            <div className="container">
                <div className="row login-box">
                    <BaseLoginn />
                    <LoginFormm />
                    {/* <LoginFormm /> */}
                </div>
            </div>
      </div>  
      </>
    )
}

export default Loginn
