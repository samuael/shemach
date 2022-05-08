import "../../libs.dart";


class DragShowInfoBloc extends Bloc<bool, bool> {
  DragShowInfoBloc() : super(false);

  Stream<bool>  mapEventToState(bool dstate) async* {
    yield dstate;
  }
}