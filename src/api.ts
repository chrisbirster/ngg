export type Comment = { id: string; gameSlug: string; author: string; body: string; score: number; createdAt: string };
export type GameSnapshot = { likes: number; comments: Comment[] };

async function post<T>(path: string, body: unknown): Promise<T> {
  const response = await fetch(path, { method: "POST", headers: { "Content-Type": "application/json" }, body: JSON.stringify(body) });
  if (!response.ok) throw new Error(`request failed: ${response.status}`);
  if (response.status === 204) return undefined as T;
  return response.json() as Promise<T>;
}
export const socialAPI = {
  like: (slug: string) => post<{likes:number}>(`/api/v1/games/${slug}/reactions`, { type: "like" }),
  comment: (slug: string, body: string) => post<Comment>(`/api/v1/games/${slug}/comments`, { author: "guest-coach", body }),
  playlist: (slug: string) => post<{saved:boolean}>(`/api/v1/games/${slug}/playlists`, { name: "play-later" }),
  report: (slug: string) => post<void>(`/api/v1/games/${slug}/reports`, { reason: "user-submitted" })
};
