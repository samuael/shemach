import "package:flutter_bloc/flutter_bloc.dart";

class IndexBloc extends Bloc<int , int >{
  IndexBloc() : super(1){
    on<int>((event , emit ){
      emit(event);
    });
  }
}