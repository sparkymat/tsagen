package templates

const PackageJSONTemplate = `{
  "name": "{{.AppName}}",
  "version": "1.0.0",
  "description": "",
  "main": "src/index.tsx",
  "scripts": {
    "build": "webpack",
    "release": "NODE_ENV=production webpack",
    "watch": "webpack --watch",
    "lint": "tslint -p ."
  },
  "author": "",
  "license": "ISC",
  "devDependencies": {
    "@types/core-js": "^0.9.45",
    "@types/history": "^4.6.2",
    "@types/node": "^9.6.1",
    "@types/react": "^16.3.0",
    "@types/react-dom": "^16.0.3",
    "@types/react-redux": "^5.0.14",
    "@types/react-router-dom": "^4.2.5",
    "@types/react-router-redux": "^5.0.13",
    "@types/systemjs": "^0.20.6",
    "clean-webpack-plugin": "^0.1.17",
    "css-loader": "^0.28.9",
		"html-webpack-plugin": "^3.2.0",
    "hashed-module-id-plugin": "^1.0.1",
    "style-loader": "^0.19.1",
    "ts-loader": "^3.2.0",
    "typescript": "^2.6.2",
    "uglifyjs-webpack-plugin": "^1.1.6",
    "webpack": "^3.10.0"
  },
  "dependencies": {
    "axios": "^0.18.0",
    "history": "^4.7.2",
    "inflect": "^0.4.0",
    "lodash": "^4.17.5",
    "moment": "^2.22.1",
    "npm": "^5.8.0",
    "path-parser": "^4.0.3",
    "react": "^16.4.0",
    "react-dom": "^16.2.0",
    "react-lifecycle-component": "^3.0.0",
    "react-redux": "^5.0.6",
    "react-router-dom": "^4.2.2",
    "react-router-redux": "^5.0.0-alpha.9",
    "reactstrap": "^5.0.0",
    "redux": "^3.7.2",
    "redux-logger": "^3.0.6",
    "redux-persist": "^5.9.1",
    "redux-saga": "^0.16.0",
    "redux-thunk": "^2.2.0",
    "serializer.ts": "0.0.12"
  }
}
`
