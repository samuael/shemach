class ImageUploadResponse {
  int id;
  String imgurl;
  int statusCode;
  String? msg;
  ImageUploadResponse(this.id, this.statusCode, this.imgurl, {this.msg});
}
