package templates

const SrcInterfaceAction = `import ServerConfig from "../models/ServerConfig";

interface Action {
	type: string;
}

export interface ConfigAction extends Action {
	serverConfig: ServerConfig;
}

export default Action;
`
