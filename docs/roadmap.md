# Community platform implementation status

The original ten visual milestones remain complete. The ten community-platform milestones now also have a vertical implementation backed by Turso.

| Milestone | Delivered |
|---|---|
| M0 architecture | Portal/game ownership boundary and route contract |
| M1 global shell | Header, search, primary and category navigation |
| M2 game page | Responsive 3:1 page layout with mobile stacking |
| M3 player | Interactive XO scoreboard, field, markers, routes |
| M4 controls | Play tabs, formation/play selection, recent plays |
| M5 metadata | Typed game manifest and creator/runtime metadata |
| M6 sidebar | CTA, game information, tags, recommendations |
| M7 social | Reactions, comments, playlist, share and report flows |
| M8 promotion | Original NGG promotional banner |
| M9 polish | Responsive breakpoints, sticky sidebar, animation, failure-safe local player |

## Community milestones

1. Accounts and public creator profiles.
2. Game/video drafts, collaboration, credits, and submission.
3. Direct-upload tickets and immutable object keys.
4. HTML5/WASM releases with a sandbox-player boundary.
5. Video releases with a managed-transcoding boundary.
6. Content comments, ratings, favorites, and playlists.
7. Forums, mentions, and notifications.
8. Reports, moderation queue, actions, and an audit event stream.
9. SDK achievements, leaderboards, and versioned cloud saves.
10. Five authoritative instant-game rooms: Tic-Tac-Toe, Four in a Row, Checkers, Chess, and Yacht Dice.

## Next production hardening

1. Replace the development session bootstrap with verified magic-link delivery.
2. Connect upload tickets to R2/S3 signing and virus/archive inspection workers.
3. Connect video releases to a managed transcoder and signed playback.
4. Serve uploaded games from a dedicated sandbox origin with a strict CSP.
5. Add moderator role enforcement and rate limiting at every mutation boundary.
6. Add WebSocket fanout on top of the authoritative arcade room version protocol.
7. Add full rules engines and interactive boards for all five arcade games.
8. Add search indexes and recommendation ranking.
