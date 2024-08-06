import http from "../httpbase";
import LoginPage from "../components/imports/LoginPageImp"

class prodcutService {

  //User Login
  userlogin(data) {
    return http.post("/login", data);
  }




  getOne() {
    return http.get("/product");
  }

  getAll() {
    return http.get("/products");
  }

  get(id) {
    return http.get(`/product?id=${id}`);
  }

  // create(data) {
  //   return http.post("/superadmin/product/new", data);
  // }
      //Create new product
  create(data, token) {
    return http.post("/superadmin/product/new", data , {
      headers: {
      "Authorization"  : "Bearer "+ token ,
    }
  });
  }

  update(data, token) {
    return http.put("/infoadmin/product", data , {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
      
    });
   
  }

  // update(data){
  //    return http.put("/infoadmin/product", data);

  // }

  delete(id) {
    return http.delete(`/products/${id}`);

  }

  deleteAll() {
    return http.delete(`/products`);
  }

  findByTitle(name) {
    return http.get(`/product/search?text=${name}`);
  }

  //For Admins

  getAllAdmins() {
    return http.get("/infoadmins");
  }

  findAdminByName(name) {
    return http.get(`/infoadmin/search?name=${name}`);
  }


  //superadmin

      //superadmin register new infoadmin
  registerInfoAdmin(data, token) {
    return http.post("/superadmin/infoadmin/new", data, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
     );
  }
 
  // registerAdmin(data, token) {
  //   return http.post("/superadmin/admin/new", data, {
  //     headers: {
  //       "Authorization"  : "Bearer "+ token ,
  //     }
  //   }
  //    );
  // }
      

  //superadmin create dictionary
  createdict(data, token) {
    return http.post("/superadmin/dictionary/new", data, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
     );
  }

 // Superadmin SearchWord
  superSearchWord(data, token) {
    return http.post("/dictionary/translate", data, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
    );
  }

   // Superadmin ListRecent
   superListRecent(token) {
    return http.get(`/dictionaries?offset=0&limit=9`, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
    );
  }

   // Superadmin getDict
   getdict(id, token) {
    return http.get(`/superadmin/dictionary?id=${id}`, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
    );
  }

   // Superadmin updatedict
   updatedict(data, token) {
    return http.put("/superadmin/dictionary",data, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
    );
  }

     // Superadmin deletewprd
    //  deleteDictWord(data, token) {
    //   return http.delete(`/superadmin/dictionary?id=${data.id}`,data, {
    //     headers: {
    //       "Authorization"  : "Bearer "+ token ,
    //     }
    //   }
    //   );
    // }

    deleteDictWord(id, token) {
      return http.delete(`/superadmin/dictionary?id=${id}`, {
        headers: {
          "Authorization"  : "Bearer "+ token ,
        }
      }
      );
    }



    //Superadmin SearchWord
    // superSearchWord(data) {
    //   return http.get("/dictionary/translate", data);
    // }



  deleteInfoAdmin(id, token) {
    return http.delete(`/superadmin/infoadmin?id=${id}`, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
    );

  }

  getAdmin(id) {
    return http.get(`/infoadmin?id=${id}`);
  }



  //Superadmin messages
  // Superadmin sendmessage
   superSendMessage(data, token) {
    return http.post("/message/new", data, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
    );
  }

  //Superadmin getall messages
  superGetAllMessage(token) {
    return http.get("/admins/messages", {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
    );
  }

  //superdeletemessage
  deleteMessageSuper(id, token) {
    return http.delete(`/message/${id}`, {
      headers: {
        "Authorization"  : "Bearer "+ token ,
      }
    }
    );

  }



  
}

export default new prodcutService();





// import http from "../httpbase";

// class prodcutService {
//   getAll() {
//     return http.get("/tutorials");
//   }

//   get(id) {
//     return http.get(`/tutorials/${id}`);
//   }

//   create(data) {
//     return http.post("/tutorials", data);
//   }

//   update(id, data) {
//     return http.put(`/tutorials/${id}`, data);
//   }

//   delete(id) {
//     return http.delete(`/tutorials/${id}`);
//   }

//   deleteAll() {
//     return http.delete(`/tutorials`);
//   }

//   findByTitle(title) {
//     return http.get(`/tutorials?title=${title}`);
//   }
// }

// export default new prodcutService();