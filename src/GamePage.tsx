import * as stylex from "@stylexjs/stylex";
import { GameHeading, GamePlayer, PromoBanner, Sidebar, SocialPanel } from "./components";
import { xoArena } from "./catalog";
import { s } from "./styles.stylex";

export default function GamePage() {
  const game = xoArena;
  return <main {...stylex.attrs(s.container)}>
    <PromoBanner />
    <div {...stylex.attrs(s.layout)}>
      <section id="player" {...stylex.attrs(s.gameCard)}><GameHeading game={game} /><GamePlayer /></section>
      <Sidebar game={game}/>
      <SocialPanel slug={game.slug}/>
    </div>
    <footer {...stylex.attrs(s.footer)}><span>© 2026 NGG.GG · Built for creators.</span><span>Games · Community · Privacy · Terms</span></footer>
  </main>;
}
