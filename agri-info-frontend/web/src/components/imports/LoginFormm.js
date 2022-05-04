import React from 'react'
import './loginFormm.css'
import {Switch, Route, Link} from 'react-router-dom'


const LoginFormm = () => {
    return (
            <div className="col-sm-7 bg-color align-self-center">
                <div className="form-section">
                    <div className="title">
                    <h3>Sign into your account</h3>
                    </div>
                    <div className="login-inner-form">
                        <form method="POST">

                            <div className="form-group form-box">
                                <input type="text" id="email" className="input-text" placeholder="Email Address" />
                                <i className="icon email"></i>
                            </div>

                            <div className="form-group form-box">
                                <input type="text" id="password" className="input-text" placeholder="Password" />
                                <i className="icon lock"></i>
                            </div>

                            {/* {
                                errorMessage && <ErrorAlter errorMessage={errorMessage} clearError={() => setError(undefined) }></ErrorAlter>
                            } */}

                            <div className="form-group">
                                <button className="btn primary-btn">Login</button>
                            </div>

                            <div className="form-group">
                                {/* <Link to='/products'><button className="btn secondary-btn">Forgot Password ?</button></Link> */}
                                <button className="btn secondary-btn">Forgot Password ?</button>
                            </div>

                        </form>
                    </div>git 
                </div>
            </div> 
    )
}

export default LoginFormm
