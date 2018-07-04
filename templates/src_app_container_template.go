package templates

const SrcAppContainerTemplate = `import { connect } from "react-redux";
import {{.AppName}}View, { {{.AppName}}Props } from "../components/{{.AppName}}";
import { {{.AppName}}State } from "../reducers";

interface StateProps {
}

const mapStateToProps = (state: {{.AppName}}State) => ({
});

interface DispatchProps {
}

const mapDispatchToProps = (dipatch: any) => ({
});

const mergeProps = (stateProps: StateProps, dispatchProps: DispatchProps): {{.AppName}}Props => ({
});

const {{.AppName}} = connect(
	mapStateToProps,
	mapDispatchToProps,
	mergeProps,
)({{.AppName}}View);

export default {{.AppName}};
`
