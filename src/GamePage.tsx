import * as stylex from "@stylexjs/stylex";
import { GamePlayer, PromoBanner, Sidebar, SocialPanel } from "./components";
import { xoArena } from "./catalog";
import { s } from "./styles.stylex";

export default function GamePage() {
  const game = xoArena;
  return <main {...stylex.props(s.container)}>
    <PromoBanner />
    <div {...stylex.props(s.titleRow)}>
      <div><div {...stylex.props(s.eyebrow)}>Featured game · by <a href={game.creator.url}>{game.creator.name}</a></div><h1 {...stylex.props(s.title)}>{game.title}</h1><div {...stylex.props(s.subtitle)}>{game.genres.join(" · ")} · Multiplayer · HTML5 · Free</div></div>
      <span {...stylex.props(s.livePill)}>● PLAYABLE NOW</span>
    </div>
    <div {...stylex.props(s.layout)}>
      <div id="player"><GamePlayer /><SocialPanel slug={game.slug}/></div>
      <Sidebar game={game}/>
    </div>
    <footer {...stylex.props(s.footer)}><span>© 2026 NGG.GG · Built for creators.</span><span>Games · Community · Privacy · Terms</span></footer>
  </main>;
}
