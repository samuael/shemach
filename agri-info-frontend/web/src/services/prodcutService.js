import http from "../httpbase";

class prodcutService {

  getOne() {
    return http.get("/product");
  }

  getAll() {
    return http.get("/products");
  }

  get(id) {
    return http.get(`/product?id=${id}`);
  }

  create(data) {
    return http.post("/superadmin/product/new", data);
  }

  // create(data, token) {
  //   return http.post("/superadmin/product/new", data , {
  //     headers: {
  //     "Authorization"  : "Bearer "+ token ,
  //   }
  // });
  // }

  // update(data, token) {
  //   return http.put("/infoadmin/product", data , {
  //     headers: {
  //       "Authorization"  : "Bearer "+ token ,
  //     }
      
  //   });
   
  // }

  update(data){
     return http.put("/infoadmin/product", data);

  }

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

  registerAdmin(data) {
    return http.post("/superadmin/admin/new", data);
  }

  deleteAdmin(id) {
    return http.delete(`/superadmin/admin?id=${id}`);

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