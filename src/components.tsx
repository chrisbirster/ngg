import { For, Show, createSignal } from "solid-js";
import * as stylex from "@stylexjs/stylex";
import { socialAPI, type Comment } from "./api";
import { relatedGames, type GameManifest } from "./catalog";
import { s } from "./styles.stylex";

export function Header() {
  return <header {...stylex.props(s.top)}>
    <div {...stylex.props(s.nav)}>
      <a href="/" {...stylex.props(s.logo)}>NGG.GG</a>
      <nav {...stylex.props(s.mainNav)}>
        <For each={["Games","Movies","Audio","Art","Community"]}>{item => <a href="#" {...stylex.props(s.navLink)}>{item}</a>}</For>
      </nav>
      <input aria-label="Search NGG" placeholder="Search games…" {...stylex.props(s.search)} />
      <div {...stylex.props(s.auth)}><a href="#">Log in</a><a href="#" {...stylex.props(s.signup)}>Sign up</a></div>
    </div>
    <div {...stylex.props(s.categories)}><div {...stylex.props(s.categoryInner)}>
      <For each={["Featured","New Games","Popular","Action","Sports","Strategy","Multiplayer","Game Jams"]}>{item => <a href="#">{item}</a>}</For>
    </div></div>
  </header>;
}

export function PromoBanner() {
  return <section {...stylex.props(s.promo)}>
    <div><div {...stylex.props(s.promoTitle)}>SMALL GAMES.<br />BIG CREATORS.</div><div {...stylex.props(s.promoSub)}>INDIE GAMES FOREVER · PLAY SOMETHING WEIRD</div></div>
    <div aria-hidden="true" {...stylex.props(s.promoArt)}><span>👾</span><span>🏈</span><span>🐝</span><span>🔥</span></div>
  </section>;
}

const initialOffense = [[14,26],[22,34],[31,42],[40,46],[48,48],[56,46],[65,42],[76,34],[86,26],[39,62],[48,71]];
const initialDefense = [[18,70],[27,65],[36,61],[45,59],[54,59],[63,61],[72,65],[82,70],[33,80],[50,82],[68,80]];

export function GamePlayer() {
  const [snap, setSnap] = createSignal(false);
  const [score, setScore] = createSignal([7,3]);
  const [clock, setClock] = createSignal("08:42");
  const [plays, setPlays] = createSignal(["1st & 10 · Pass, 8 yards","2nd & 2 · Run, 3 yards","1st & 10 · Sack, -6 yards","2nd & 16 · Incomplete"]);
  const [side, setSide] = createSignal("Offense");
  const run = () => {
    const next = !snap(); setSnap(next);
    setClock(next ? "08:35" : "08:28");
    setPlays(items => [next ? "3rd & 7 · PA Cross, 18 yards" : "1st & 10 · Inside zone, 5 yards", ...items].slice(0,4));
    if (!next) setScore([14,3]);
  };
  return <div {...stylex.props(s.player)}>
    <div {...stylex.props(s.scoreboard)}>
      <div><div {...stylex.props(s.team)}>TEAM X</div><div {...stylex.props(s.clock)}>HARRISBURG</div></div>
      <div><div {...stylex.props(s.score)}>{score()[0]} — {score()[1]}</div><div {...stylex.props(s.clock)}>Q2 · {clock()}</div></div>
      <div><div {...stylex.props(s.teamRight)}>TEAM O</div><div {...stylex.props(s.clock)}>BALTIMORE</div></div>
    </div>
    <div {...stylex.props(s.field)}>
      <div {...stylex.props(s.midfield)} />
      <div {...stylex.props(s.situation)}>3RD & 7 · X 43</div>
      <svg viewBox="0 0 100 100" preserveAspectRatio="none" style={{position:"absolute",inset:"0",width:"100%",height:"100%",opacity:snap()?".25":".8"}}>
        <path d="M 31 42 C 42 20, 58 20, 70 38" fill="none" stroke="#ffd633" stroke-width="1.2" stroke-dasharray="2 1"/>
        <path d="M 65 42 C 74 40, 82 33, 89 20" fill="none" stroke="#ffd633" stroke-width="1.2"/>
      </svg>
      <For each={initialOffense}>{(p,i) => <div {...stylex.props(s.marker,s.offense)} style={{left:`${p[0]+(snap()&&i()%3===0?9:0)}%`,top:`${p[1]-(snap()&&i()<9?8:0)}%`}}>X</div>}</For>
      <For each={initialDefense}>{(p,i) => <div {...stylex.props(s.marker,s.defense)} style={{left:`${p[0]-(snap()&&i()%2===0?5:0)}%`,top:`${p[1]-(snap()?9:0)}%`}}>O</div>}</For>
      <div {...stylex.props(s.ball)} style={{left:snap()?"60%":"48%",top:snap()?"34%":"58%"}} />
    </div>
    <div {...stylex.props(s.controls)}>
      <div {...stylex.props(s.playControls)}>
        <div {...stylex.props(s.tabs)}><For each={["Offense","Defense","Special Teams"]}>{item => <button onClick={() => setSide(item)} {...stylex.props(s.tab,side()===item&&s.activeTab)}>{item}</button>}</For></div>
        <div {...stylex.props(s.selectRow)}>
          <label {...stylex.props(s.label)}>Formation<select {...stylex.props(s.select)}><option>Shotgun</option><option>Singleback</option><option>Pistol</option></select></label>
          <label {...stylex.props(s.label)}>Play<select {...stylex.props(s.select)}><option>PA Cross</option><option>Inside Zone</option><option>Four Verticals</option></select></label>
        </div>
        <div {...stylex.props(s.playName)}>PA Cross</div><p {...stylex.props(s.playDesc)}>Play action with crossing routes that punish an aggressive second level.</p>
        <button onClick={run} {...stylex.props(s.primary)}>{snap() ? "NEXT SNAP ▶" : "RUN PLAY ▶"}</button>
      </div>
      <div {...stylex.props(s.recent)}><h3 {...stylex.props(s.sectionLabel)}>RECENT PLAYS</h3><For each={plays()}>{play => { const [down,result]=play.split(" · "); return <div {...stylex.props(s.playItem)}><strong>{down}</strong><span>{result}</span></div>; }}</For></div>
    </div>
  </div>;
}

export function Sidebar(props:{game:GameManifest}) {
  return <aside {...stylex.props(s.sticky)}>
    <section {...stylex.props(s.card)}><div {...stylex.props(s.cover)}>X/O</div><div {...stylex.props(s.panelPad)}>
      <h2 style={{margin:"0 0 7px","font-size":"19px"}}>{props.game.title}</h2><p {...stylex.props(s.description)}>{props.game.description}</p>
      <a href="#player" {...stylex.props(s.primary)} style={{display:"block","text-align":"center"}}>PLAY NOW ▶</a>
      <div {...stylex.props(s.meta)} style={{"margin-top":"16px"}}>
        <div {...stylex.props(s.metaLine)}><span>Published</span><strong>{props.game.publishedAt}</strong></div>
        <div {...stylex.props(s.metaLine)}><span>Updated</span><strong>{props.game.updatedAt}</strong></div>
        <div {...stylex.props(s.metaLine)}><span>Runtime</span><strong>Web · No download</strong></div>
        <div {...stylex.props(s.tags)}><For each={props.game.tags}>{tag => <span {...stylex.props(s.tag)}>#{tag}</span>}</For></div>
      </div>
    </div></section>
    <section {...stylex.props(s.card,s.panelPad)}><h3 {...stylex.props(s.sectionLabel)}>YOU MIGHT ALSO LIKE</h3><For each={relatedGames}>{game => <a href="#" {...stylex.props(s.relatedGrid)}><div {...stylex.props(s.thumb)} style={{"background-color":game.color}}>{game.glyph}</div><div><strong style={{"font-size":"12px"}}>{game.title}</strong><div style={{color:"#8d99aa","font-size":"10px","margin-top":"3px"}}>{game.genre}</div></div></a>}</For></section>
  </aside>;
}

const seedComments: Comment[] = [
  {id:"c1",gameSlug:"xo-arena-football",author:"pixelcoach",body:"The X/O presentation makes play calling instantly readable.",score:42,createdAt:"2026-09-14"},
  {id:"c2",gameSlug:"xo-arena-football",author:"fourthandone",body:"Give me one more drive. Then another.",score:27,createdAt:"2026-09-14"}
];

export function SocialPanel(props:{slug:string}) {
  const [likes,setLikes]=createSignal(2400), [liked,setLiked]=createSignal(false), [saved,setSaved]=createSignal(false);
  const [comments,setComments]=createSignal(seedComments), [body,setBody]=createSignal("");
  const like=async()=>{ if(liked()) return; setLiked(true); setLikes(v=>v+1); try{const x=await socialAPI.like(props.slug);setLikes(x.likes)}catch{} };
  const submit=async(e:SubmitEvent)=>{e.preventDefault(); const value=body().trim(); if(!value)return; setBody(""); const optimistic:Comment={id:`local-${Date.now()}`,gameSlug:props.slug,author:"guest-coach",body:value,score:0,createdAt:new Date().toISOString()}; setComments(c=>[optimistic,...c]); try{const saved=await socialAPI.comment(props.slug,value);setComments(c=>[saved,...c.filter(x=>x.id!==optimistic.id)])}catch{}};
  return <section {...stylex.props(s.card,s.social)}>
    <div {...stylex.props(s.actionRow)}>
      <button onClick={like} {...stylex.props(s.action)}>{liked()?"♥":"♡"} {likes().toLocaleString()}</button>
      <a href="#comments" {...stylex.props(s.action)}>💬 {comments().length}</a>
      <button onClick={()=>navigator.clipboard?.writeText(location.href)} {...stylex.props(s.action)}>Share</button>
      <button onClick={async()=>{setSaved(true);try{await socialAPI.playlist(props.slug)}catch{}}} {...stylex.props(s.action)}>{saved()?"✓ Saved":"＋ Playlist"}</button>
      <button onClick={()=>socialAPI.report(props.slug).catch(()=>{})} {...stylex.props(s.action)}>Report</button>
    </div>
    <form onSubmit={submit} {...stylex.props(s.commentForm)}><input value={body()} onInput={e=>setBody(e.currentTarget.value)} maxlength={1000} placeholder="Tell the creator what you think…" {...stylex.props(s.input)}/><button {...stylex.props(s.primary)} style={{width:"auto"}}>Post</button></form>
    <div id="comments" {...stylex.props(s.comments)}><For each={comments()}>{comment=><article {...stylex.props(s.comment)}><div {...stylex.props(s.commentHead)}><span>@{comment.author}</span><span>▲ {comment.score}</span></div><p {...stylex.props(s.commentBody)}>{comment.body}</p></article>}</For></div>
  </section>;
}
