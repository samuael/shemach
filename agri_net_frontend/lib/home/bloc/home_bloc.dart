import '../../libs.dart';

class HomeBloc extends Bloc<HomeEvent, HomeBlocState> {
  HomeBloc(/*{required this.repo}*/) : super(HomeStateInit());
  @override
  Stream<HomeBlocState> mapEventToState(HomeEvent event) async* {
    if (event is AuthLoginEvent) {
      yield (HomeStateInit());
    }
  }
}
