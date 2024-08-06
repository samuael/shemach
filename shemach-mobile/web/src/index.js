import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';
import 'bootstrap/dist/css/bootstrap.min.css';
import { Provider } from 'react-redux';

//import  store from './container/store'




 ReactDOM.render(<App />, document.getElementById('root'));


// ReactDOM.render(
//     <Provider store={store}>
//       <App />
//     </Provider>,
//     document.getElementById('root')
//   );
  
  //console.log(store.getState())
