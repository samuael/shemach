import React, { Component } from "react";
import currentDictionaryDataService from "../../services/prodcutService";
import './superAdminDictionaryEdit.css'
import { withRouter } from 'react-router-dom'
import './headerInfo.css'   //For Header
import {Switch, Route, Link} from 'react-router-dom'

class SuperAdmincurrentDictionaryEdit extends Component {
  constructor(props) {
    super(props);
    this.onChangeCurrLang = this.onChangeCurrLang.bind(this);
    this.onChangeCurrText = this.onChangeCurrText.bind(this);
    this.onChangeCurrTranslation = this.onChangeCurrTranslation.bind(this);
    this.getWord = this.getWord.bind(this);
    this.updateWord = this.updateWord.bind(this);
    this.deleteWord = this.deleteWord.bind(this);
   // this.dictdata = this.dictdata.bind(this);

    this.state = {
   //  currentDictionary: this.props.location.state,
      // currentDictionary: { 
         msg: "",
          id: null,
          lang: "",
          text: "",
          translation: ""
        // },
       
    };
  }

  componentDidMount() {
    this.getWord();
    // var dicdataa = this.props.location.state;
    // console.log(this.props.location.state);
    // console.log(dicdataa);


  }

  onChangeCurrLang(e) {
      const lang = e.target.value;
    this.setState({
        lang: lang
      });

      // this.setState(function(prevState) {
      //   return {
      //     currentDictionary: {
      //       ...prevState.currentDictionary,
      //       lang: lang
      //     }
      //   };
      // });
  }

  onChangeCurrText(e) {
    this.setState({
        text: e.target.value
      });
  }


  onChangeCurrTranslation(e) {
    this.setState({
        translation: e.target.value
      });
  }

//   onChangeSearchAdmin(e) {
//     const searchAdmin = e.target.value;

//     this.setState({
//       searchAdmin: searchAdmin
//     });
//   }


  // onChangeCurrLang(e) {
  //   this.setState({
  //     current_price: e.target.value
  //   });
  // }



  getWord() {
    // currentDictionaryDataService.getdict(id)
    //  .then(response => {
        this.setState({
        //  currentDictionary: response.data
        // dicdata: this.props.location.state,
      //    msg:"",
      //    status_code: 200,
      //  currentDictionary: this.props.location.state,
        //   currentDictionary: response.data.currentDictionary
        // currentDictionary: { 
              id: this.props.location.state.id,
              lang: this.props.location.state.lang,
              text: this.props.location.state.text,
              translation: this.props.location.state.translation
        //     },
            
        });
        console.log(this.props.location.state)
      //  console.log(response.data);
      //  console.log(this.state);
      //  console.log(this.state.currentDictionary);
       // console.log(response.data.product);
        
    //   })
    //   .catch(e => {
    //     console.log(e);
    //   });
  }

  updateWord() {
    var data = {
      id:parseInt(this.state.id),
      lang:this.state.lang,
      text:this.state.text,
      translation:this.state.translation
      // id:28,
      // lang: 'oro',
      // text:'faniman',
      // translation:'daniman'

    }

  //  var token = this.state.tkValue;
  var token = window.token

    
    console.log(data);
    console.log(token);
    currentDictionaryDataService.updatedict(data, token
      // this.state.currentDictionary.id,
      // this.state.currentDictionary.current_price
      // parseInt(this.state.currentDictionary.id),
      // parseInt(this.state.currentDictionary.current_price)
    )
      .then(response => {
        console.log(response.data);
        this.setState({
           msg: "Translation updated successfully!"
        });
        // this.setState({

        //   msg: "The product was updated successfully!"
        // });
      })
      .catch(e => {
        console.log(e);
      });
  }

  deleteWord() {    
      var token = window.token;
      currentDictionaryDataService.deleteDictWord(this.state.id, token)
      .then(response => {
        console.log(response.data);
        this.props.history.push('/super-admin/dic')
      })
      .catch(e => {
        console.log(e);
      });
  }

  render() {
   const {currentDictionary, dicdata} = this.state;
   console.log(dicdata);
   console.log(currentDictionary);
   console.log(this.state.currentDictionary);
  //  console.log(res_data.product);
  //  console.log(res_data.msg);
  // console.log(res_data.product.lang);

    return (
      <>
                <header id="navitem">
                <nav className="navbar navbar-expand-lg">
                <div className="container">
                <Link className="navbar-brand text-white" to='/super-admin/dic'><i className="fa-solid fa-angles-left"></i> Home </Link>
                <div className="collapse navbar-collapse" id="nvbCollapse">
                    <ul className="navbar-nav ml-auto nav-menu">
                        <p data-cy="naveditdict" className="nav-item pl-1 nav-link">Super Admin currentDictionary Edit Page</p>
                    </ul>
                </div>
                </div>
                </nav>
	</header>

      <div id="dictcomponent">
        {this.state ? (
          <div className="edit-form mt-4">
            <h4>Word</h4>
            <form className="row infoeditprice">
              <div className="col-sm-6">
              {/* <div className="form-group">
                <label htmlFor="title">Lang</label>
                <input
                  type="text"
                  className="form-control"
                //  disabled="disabled"
                  id="title"
                  value={this.state.lang}
                  onChange={this.onChangeCurrLang}
                />
              </div> */}

              <div className="form-group">
              <label data-cy="langlabel" htmlFor="lang">Lang</label>
              <select data-cy="selecttranslation" name="lang" id="lang" required className="form-control" value={this.state.lang} onChange={this.onChangeCurrLang} type="text">
                  {/* <option value="" selected disabled hidden>Choose here</option> */}
                      <option value="amh">amh</option>
                      <option value="oro">oro</option>
                      <option value="som">som</option>
                      <option value="sid">sid</option>
             </select>
            </div>


              <div className="form-group">
                <label data-cy="textlabel" htmlFor="description">Text</label>
                <input
                  type="text"
                  className="form-control"
                //  disabled="disabled"
                  id="description"
                  value={this.state.text}
                  onChange={this.onChangeCurrText}
                  data-cy="texttranslation"
                />
              </div>
              <div className="form-group">
                <label data-cy="translabel"htmlFor="productionarea">Translation</label>
                <input
                  type="text"
                  className="form-control"
               //   disabled="disabled"
                  id="productionarea"
                  value={this.state.translation}
                  onChange={this.onChangeCurrTranslation}
                  data-cy="wordtranslation"
                />
              </div>
  
              </div>
            </form>

            {/* {product.published ? (
              <button
                classlang="badge badge-primary mr-2"
                onClick={() => this.updatePublished(false)}
              >
                UnPublish
              </button>
            ) : (
              <button
                classlang="badge badge-primary mr-2"
                onClick={() => this.updatePublished(true)}
              >
                Publish
              </button>
            )} */}


            <button
              type="submit"
              className="badge badge-success btn btn-primary"
            
              onClick={this.updateWord}
              data-cy="updatetranslation"
            >
              Update
            </button>
            <p>{this.state.msg}</p>

            <button
              className="badge badge-danger mr-2 btn btn-primary"
              onClick={this.deleteWord}
              data-cy="deletetranslation"
            >
              Delete
            </button> 
            {/* <p>{this.state.message}</p> */}

  
          </div>
        ) : (
          <div>
            <br />
            <p>Please click on word</p>
          </div>
        )}
      </div>
      </>
    );
  }
}

export default withRouter(SuperAdmincurrentDictionaryEdit)

