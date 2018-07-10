package templates

const SrcSagaIndex = `import axios, { AxiosPromise } from "axios";
import { select, call, put, all, takeEvery } from "redux-saga/effects";
import Action from "../interfaces/Action";
import {
	ON_APP_INIT,
	fetchConfig,
} from "../actions";

function* initialiseApp(action: Action) {
	yield put(fetchConfig());
}  

function* onAppInit() {
	yield takeEvery(ON_APP_INIT, initialiseApp);
}  

export default function* root() {
	let sagas = [
	  onAppInit(),
	];
  
	//sagas = sagas.concat(PageSagas);
  
	yield all(sagas);
  }
  
`
