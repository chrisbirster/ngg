import { Header } from "./components";
import GamePage from "./GamePage";
import * as stylex from "@stylexjs/stylex";
import { s } from "./styles.stylex";

export default function App() {
  return <div {...stylex.attrs(s.app)}><Header /><GamePage /></div>;
}
