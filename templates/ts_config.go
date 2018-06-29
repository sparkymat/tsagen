package templates

const TsConfigTemplate = `{
	"compilerOptions": {
	  "jsx":                    "react",
	  "module":                 "commonjs",
	  "noImplicitAny":          true,
	  "sourceMap":              true,
	  "target":                 "ES6",
	  "experimentalDecorators": true
	},
	"lib": [
	  "es5",
	  "es2015",
	  "dom",
	  "scripthost"
	],
	"files": [
	  "src/index.tsx"
	]
  }
`
