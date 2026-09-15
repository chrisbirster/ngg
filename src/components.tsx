import { For, createSignal } from "solid-js";
import * as stylex from "@stylexjs/stylex";
import { socialAPI, type Comment } from "./api";
import { relatedGames, type GameManifest } from "./catalog";
import { s } from "./styles.stylex";

export function Header() {
  const primary = ["Games", "Movies", "Audio", "Art", "Community"];
  const categories = ["Featured", "Top Rated", "New", "Action", "Adventure", "Puzzle", "Racing", "Sports", "Multiplayer", "More⌄"];
  return <header {...stylex.attrs(s.top)}>
    <div {...stylex.attrs(s.nav)}>
      <a href="/" {...stylex.attrs(s.logoWrap)} aria-label="NGG.GG home">
        <span {...stylex.attrs(s.logo)}><span {...stylex.attrs(s.logoGold)}>NGG</span><span {...stylex.attrs(s.logoWhite)}>.GG</span></span>
        <span {...stylex.attrs(s.logoTag)}>PLAY CREATE SHARE</span>
      </a>
      <nav {...stylex.attrs(s.mainNav)} aria-label="Primary navigation">
        <For each={primary}>{(item, index) => <a href="#" {...stylex.attrs(s.navLink, index() === 0 && s.activeNav)}>{item}</a>}</For>
      </nav>
      <label {...stylex.attrs(s.searchWrap)}>
        <input aria-label="Search NGG" placeholder="Search games, creators, tags…" {...stylex.attrs(s.search)} />
        <span aria-hidden="true" {...stylex.attrs(s.searchIcon)}>⌕</span>
      </label>
      <div {...stylex.attrs(s.auth)}><a href="#">Log In</a><a href="#" {...stylex.attrs(s.signup)}>Sign Up</a></div>
    </div>
    <nav {...stylex.attrs(s.categories)} aria-label="Game categories"><div {...stylex.attrs(s.categoryInner)}>
      <For each={categories}>{(item, index) => <a href="#" {...stylex.attrs(index() === 0 && s.categoryActive)}>{item}</a>}</For>
    </div></nav>
  </header>;
}

export function PromoBanner() {
  return <section {...stylex.attrs(s.promo)} aria-label="Small games, big creators">
    <div {...stylex.attrs(s.promoShade)} />
    <div {...stylex.attrs(s.promoCopy)}>
      <div {...stylex.attrs(s.promoTitle)}>SMALL GAMES.<br />BIG CREATORS.</div>
      <div {...stylex.attrs(s.promoRight)}>INDIE GAMES<br />FOREVER ♛</div>
    </div>
  </section>;
}

export function GameHeading(props: { game: GameManifest }) {
  return <>
    <div {...stylex.attrs(s.gameHeading)}>
      <div aria-hidden="true" {...stylex.attrs(s.gameIcon)}>🏈</div>
      <div><h1 {...stylex.attrs(s.gameTitle)}>{props.game.title}</h1><div {...stylex.attrs(s.byline)}>by {props.game.creator.name}</div></div>
      <div {...stylex.attrs(s.genreTags)}>
        <For each={[...props.game.genres, "Multiplayer", "HTML5", "Free"]}>{tag => <span {...stylex.attrs(s.genreTag)}>{tag}</span>}</For>
      </div>
    </div>
    <p {...stylex.attrs(s.gameLead)}>Call the plays. Coach your team. Build a dynasty. A browser-based football strategy universe.</p>
  </>;
}

const offense = [[36,35],[31,43],[36,50],[31,58],[37,65],[43,48],[43,59],[42,72],[48,54],[47,65],[50,58]];
const defense = [[56,35],[56,48],[60,42],[64,36],[60,58],[68,50],[62,68],[54,72],[69,64],[58,80],[67,80]];
const yardNumbers = [
  [18,"10"],[27,"20"],[36,"30"],[45,"40"],[54,"50"],[63,"40"],[72,"30"],[81,"20"],
] as const;

function MiniPlayDiagram() {
  return <div {...stylex.attrs(s.miniPlay)} aria-label="PA Cross play diagram">
    <svg viewBox="0 0 110 90" width="100%" height="100%" aria-hidden="true">
      <path d="M18 70 C24 48 34 34 42 25" stroke="#54b9f8" stroke-width="2" fill="none" />
      <path d="M89 72 L89 28 L68 28" stroke="#ffc928" stroke-width="2" fill="none" />
      <path d="M55 70 C55 54 58 42 70 35" stroke="#ff5c5c" stroke-width="2" fill="none" />
      <g fill="#58b9f5"><circle cx="18" cy="70" r="5"/><circle cx="31" cy="74" r="5"/><circle cx="43" cy="77" r="5"/><circle cx="55" cy="77" r="5"/><circle cx="68" cy="77" r="5"/><circle cx="82" cy="74" r="5"/></g>
      <circle cx="55" cy="57" r="5" fill="#fff"/><circle cx="55" cy="57" r="2" fill="#ff5656"/>
    </svg>
  </div>;
}

export function GamePlayer() {
  const [snap, setSnap] = createSignal(false);
  const [score, setScore] = createSignal([14, 10]);
  const [clock, setClock] = createSignal("6:42");
  const [plays, setPlays] = createSignal(["1st & 10 · Pass, 8 yards", "2nd & 2 · Run, 3 yards", "1st & 10 · Sack, -6 yards", "3rd & 7 · Incomplete", "2nd & 7 · Run, 4 yards"]);
  const [side, setSide] = createSignal("Offense");
  const run = () => {
    const next = !snap(); setSnap(next); setClock(next ? "6:35" : "6:28");
    setPlays(items => [next ? "2nd & 7 · PA Cross, 18 yards" : "1st & 10 · Inside zone, 5 yards", ...items].slice(0, 5));
    if (!next) setScore([21, 10]);
  };
  return <div {...stylex.attrs(s.player)}>
    <div {...stylex.attrs(s.scoreboard)}>
      <div {...stylex.attrs(s.teamBlock)}><span {...stylex.attrs(s.teamSymbolX)}>X</span><div><div {...stylex.attrs(s.team)}>TEAM X</div><div {...stylex.attrs(s.record)}>3 - 2</div></div></div>
      <div {...stylex.attrs(s.score)}>{score()[0]}</div>
      <div {...stylex.attrs(s.gameClock)}>Q2&nbsp; {clock()}<div {...stylex.attrs(s.down)}>2nd &amp; 7</div></div>
      <div {...stylex.attrs(s.score)}>{score()[1]}</div>
      <div {...stylex.attrs(s.teamRightBlock)}><span {...stylex.attrs(s.teamSymbolO)} /><div><div {...stylex.attrs(s.team)}>TEAM O</div><div {...stylex.attrs(s.record)}>4 - 1</div></div></div>
    </div>
    <div {...stylex.attrs(s.field)}>
      <div {...stylex.attrs(s.endzoneX)}><span {...stylex.attrs(s.endzoneText)}>TEAM X</span></div>
      <div {...stylex.attrs(s.endzoneO)}><span {...stylex.attrs(s.endzoneText)}>TEAM O</span></div>
      <div {...stylex.attrs(s.hashesTop)} /><div {...stylex.attrs(s.hashesMid)} />
      <For each={yardNumbers}>{item => <span {...stylex.attrs(s.yardNumber)} style={{left:`${item[0]}%`}}>{item[1]}</span>}</For>
      <svg viewBox="0 0 100 100" preserveAspectRatio="none" style={{position:"absolute",inset:"0",width:"100%",height:"100%",opacity:snap()?".2":".9","z-index":"1"}} aria-hidden="true">
        <defs><marker id="route-arrow" markerWidth="5" markerHeight="5" refX="4" refY="2.5" orient="auto"><path d="M0,0 L5,2.5 L0,5" fill="none" stroke="#55b9ff" stroke-width="1.2"/></marker></defs>
        <path d="M36 35 C42 29,48 29,53 35" fill="none" stroke="#55b9ff" stroke-width=".7" marker-end="url(#route-arrow)"/>
        <path d="M42 72 C45 78,50 80,55 73" fill="none" stroke="#55b9ff" stroke-width=".7" marker-end="url(#route-arrow)"/>
      </svg>
      <For each={offense}>{(p,i) => <div {...stylex.attrs(s.marker,s.offense)} style={{left:`${p[0]+(snap()&&i()%3===0?7:0)}%`,top:`${p[1]-(snap()&&i()<8?6:0)}%`}}>X</div>}</For>
      <For each={defense}>{(p,i) => <div {...stylex.attrs(s.marker,s.defense)} style={{left:`${p[0]-(snap()&&i()%2===0?4:0)}%`,top:`${p[1]-(snap()?5:0)}%`}}>O</div>}</For>
      <div {...stylex.attrs(s.ball)} style={{left:snap()?"55%":"46%",top:snap()?"44%":"55%"}} />
    </div>
    <div {...stylex.attrs(s.controls)}>
      <div {...stylex.attrs(s.playControls)}>
        <div {...stylex.attrs(s.tabs)}><For each={["Offense","Defense","Special Teams"]}>{item => <button onClick={() => setSide(item)} {...stylex.attrs(s.tab,side()===item&&s.activeTab)}>{item}</button>}</For></div>
        <div {...stylex.attrs(s.callGrid)}>
          <label {...stylex.attrs(s.label)}>Formation<select {...stylex.attrs(s.select)}><option>Shotgun</option><option>Singleback</option><option>Pistol</option></select></label>
          <label {...stylex.attrs(s.label)}>Play<select {...stylex.attrs(s.select)}><option>PA Cross</option><option>Inside Zone</option><option>Four Verticals</option></select></label>
          <MiniPlayDiagram />
          <div {...stylex.attrs(s.callInfo)}><div><div {...stylex.attrs(s.playName)}>PA Cross</div><p {...stylex.attrs(s.playDesc)}>Play action with crossing routes.</p></div><button onClick={run} {...stylex.attrs(s.primary)}>{snap() ? "NEXT SNAP ▶" : "RUN PLAY ▶"}</button></div>
        </div>
      </div>
      <div {...stylex.attrs(s.recent)}><h3 {...stylex.attrs(s.sectionLabel)}>RECENT PLAYS</h3><For each={plays()}>{(play,index) => { const [down,result]=play.split(" · "); return <div {...stylex.attrs(s.playItem)}><span>{down}</span><span {...stylex.attrs(index()%3===2?s.playTeamO:s.playTeamX)}>{index()%3===2?"":"X"}</span><span>{result}</span></div>; }}</For></div>
    </div>
  </div>;
}

export function Sidebar(props:{game:GameManifest}) {
  const more = [["XO ARENA","XO Arena: Franchise","Simulation"],["X ↗ O","XO Arena: Playbooks","Strategy"],["X  vs  O","XO Arena: Quick Match","Sports"]];
  return <aside {...stylex.attrs(s.sticky)}>
    <section {...stylex.attrs(s.card)}>
      <div {...stylex.attrs(s.cover)}><img src="/assets/xo-arena-cover.webp" alt="" {...stylex.attrs(s.coverImage)} /><div {...stylex.attrs(s.coverShade)}><div {...stylex.attrs(s.coverTitle)}>XO ARENA<div {...stylex.attrs(s.coverSub)}>FOOTBALL</div></div></div></div>
      <div {...stylex.attrs(s.panelPad)}><p {...stylex.attrs(s.sideDescription)}>XO Arena Football is a free browser-based football strategy game where you call plays, coach your team, or run a full franchise.</p><a href="#player" {...stylex.attrs(s.primary)} style={{display:"block","text-align":"center"}}>Play Now</a></div>
    </section>
    <section {...stylex.attrs(s.card,s.sideSection)}>
      <h2 {...stylex.attrs(s.sideTitle)}>More from XO Arena</h2>
      <For each={more}>{item => <a href="#" {...stylex.attrs(s.moreItem)}><span {...stylex.attrs(s.moreThumb)}>{item[0]}</span><span><span {...stylex.attrs(s.moreName)}>{item[1]}</span><span {...stylex.attrs(s.moreGenre)} style={{display:"block"}}>{item[2]}</span></span></a>}</For>
      <div {...stylex.attrs(s.meta)}>
        <div {...stylex.attrs(s.metaLine)}><span>Published</span><span {...stylex.attrs(s.metaValue)}>{props.game.publishedAt}</span></div>
        <div {...stylex.attrs(s.metaLine)}><span>Updated</span><span {...stylex.attrs(s.metaValue)}>{props.game.updatedAt}</span></div>
        <div {...stylex.attrs(s.metaLine)}><span>Game Size</span><span {...stylex.attrs(s.metaValue)}>Web (No Download)</span></div>
        <div {...stylex.attrs(s.metaLine)}><span>Platform</span><span {...stylex.attrs(s.metaValue)}>Browser (Desktop &amp; Mobile)</span></div>
        <div {...stylex.attrs(s.metaLine)}><span>Tags</span><span {...stylex.attrs(s.metaValue)}>Sports, Football, Strategy, Multiplayer</span></div>
        <div {...stylex.attrs(s.metaLine)}><span>Language</span><span {...stylex.attrs(s.metaValue)}>English</span></div>
        <div {...stylex.attrs(s.metaLine)}><span>Developer</span><span {...stylex.attrs(s.metaValue)}>XO Arena</span></div>
      </div>
    </section>
    <section {...stylex.attrs(s.card,s.sideSection)}><h2 {...stylex.attrs(s.sideTitle)}>You might also like</h2><For each={relatedGames.slice(0,2)}>{game => <a href="#" {...stylex.attrs(s.relatedGrid)}><div {...stylex.attrs(s.thumb)} style={{"background-color":game.color}}>{game.glyph}</div><div><strong style={{"font-size":"11px"}}>{game.title}</strong><div {...stylex.attrs(s.moreGenre)}>{game.genre}</div></div></a>}</For></section>
  </aside>;
}

const seedComments: Comment[] = [
  {id:"c1",gameSlug:"xo-arena-football",author:"PixelQB",body:"This is sick. Finally a football game that actually feels good in the browser. The play calling is 🔥",score:124,createdAt:"3 days ago"},
  {id:"c2",gameSlug:"xo-arena-football",author:"GridironGhost",body:"Been waiting for something like this on NGG. Can't wait to see multiplayer leagues!",score:37,createdAt:"2 days ago"},
];

export function SocialPanel(props:{slug:string}) {
  const [likes,setLikes]=createSignal(2400), [liked,setLiked]=createSignal(false), [saved,setSaved]=createSignal(false);
  const [comments,setComments]=createSignal(seedComments), [body,setBody]=createSignal("");
  const like=async()=>{ if(liked()) return; setLiked(true); setLikes(v=>v+1); try{const x=await socialAPI.like(props.slug);setLikes(x.likes)}catch{} };
  const submit=async(e:SubmitEvent)=>{e.preventDefault(); const value=body().trim(); if(!value)return; setBody(""); const optimistic:Comment={id:`local-${Date.now()}`,gameSlug:props.slug,author:"guest-coach",body:value,score:0,createdAt:"now"}; setComments(c=>[optimistic,...c]); try{const saved=await socialAPI.comment(props.slug,value);setComments(c=>[saved,...c.filter(x=>x.id!==optimistic.id)])}catch{}};
  return <div {...stylex.attrs(s.social)}>
    <div {...stylex.attrs(s.actionRow)}>
      <button onClick={like} {...stylex.attrs(s.action)}>{liked()?"♥":"♡"} {likes().toLocaleString()}</button><a href="#comments" {...stylex.attrs(s.action)}>▢ {comments().length + 319}</a><button onClick={()=>navigator.clipboard?.writeText(location.href)} {...stylex.attrs(s.action)}>↗ Share</button><button onClick={async()=>{setSaved(true);try{await socialAPI.playlist(props.slug)}catch{}}} {...stylex.attrs(s.action)}>{saved()?"✓ Saved":"▣ Add to Playlist"}</button><button onClick={()=>socialAPI.report(props.slug).catch(()=>{})} {...stylex.attrs(s.action)}>⚑ Report</button>
    </div>
    <section id="comments" {...stylex.attrs(s.commentsCard)}>
      <h2 {...stylex.attrs(s.sideTitle)}>Comments ({comments().length + 319})</h2>
      <form onSubmit={submit} {...stylex.attrs(s.commentForm)}><input value={body()} onInput={e=>setBody(e.currentTarget.value)} maxlength={1000} placeholder="Join the conversation…" {...stylex.attrs(s.input)}/><button {...stylex.attrs(s.primary)} style={{width:"72px"}}>Post</button></form>
      <div {...stylex.attrs(s.comments)}><For each={comments()}>{comment=><article {...stylex.attrs(s.comment)}><div {...stylex.attrs(s.avatar)}>{comment.author.slice(0,2).toUpperCase()}</div><div><div {...stylex.attrs(s.commentHead)}><span {...stylex.attrs(s.commentAuthor)}>{comment.author}</span><span {...stylex.attrs(s.commentDate)}>{comment.createdAt}</span></div><p {...stylex.attrs(s.commentBody)}>{comment.body}</p><div {...stylex.attrs(s.commentActions)}>♡ {comment.score}&nbsp;&nbsp;&nbsp; Reply</div></div></article>}</For></div>
    </section>
  </div>;
}
