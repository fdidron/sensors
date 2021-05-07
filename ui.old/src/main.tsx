import { render } from "preact";

import { store, context } from "./store";
import App from "./app";
import "./index.scss";

const Provider = context.Provider;

render(
  <Provider value={store}>
    <App />
  </Provider>,
  document.getElementById("app")!
);
