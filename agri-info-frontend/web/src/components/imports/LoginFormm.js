import React from 'react'
import './loginFormm.css'
import {Switch, Route, Link} from 'react-router-dom'
import ErrorAlter from './ErrorAlter'


// const LoginFormm = () => {

//     // let { handleSubmit, setEmail, setPassword, errorMessage, setError } = props.loginState;
//     return (
//             <div className="col-sm-7 bg-color align-self-center">
//                 <div className="form-section">
//                     <div className="title">
//                     <h3>Sign into your account</h3>
//                     </div>
//                     <div className="login-inner-form">
//                         <form method="POST" action="/login">

//                             <div className="form-group form-box">
//                                 <input type="text" id="email"  className="input-text" placeholder="Email Address" />
//                                 <i className="icon email"></i>
//                             </div>

//                             <div className="form-group form-box">
//                                 <input type="text" id="password" className="input-text" placeholder="Password" />
//                                 <i className="icon lock"></i>
//                             </div>

//                             {/* {
//                                 errorMessage && <ErrorAlter errorMessage={errorMessage} clearError={() => setError(undefined) }></ErrorAlter>
//                             } */}

//                             <div className="form-group">
//                                 <button type="submit" className="btn primary-btn">Login</button>
//                             </div>

//                             <div className="form-group">
//                                 {/* <Link to='/products'><button className="btn secondary-btn">Forgot Password ?</button></Link> */}
//                                 <button className="btn secondary-btn">Forgot Password ?</button>
//                             </div>

//                         </form>
//                     </div>
//                 </div>
//             </div> 
//     )
// }

// export default LoginFormm


export default function LoginFormm(props) {

    let { handleSubmit, setEmail, setPassword, errorMessage, setError } = props.loginState;

    return (
        <div className="col-sm-7  bg-color align-self-center">
        <div className="form-section">
            <div className="title">
            <h3>Sign into your account</h3>
            </div>
            <div className="login-inner-form">
                <form method="POST" onSubmit={handleSubmit}>

                    <div className="form-group form-box">
                        <input type="text" id="email" onChange={e => setEmail(e.target.value)} className="input-text" placeholder="Email Address" />
                        <i className="icon email"></i>
                    </div>

                    <div className="form-group form-box">
                        <input type="text" id="password" onChange={e => setPassword(e.target.value)} className="input-text" placeholder="Password" />
                        <i className="icon lock"></i>
                    </div>

                    {
                        errorMessage && <ErrorAlter errorMessage={errorMessage} clearError={() => setError(undefined) }></ErrorAlter>
                    }

                    <div className="form-group">
                        <button className="btn primary-btn">Login</button>
                    </div>

                </form>
            </div>
        </div>
        </div>
    )
}

