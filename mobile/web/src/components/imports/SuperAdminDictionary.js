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
    this.refreshList = this.refreshList.bind(this);
    this.setActiveWord = this.setActiveWord.bind(this);
    this.onChangeTranslation = this.onChangeTranslation.bind(this);
    this.saveDictionary = this.saveDictionary.bind(this);
   this.onChangeSearchWord = this.onChangeSearchWord.bind(this);
    this.retrieverecentDicts = this.retrieverecentDicts.bind(this);
    this.deleteWord = this.deleteWord.bind(this);
     this.searchWord = this.searchWord.bind(this);

    this.state = {
    //  tkValue: this.props.location.state,
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
        dictionaries: [],
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

componentDidMount() {
  console.log(this.props)
  // this.retrieverecentDicts();

  var tokenn = window.token;
  if(!tokenn){
   this.props.history.push('/')
  }else{
   this.retrieverecentDicts();
  }

}



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

  refreshList() {
    this.retrieverecentDicts();
    this.setState({
      currentWord: null,
      currentIndex: -1
    });
  }

  setActiveWord(Word, index) {
    this.setState({
      currentWord: Word,
      currentIndex: index
    });
  }
  

  retrieverecentDicts() {
    var token = window.token
    console.log(token)
    DictionaryDataService.superListRecent(token)
      .then(response => {
        this.setState({
          msg: response.data.msg,
          status_code: response.data.status_code,
          dictionaries: response.data.dictionaries,
        });
        console.log(response.data);
        console.log(this.state.dictionaries);
      })
      .catch(e => {
        console.log(e);
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
          // word: {
          //   id: response.data.word.id,
          //   name: response.data.word.name,
          //   Wordion_area: response.data.word.Wordion_area,
          //   unit_id: response.data.word.unit_id,
          //   current_price: response.data.word.current_price,
          //   created_by:response.data.word.current_price,
          //   created_at: response.data.word.created_at,
          //   last_update_time: response.data.word.last_update_time
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
        this.setState({
          status_code: response.data.status_code,
          dictionaries: response.data.dictionary,
          // dictionary: {
          //   id: response.data.dictionary.id,
          //   lang: response.data.dictionary.lang,
          //   text: response.data.dictionary.text,
          //   translation: response.data.dictionary.translation
          // },
         // dictionaries: Object.keys(response.data.dictionary)
        
          
         
        });
        // console.log(this.state.dictionary);
        // var obj = this.state.dictionary;  
        // console.log(obj);  
        // this.state.dictionaries = [];  
        // this.state.dictionaries = Object.values(obj);   
        // this.state.dictionaries[0] = obj; 
        // for(var i in obj){
        //     this.state.dictionaries.push(obj[i]);
        // }
        console.log(this.state.dictionaries);



        // var dictionaries = [];
        // for (var key in this.state.dictionary){
        //   if(this.state.dictionary.hasOwnProperty(key)){
        //     var item = this.state.dictionary[key];
        //     this.dictionaries.push({
        //       id: item.id,
        //       lang: item.lang,
        //       text: item.text,
        //       translation: item.translation
        //     });
        //   }
        // }
        // console.log(this.dictionaries);
        console.log(response.data);
      })
      .catch(e => {
        console.log(e);
      });
  }


  deleteWord() {  
    
    var data = {
      id : this.state.currentWord.id,
      lang: this.state.currentWord.lang,
      text: this.state.currentWord.text,
      translation: this.state.currentWord.translation
    }

    console.log(data);
    console.log(data.id);
    var token  = window.token
    console.log(token); 
    console.log(this.state.currentWord.id);

    DictionaryDataService.deleteDictWord(this.state.currentWord.id, token)
      .then(response => {
        console.log(response.data);
        // this.searchWord()
       // this.props.history.push('/super-admin/control-admins')
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
   const { dictionaries, dictionary, currentIndex, currentWord} = this.state;
   console.log(dictionaries);
   console.log(currentWord);
   console.log(dictionary);
//console.log(response_data.word);
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
                                    <select data-cy="selectlang" name="lang" id="lang" required  value={this.state.lang} onChange={this.onChangeLanguage} type="text">
                                            <option value="DEFAULT" >Select Language</option>
                                            <option value="amh">Amharic</option>
                                            <option value="oro">Oromic</option>
                                            <option value="som">Somalic</option>
                                            <option value="sid">Sidama</option>
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
                                        data-cy = "inputkey"
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
                                    data-cy = "inputvalue"
                                    />
                                </div>
                            </div>

                            <div className="col-sm-2 mt-4 addbutton">
                                <button data-cy="addictbtn" className="btn btn-primary" type="submit" onClick={this.saveDictionary} >Add</button>
                            </div>

                            <div className="col-sm-2 mt-4">
                            <p>{this.state.msg}</p>
                            {/* <p>Faniman</p> */}
                            </div>
                     </div>
                 </form>

                
                    <div className="row mt-4 ml-5">
                    <div className="col-md-8 col-md-offset-2 main">
                      <a className="page-header" onClick={this.retrieverecentDicts}>Recent Translations</a>
                    </div>
                </div>

                <div className="row">
                    <div className="col-sm-10">
                    <div>
                        <table className="table table-striped table-bordered table-primary table-hover">
                        <thead>
                          <tr>
                            <th scope="col">Language</th>
                            <th scope="col">Key(Eng)</th>
                            <th scope="col">Value</th>
                            {/* <th scope="col">Action</th> */}
                          </tr>
                        </thead>
                        </table>
                  </div>

                  <div className="diclist">
                  <ul className="list-group">
            {
              dictionaries.map((word, index) => (
                <div className={
                  "list-group-item " +
                  (index === currentIndex ? "active" : "")
                }
                onClick={() => this.setActiveWord(word, index)}
                key={index}
                >                
                    <div className="row listofdicts">
                    <p className="col-sm-5">{word.lang}</p>
                              <p className="col-sm-4 listd">{word.text}</p>
                              <p className="col-sm-3 listd">{word.translation}</p>
                              {/* <p className="col-sm-3 listd">
                              <Link
                                to={"/super-admin/edit-dict/" }
                                  className=""
                                >
                                  Edit
                                </Link>   --- 

                                <button
                                    className="btn-danger"
                                    onClick={this.deleteWord}
                                  >
                                    Delete
                                </button>


                              </p> */}
                    </div>
                
                </div>

              ))}
          </ul>
          </div>
                  </div>

                  

                  <div className="col-sm-2 editing">
                    {currentWord ? (
                      <div className="row ">
                          <div className="col-sm-1 alignbuttons">
                            <Link
                            className="linkeditiword"
                            // to={"/super-admin/edit-dict/" + currentWord.id }
                            // state = {currentWord}

                            to={{
                              pathname: "/super-admin/edit-dict/" + currentWord.id,
                              state: currentWord
                          }}
                            >
                              Edit
                            </Link>
                          </div>

                            {/* <div className="col-sm-1">
                          <button
                          className="badge badge-danger mr-2 btn btn-primary"
                          onClick={this.deleteWord}
                        >
                          Delete
                        </button>

                          </div> */}
                      </div>

                            

                            
                    ) : (
                      <div className="wordclickon">
                       <p></p>
                      </div>
                    )}

                  </div>
                </div>
              

      
         </div>
    
    );
  }
}


export default withRouter(SuperAdminDictionary)



