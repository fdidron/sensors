import React from 'react'
import ReactDOM from 'react-dom'
import { store, context } from "./store";
import App from "./app";
import "./index.scss";

const Provider = context.Provider;

ReactDOM.render(
  <React.StrictMode>
    <Provider value={store}>
      <App />
    </Provider>,
  </React.StrictMode>,
  document.getElementById('root')
)
