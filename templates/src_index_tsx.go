package templates

const SrcIndexTsxTemplate = `import * as React from "react";
import * as ReactDOM from "react-dom";
import { createStore, compose, applyMiddleware } from "redux";
import { Provider } from "react-redux";
import thunkMiddleware from "redux-thunk";
import createSagaMiddleware from "redux-saga";
import createHistory from "history/createBrowserHistory";
import Path from "path-parser";
import { routerMiddleware } from "react-router-redux";
import { persistStore, persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";
import { PersistGate } from "redux-persist/integration/react";

import rootReducer from "./reducers";
import EnhancedWindow from "./interfaces/EnhancedWindow";
import {{.AppName}} from "./containers/{{.AppName}}";
import rootSaga from "./sagas";

const some = require("lodash/some");

const initializeApp = (element: HTMLElement) => {
  const history = createHistory();
  const sagaMiddleware = createSagaMiddleware();
  const appRouterMiddleware = routerMiddleware(history);
  const middleware = [sagaMiddleware, thunkMiddleware, appRouterMiddleware];

  if (process.env.NODE_ENV !== "production") {
    const { logger } = require("redux-logger");
    middleware.push(logger);
  }

  const persistConfig = {
    key: "root",
    storage,
  };
  const persistedReducer = persistReducer(persistConfig, rootReducer);

  const composeEnhancers = (window as EnhancedWindow).__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
  const appStore = createStore(
    persistedReducer,
    composeEnhancers(applyMiddleware(...middleware)),
  );
  const persistor = persistStore(appStore);
  sagaMiddleware.run(rootSaga);

  ReactDOM.render(
    <Provider store={ appStore }>
      <PersistGate loading={null} persistor={persistor}>
        <{{.AppName}} history={ history } />
      </PersistGate>
    </Provider>,
    element,
  );
};

document.addEventListener("DOMContentLoaded", (event) => {
  const element = document.getElementById("app");
  if (element) {
    initializeApp(element);
  }
});
`
