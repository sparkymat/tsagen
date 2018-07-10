package templates

const SrcAppContainerTemplate = `import { connect } from "react-redux";
import {{.AppName}}View, { {{.AppName}}Props } from "../components/{{.AppName}}";
import { {{.AppName}}State } from "../reducers";
import { onAppInit } from "../actions";

const ReactLifecycleComponent = require("react-lifecycle-component");

interface StateProps {
}

const mapStateToProps = (state: {{.AppName}}State): StateProps => ({
});

interface DispatchProps {
	onAppInit(): void;
}

const mapDispatchToProps = (dispatch: any): DispatchProps => ({
	onAppInit: () => dispatch( onAppInit() ),
});

interface LifecycleProps {
  componentDidMount(): void;
}

const mergeProps = (stateProps: StateProps, dispatchProps: DispatchProps): {{.AppName}}Props & LifecycleProps => ({
	componentDidMount: dispatchProps.onAppInit,
});

const {{.AppName}} = ReactLifecycleComponent.connectWithLifecycle(
	mapStateToProps,
	mapDispatchToProps,
	mergeProps,
)({{.AppName}}View);

export default {{.AppName}};
`
