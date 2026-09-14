import { render } from "@solidjs/web";
import { Router } from "@solidjs/router";
import App from "./App";
import "./reset.css";

render(() => <Router><App /></Router>, document.getElementById("root")!);
