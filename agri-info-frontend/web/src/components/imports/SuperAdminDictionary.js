import React, { useState } from 'react';
import './superAdminDictionary.css'

const SuperAdminDictionary = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
    return (
        <>
            <div id="dict">
                <div className="row">
                    <div className="col-sm-8">
                        <h3 className="topic">Add Dictionary Key : Value Here</h3>
                    </div>
                </div>

                <form className="row dictionary">
                        <div className="row">
                            <div className="col-sm-5 mt-4">
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
                        </div>

                        <div className="row">
                            <div className="col-sm-4">
                                <div className="input-group mb-5">
                                    <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Search words"
                                    //   value={searchProduct}
                                    //   onChange={this.onChangesearchProduct}
                                    />
                                    <div className="input-group-append">
                                        <button
                                            className="btn btn-outline-secondary searchbuttom"
                                            type="button"
                                            // onClick={this.searchProduct}
                                        >
                                            Search
                                        </button>
                                    </div>
                                </div>
                            </div>
                        </div>


                        <div className="row mb-3">
                            <div className="col-sm-4">
                                    <h6>Key</h6>
                                    <div className="input-group mb-3">
                                        <input
                                        type="text"
                                        className="form-control"
                                        placeholder="Enter Key"
                                        //   value={searchProduct}
                                        //   onChange={this.onChangesearchProduct}
                                        />
                                    </div>
                            </div>

                            <div className="col-sm-4">
                                <h6>Value</h6>
                                <div className="input-group mb-3">
                                    <input
                                    type="text"
                                    className="form-control"
                                    placeholder="Enter Value"
                                    //   value={searchProduct}
                                    //   onChange={this.onChangesearchProduct}
                                    />
                                </div>
                            </div>

                            <div className="col-sm-4 mt-4 addbutton">
                                <button className="btn btn-primary" type="submit">Add</button>
                            </div>
                     </div>
                 </form>
            </div>
       </>

    )
}

export default SuperAdminDictionary


