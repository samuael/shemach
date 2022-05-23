import React, { Component } from "react";
import DictionaryDataService from "../../services/prodcutService";
import { Link } from "react-router-dom";
import { withRouter } from 'react-router-dom'
import './superAdminDictionary.css'
import './headerInfo.css'    //For Header



class SuperAdminDictionary extends Component {
  constructor(props) {
    super(props);
  //  this.onChangeProdName = this.onChangeProdName.bind(this);
    this.onChangeLanguage = this.onChangeLanguage.bind(this);
    this.onChangeText = this.onChangeText.bind(this);
   // this.onChangeMeasurement = this.onChangeMeasurement.bind(this);
    this.onChangeTranslation = this.onChangeTranslation.bind(this);
    this.saveDictionary = this.saveDictionary.bind(this);
   this.onChangeSearchWord = this.onChangeSearchWord.bind(this);
  //  this.response_data = this.response_data.bind(this);
     this.searchWord = this.searchWord.bind(this);

    this.state = {
    //   tkValue: this.props.location.state,
    //  response_data: {
        msg: "",
        status_code: 0,
        dictionary: {
          id: null,
          lang: "",
          text: "",
          translation: ""
        },
        
    //  },
        submitted: false,
        currentWord: null,
        currentIndex: -1,
        searchWord: ""

    };
  }

//   onChangeProdName(e) {
//     this.setState({
//       name: e.target.value
//       });
//   }

onChangeSearchWord(e) {
  // const searchWord = e.target.value;

  this.setState({
    searchWord: e.target.value
  });
}



  onChangeLanguage(e) {
    this.setState({
        lang: e.target.value
      });
  }


  

 

  onChangeText(e) {
    this.setState({
        text: e.target.value
      });
  }

  


  onChangeTranslation(e) {
    this.setState({
        translation: e.target.value
      });
  }
  

 


  saveDictionary(event) {
    event.preventDefault();
    var data = {
     // unit_id: parseInt(this.state.unit_id),
      lang: this.state.lang,
      text: this.state.text,
      translation:this.state.translation

      // unit_id: 6,
    //   lang: "Faniman",
    //   text: "AA",
    //   translation:"Faaya"
    };

    // var token = this.state.tkValue;
    var token = window.token;
    console.log(data);
    console.log(token);

    DictionaryDataService.createdict(data, token)
      .then(response => {
        console.log(response.data)
        this.setState({

       // response_data: {parseInt(
          msg: response.data.msg,
          status_code: response.data.status_code,
          dictionary: {
            id: response.data.dictionary.id,
            lang: response.data.dictionary.lang,
            text: response.data.dictionary.text,
            translation: response.data.dictionary.translation
          },
          // product: {
          //   id: response.data.product.id,
          //   name: response.data.product.name,
          //   production_area: response.data.product.production_area,
          //   unit_id: response.data.product.unit_id,
          //   current_price: response.data.product.current_price,
          //   created_by:response.data.product.current_price,
          //   created_at: response.data.product.created_at,
          //   last_update_time: response.data.product.last_update_time
          // },
          
       // }
          submitted: true
        
        });
        console.log(response.data);
      //  console.log(response_data);
      })
      .catch(e => {
        console.log(e);
      });
  }

  searchWord() {
    this.setState({
      currentWord: null,
      currentIndex: -1
    });

    var data = {
      lang: this.state.lang,
      text: this.state.searchWord
    }

    var token = window.token

    console.log(data);
    console.log(token);

    DictionaryDataService.superSearchWord(data, token)
      .then(response => {
        console.log(response.data)
        // this.setState({
        //   status_code: response.data.status_code,
        //   dictionary: {
        //     id: response.data.dictionary.id,
        //     lang: response.data.dictionary.lang,
        //     text: response.data.dictionary.text,
        //     translation: response.data.dictionary.translation
        //   }
        // });
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }



//   newDictionary() {
//     this.state = {
//     //  response_data: {
//         msg: "",
//         product: {
//           id: null,
//           name: "",
//           production_area: "",
//           unit_id: 0,
//           current_price: 0,
//           created_by:0,
//           created_at: 0,
//           last_update_time: 0
//         },
//         status_code: 0,
//     //  }
//         submitted: false
//     };
//   }

  render() {
  //  const {product} = this.state;
  //  console.log(product);
//console.log(response_data.product);
//const { searchWord } = this.state;
    return (

        
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
                                    <label htmlFor="lang">Choose a Language :</label>
                                    <select name="lang" id="lang" required  value={this.state.lang} onChange={this.onChangeLanguage} type="text">
                                            <option value="DEFAULT" >Select Language</option>
                                            <option value="amh">Amharic</option>
                                            <option value="oro">Oromic</option>
                                            <option value="som">Somalic</option>
                                            <option value="sid">Sidex</option>
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
                                    value={this.state.searchWord}
                                    onChange={this.onChangeSearchWord}
                                    />
                                    <div className="input-group-append">
                                        <button
                                            className="btn btn-outline-secondary searchbuttom"
                                            type="button"
                                            onClick={this.searchWord}
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
                                        required
                                        value={this.state.text}
                                        onChange={this.onChangeText}
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
                                    required
                                    value={this.state.translation}
                                    onChange={this.onChangeTranslation}
                                    />
                                </div>
                            </div>

                            <div className="col-sm-2 mt-4 addbutton">
                                <button className="btn btn-primary" type="submit" onClick={this.saveDictionary} >Add</button>
                            </div>

                            <div className="col-sm-2 mt-4">
                            <p>{this.state.msg}</p>
                            {/* <p>Faniman</p> */}
                            </div>
                     </div>
                 </form>

                
                    <div className="row mt-4 ml-5">
                    <div className="col-md-8 col-md-offset-2 main">
                      <h3 className="page-header">Recent Searches</h3>

                      <div className="table">
                        <table className="table table-striped">
                          <thead>
                            <tr>
                              <th>Language</th>
                              <th>Key(Eng)</th>
                              <th>Value</th>
                              <th>Action</th>
                            </tr>
                          </thead>
                          <tbody>
                          <td>Faniman</td>
                              <td>Daniman</td>
                              <td>English</td>
                              <td><a href="dictionary?id={{Lang.ID}}">edit</a> -  <a href="/dictionary?id={{.ID}}">delete</a></td>
                            {/* {{range .AllBooks}}
                            <tr>
                              <td>{{.ID}}</td>
                              <td>{{.Name}}</td>
                              <td>{{.Author}}</td>
                              <td>{{.Pages}}</td>
                              <td>{{.PublicationDateStr}}</td>
                              <td><a href="book.html?id={{Lang.ID}}">edit</a> -  <a href="/delete?id={{.ID}}">delete</a></td>
                            </tr>
                            {{end}} */}
                            {/* <tr>
                              <td></td>
                              <td></td>
                              <td></td>
                              <td></td>
                              <td></td>
                              <td><a href="book.html" class="btn btn-primary">New book</a></td>
                            </tr> */}
                          </tbody>
                        </table>
                      </div>
                    </div>
                </div>

              </div>
    
    );
  }
}


export default withRouter(SuperAdminDictionary)



