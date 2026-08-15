import { readFileSync, readdirSync, writeFileSync } from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const workspace = path.join(root, "packages", "typescript");
const sourceRoot = path.join(workspace, "src");
const outputRoot = path.join(workspace, "dist");

function filesBelow(directory) {
  return readdirSync(directory, { withFileTypes: true }).flatMap((entry) => {
    const child = path.join(directory, entry.name);
    return entry.isDirectory() ? filesBelow(child) : [child];
  });
}

for (const mapPath of filesBelow(outputRoot).filter((filePath) => filePath.endsWith(".d.ts.map"))) {
  const sourceMap = JSON.parse(readFileSync(mapPath, "utf8"));
  sourceMap.sourcesContent = sourceMap.sources.map((source) => {
    const sourcePath = path.resolve(path.dirname(mapPath), source);
    const relative = path.relative(sourceRoot, sourcePath);
    if (relative.startsWith("..") || path.isAbsolute(relative)) {
      throw new Error(`Declaration map source escapes the TypeScript source tree: ${source}`);
    }
    return readFileSync(sourcePath, "utf8");
  });
  writeFileSync(mapPath, JSON.stringify(sourceMap));
}
