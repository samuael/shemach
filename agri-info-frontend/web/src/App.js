import React from 'react';
import Navbar from './components/Navbar';
import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './components/pages/LoginPage';
// import Home from './components/pages/Loginn';
import { BrowserRouter as Router, Switch, Route, withRouter } from 'react-router-dom';
import Products from './components/pages/InfoAdmin';
import SuperProducts from './components/pages/SuperAdminProd';
import ProductComponent from './components/imports/ProductComponent';
import SuperControlAdmins from './components/pages/SuperControlAdmins';
import SuperAdminAddPr from './components/imports/SuperAdminAddPr';
import SuperAdminRegAdmin from './components/imports/SuperAdminRegisterAdmin';
import Broadreceive from './components/pages/BroadcastRec';
// import BroadCreate from './components/pages/BroadcastCreate';
// import SuperProductEdit from './components/imports/SuperProductEdit'
// import SuperEditAdmin from './components/imports/SuperEditAdmin'
import SuperAdminDictionary from './components//pages/SuperAdminDictionary'
import SuperEditDictionary  from './components/imports/SuperAdminDictionaryEdit'
// import ProductComponent '/components/imports/ProductComponent';

function App() {
  return (
    // <div className="container mt-3"></div>
    <Router>
      {/* <Navbar /> */}
      <Switch>
        <Route path='/' exact component={Home} />
        <Route path='/login' exact component={Products} />
        <Route path='/products' component={Products} />
        <Route path='/super-admin/products' component={SuperProducts} />
        {/* <Route path='/products' exact component={Products} /> */}
        <Route path='/super-admin/broadcast/received' exact component={Broadreceive} />
        {/* <Route path='/super-admin/broadcast/create' exact component={BroadCreate} /> */}
        {/* <Route path="/super/products/:id" component={SuperProductEdit} /> */}
        {/* <Route path='/super-admin/admins/:id' component={SuperEditAdmin} /> */}
        <Route path="/info/product/:id" component={ProductComponent} />
        <Route path='/super-admin/add-product' component={SuperAdminAddPr} />
        <Route path='/super-admin/reg-admin' component={SuperAdminRegAdmin} />
        <Route path='/super-admin/control-admins' component={SuperControlAdmins} />
        <Route path='/super-admin/edit-dict/:id' component={SuperEditDictionary} />
        <Route path='/super-admin/dic' component={SuperAdminDictionary} />
      </Switch>
    </Router>

  );
}

export default App;
