export type Creator = { name: string; handle: string; url: string };
export type GameManifest = {
  slug: string;
  title: string;
  creator: Creator;
  description: string;
  genres: string[];
  tags: string[];
  runtime: "html5" | "wasm";
  multiplayer: boolean;
  mobile: boolean;
  publishedAt: string;
  updatedAt: string;
  gameUrl: string;
  accent: string;
  launchMode: "local" | "embed";
};

export const xoArena: GameManifest = {
  slug: "xo-arena-football",
  title: "XO Arena Football",
  creator: { name: "XO Arena", handle: "xo-arena", url: "https://game.vutadex.com" },
  description: "Call the plays. Coach your team. Build a dynasty.",
  genres: ["Sports", "Strategy", "Simulation"],
  tags: ["football", "multiplayer", "html5", "free"],
  runtime: "html5",
  multiplayer: true,
  mobile: true,
  publishedAt: "Sep 14, 2026",
  updatedAt: "Sep 14, 2026",
  gameUrl: "https://game.vutadex.com",
  accent: "#ffd633",
  launchMode: "local"
};

export const relatedGames = [
  { title: "Fourth & Forever", genre: "Sports", color: "#ec5d41", glyph: "4TH" },
  { title: "Pocket General", genre: "Strategy", color: "#4cc6be", glyph: "PG" },
  { title: "Blitz Board", genre: "Arcade", color: "#8d6bf2", glyph: "BB" }
];
