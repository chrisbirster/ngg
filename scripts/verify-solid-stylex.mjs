import { readFile } from "node:fs/promises";

const files = ["src/App.tsx", "src/GamePage.tsx", "src/PortalPage.tsx", "src/components.tsx"];
const invalid = [];
let attributeCalls = 0;

for (const file of files) {
  const source = await readFile(new URL(`../${file}`, import.meta.url), "utf8");
  const propCalls = source.match(/stylex\.props\(/g) ?? [];
  const attrsCalls = source.match(/stylex\.attrs\(/g) ?? [];

  if (propCalls.length > 0) invalid.push(`${file}: ${propCalls.length} stylex.props call(s)`);
  attributeCalls += attrsCalls.length;
}

if (invalid.length > 0) {
  console.error("Solid DOM elements must use stylex.attrs(), which emits `class`. stylex.props() emits React-only `className`.");
  console.error(invalid.join("\n"));
  process.exit(1);
}

if (attributeCalls === 0) {
  console.error("No stylex.attrs() calls found; the production stylesheet may be disconnected.");
  process.exit(1);
}

console.log(`Verified ${attributeCalls} Solid-compatible StyleX attribute calls.`);
