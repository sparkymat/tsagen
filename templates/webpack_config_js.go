package templates

const WebpackConfigTemplate = `const webpack = require("webpack");
const path = require("path");

const HashedModuleIdsPlugin = require("hashed-module-id-plugin");
const CleanWebpackPlugin = require("clean-webpack-plugin");
const HtmlWebpackPlugin = require("html-webpack-plugin");
const UglifyJsPlugin = require("uglifyjs-webpack-plugin");

const isProd = process.env.NODE_ENV === "production" ? true : false;

var plugins = [
	new HtmlWebpackPlugin({title: "{{.AppTitle}}", template: "src/html/index.html"})
];

plugins.push(new CleanWebpackPlugin(["dist/js/"], {allowExternal: true}));
if (isProd) {
    plugins.push(new UglifyJsPlugin());
    plugins.push(new HashedModuleIdsPlugin());
}

const config = {
  /*
   * app.ts represents the entry point to your web application. Webpack will
   * recursively go through every "require" statement in app.ts and
   * efficiently build out the application's dependency tree.
   */
  entry: ["./src/index.tsx"],

  /*
   * The combination of path and filename tells Webpack what name to give to
   * the final bundled JavaScript file and where to store this file.
   */
  output: {
    path: path.resolve(__dirname, "dist", "js"),
    filename: isProd ? "{{.JsFileName}}-[chunkhash].js" : "{{.JsFileName}}-development.js"
  },

  /*
   * resolve lets Webpack now in advance what file extensions you plan on
   * "require"ing into the web application, and allows you to drop them
   * in your code.
   */
  resolve: {
    extensions: [".ts", ".tsx", ".js"]
  },

  devtool: "#cheap-module-source-map",

  module: {
    /*
     * Each loader needs an associated Regex test that goes through each
     * of the files you've included (or in this case, all files but the
     * ones in the excluded directories) and finds all files that pass
     * the test. Then it will apply the loader to that file. I haven't
     * installed ts-loader yet, but will do that shortly.
     */
    loaders: [
      {
        test: /\.tsx?$/,
        loader: "ts-loader",
        exclude: /node_modules/
      },
      {
        test: /\.css$/,
        use: [ 'style-loader', 'css-loader' ]
      }
    ]
  },

  plugins: plugins
};

module.exports = config;
`
