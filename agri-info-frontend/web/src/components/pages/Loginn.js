import React, { useState } from 'react'
import './loginn.css'
import Navbar from '../../components/Navbar';

import BaseLoginn from '../imports/Baseloginn'
import LoginFormm from '../imports/LoginFormm'

import { useDispatch, useStore } from 'react-redux';
import { loginAction  } from '../../container/actions'
import { useHistory } from 'react-router-dom';

export default function Login() {

    const [email, setEmail] = useState("")
    const [password, setPassword] = useState("")
    const [errorMessage, setError] = useState("")

    const dispatch = useDispatch()
    const store = useStore()
    const history = useHistory()

    // handle Submit handler function
    const handleSubmit = (event) =>{
        event.preventDefault();

        const userCredential = {
            email,
            password
        }
        // const userdata = { email: "fantahunfekadu1@gmail.com", password : "admin123"}
        const login = dispatch(loginAction(userCredential))
        login
            .then(data => {
                if (data.role == "infoadmin"){
                    history.push('/products')
                }
                else if (data.role == "superadmin"){
                    history.push('/super-admin/products')
                }
                else if (data.role == "admin"){
                    history.push('/super-admin/products')
                }
                else if (data.role == "agent"){
                    history.push('/super-admin/products')
                }
                else if (data.role == "merchant"){
                    history.push('/super-admin/products')
                }
                else {
                    history.push('/super-admin/products')
                }
                console.log(data.role);
                
            }
                // history.push('/products') 
            
            )
            .catch(error => { 
                console.log(error.err)
                setError(error.err)
            })

        // dispatch(loginAction(userCredential)).then(data => history.push('/products'))
        //     .catch(err => console.log(err))

            // setError(error.err)

        // console.log(store.getState())
        // console.log(userCredential);
    }

     
    return (
        <>
            <Navbar />
           <div id="login">
            {/* <Header/> */}
           
            <div className="container">
                <div className="row login-box">
                    <BaseLoginn />
                    {/* <LoginFormm  /> */}
                    <LoginFormm loginState={{handleSubmit, setEmail, setPassword, errorMessage, setError}} />

                    {/* <LoginFormm /> */}
                </div>
            </div>
      </div>  
      </>
    )
}

// export default Loginn
