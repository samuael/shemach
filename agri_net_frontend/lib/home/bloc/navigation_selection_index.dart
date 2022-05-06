import "../../libs.dart";

class NavigationIndexBloc extends Bloc<int, int> {
  NavigationIndexBloc(int initialState) : super(initialState);

  @override
  Stream<int> mapEventToState(int event) async* {
    yield event;
  }
}
