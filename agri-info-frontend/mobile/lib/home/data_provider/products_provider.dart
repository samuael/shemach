import "package:http/http.dart";
import "../../libs.dart";
import "dart:convert";

class ProductsProvider {


  Client client  = Client();

  Future<ProductsResponse> loadProducts() async {
    try{
       var response = await client.get(
        Uri(
          host: StaticDataStore.HOST,
          port: StaticDataStore.PORT,
          scheme: StaticDataStore.SCHEME,
          path: "/api/products",
        )
      );
      if (response.statusCode == 200){
        final result  = ProductsResponse.fromJson(jsonDecode(response.body));
        return result;
      }else if (response.statusCode < 500 && response.statusCode >=200 ){
        return ProductsResponse(msg : jsonDecode(response.body)["msg"] , statusCode : response.statusCode , products : []);
      }else {
        return ProductsResponse(msg :"Server Problem, please try again!", statusCode: 500, products : []);
      }
    }catch(e , a ){
      return  ProductsResponse(msg :"Connection issue!", statusCode: 999, products : []);
    }
  }

}