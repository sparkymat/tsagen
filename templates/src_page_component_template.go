package templates

const SrcPageComponentTemplate = `import * as React from "react";

export interface {{.PageName}}Props {
}

const {{.PageName}} = (props: {{.PageName}}Props) => (
	<div>
	</div>
);

export default {{.PageName}};
`
