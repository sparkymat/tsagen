package templates

const SrcPageReducerTemplate = `import Action from "../interfaces/Action";

export interface {{.PageName}}State {
}

const {{.PageName}}DefaultState: {{.PageName}}State = {
};

const reducer = (state: {{.PageName}}State = {{.PageName}}DefaultState, action: Action): {{.PageName}}State => {
	switch(action.type) {
	default:
		return state;
	}
}

export default reducer;
`
