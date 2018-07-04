package templates

const SrcRootReducerTemplate = `import { combineReducers } from "redux";

export interface {{.AppName}}State {
}

const rootReducer = combineReducers({
});

export default rootReducer;
`
