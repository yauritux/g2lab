import React from "react";
import ReactDOM from "react-dom";
import "./index.css";
//import App from './App';
import Counters from "./components/counters";
import "bootstrap/dist/css/bootstrap.css";
import registerServiceWorker from "./registerServiceWorker";

ReactDOM.render(<Counters />, document.getElementById("root"));
registerServiceWorker();
