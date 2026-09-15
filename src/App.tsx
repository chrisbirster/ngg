import { Header } from "./components";
import GamePage from "./GamePage";
import PortalPage from "./PortalPage";
import * as stylex from "@stylexjs/stylex";
import { s } from "./styles.stylex";

export default function App() {
  const path = window.location.pathname;
  const gameDetail = path === "/" || path.startsWith("/games/xo-arena-football") || path.startsWith("/play/xo-arena-football");
  return <div {...stylex.attrs(s.app)}><Header />{gameDetail ? <GamePage /> : <PortalPage path={path} />}</div>;
}
