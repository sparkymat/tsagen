package templates

const SrcAppComponentTemplate = `import * as React from "react";

export interface {{.AppName}}Props {
}

const {{.AppName}} = (props: {{.AppName}}Props) => (
	<div>
	</div>
);

export default {{.AppName}};
`
