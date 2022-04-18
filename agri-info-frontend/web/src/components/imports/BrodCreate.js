// import React from 'react'
import React, { useState } from 'react';
import styled from 'styled-components';
import './brodCreate.css'

const BrodCreate = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
    return (
        <>

          <form className="row createform">
            <div className="col-sm-5">
                  <div className="form-group">
                    <label htmlFor="name"  className="chkbox">
                    Show to Admins :
                      <input  className="chkboxiput"
                        type="checkbox"
                        id="name"
                        name="name"
                        value={name}
                        onChange={(e) => setName(e.target.value)}
                      />
                    </label>
                  </div>
                  <div className="form-group">
                    {/* <label htmlFor="email">
                      Your Email
                      <input
                        type="email"
                        id="email"
                        name="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                      />
                    </label> */}
                    <label for="prod">Choose a Product :</label>

                      <select name="prod" id="prod">
                        <option value="volvo">Teff</option>
                        <option value="saab">Nashila</option>
                        <option value="mercedes">Wheat</option>
                        <option value="audi">Sorghum</option>
                      </select>
                  </div>
                  <div className="form-group">
                    <label for="lang">Choose a Language :</label>

                      <select name="lang" id="lang">
                        <option value="volvo">Amharic</option>
                        <option value="saab">Oromic</option>
                        <option value="mercedes">Somalic</option>
                        <option value="audi">Sidex</option>
                      </select>
                  </div>
                </div>
                
                <div className="col-sm-4">
                  <div className="form-group">
                    <label htmlFor="message" className="infocreate">
                      Information : 
                      <textarea
                        type="text"
                        id="message"
                        name="message"
                        value={message}
                        placeholder="Information about the new product"
                        onChange={(e) => setMessage(e.target.value)}
                      />
                    </label>
                  </div>
                  <button className="btn btn-primary" type="submit">Send</button>
                  </div>
                  </form>
    </>

    )
}

export default BrodCreate
