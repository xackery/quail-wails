# README

## About

This is the official Wails Vue-TS template.

You can configure the project by editing `wails.json`. More information about the project settings can be found
here: https://wails.io/docs/reference/project-config

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.

## Commands ran for setup
- `cd frontend && npm install vue-router && cd ..`
- `cd frontend && npm install @shoelace-style/shoelace && cd ..` ([ref](https://shoelace.style/frameworks/vue))
- `cd frontend && npm install -D vite-plugin-static-copy && cd ..` ([ref](https://stackoverflow.com/questions/76446464/how-to-expose-shoelace-icons-in-a-vite-app))
- `cd frontend && npm install --save three && cd ..` ([ref](https://threejs.org/docs/index.html#manual/en/introduction/Installation))
- `cd frontend && npm i --save-dev @types/three && cd ..` ([ref](https://stackoverflow.com/questions/63492296/how-to-import-three-js-in-vue-js))