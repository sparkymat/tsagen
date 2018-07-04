package templates

const SrcActionIndex = `import Action, { ConfigAction } from "../interfaces/Action";
import ServerConfig from "../models/ServerConfig";

export const ON_APP_INIT = "ON_APP_INIT";
export const onAppInit = (): Action => ({
  type: ON_APP_INIT,
});

export const FETCH_CONFIG = "FETCH_CONFIG";
export const fetchConfig = (): Action => ({
  type: FETCH_CONFIG,
});

export const CONFIG_FETCH_STARTED = "CONFIG_FETCH_STARTED";
export const configFetchStarted = (): Action => ({
  type: CONFIG_FETCH_STARTED,
});

export const CONFIG_FETCH_SUCCEEDED = "CONFIG_FETCH_SUCCEEDED";
export const configFetchSucceeded = (serverConfig: ServerConfig): ConfigAction => ({
  serverConfig: serverConfig,
  type:         CONFIG_FETCH_SUCCEEDED,
});

export const CONFIG_FETCH_FAILED = "CONFIG_FETCH_FAILED";
export const configFetchFailed = (): Action => ({
  type: CONFIG_FETCH_FAILED,
});
`
