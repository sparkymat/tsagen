package templates

const SrcPageContainerTemplate = `import { connect } from "react-redux";
import {{.PageName}}View, { {{.PageName}}Props } from "./components";
import { {{.AppName}}State } from "../reducers";

interface StateProps {
}

const mapStateToProps = (state: {{.AppName}}State): StateProps => ({
});

interface DispatchProps {
}

const mapDispatchToProps = (dispatch: any): DispatchProps => ({
});

const mergeProps = (stateProps: StateProps, dispatchProps: DispatchProps): {{.PageName}}Props => ({
});

const {{.PageName}} = connect(
	mapStateToProps,
	mapDispatchToProps,
	mergeProps,
)({{.PageName}}View);

export default {{.PageName}};
`
