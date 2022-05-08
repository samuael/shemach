import http from "../httpbase";

class prodcutService {

  getOne() {
    return http.get("/product");
  }

  getAll() {
    return http.get("/products");
  }

  get(id) {
    return http.get(`/products/${id}`);
  }

  create(data) {
    return http.post("/superadmin/product/new", data);
  }

  update(id, data) {
    return http.put(`/infoadmin/product/${id}`, data);
  }

  delete(id) {
    return http.delete(`/products/${id}`);

  }

  deleteAll() {
    return http.delete(`/products`);
  }

  findByTitle(title) {
    return http.get(`/product/search?title=${title}`);
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