import { readdirSync } from "node:fs";
import { spawnSync } from "node:child_process";
import path from "node:path";
import { fileURLToPath } from "node:url";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const run = (command, args, { quiet = false, cwd = root } = {}) => {
  const result = spawnSync(command, args, { cwd, encoding: "utf8", shell: process.platform === "win32" });
  if (result.error) throw result.error;
  if (result.status !== 0) {
    process.stdout.write(result.stdout ?? ""); process.stderr.write(result.stderr ?? "");
    process.exit(result.status ?? 1);
  }
  if (!quiet) process.stdout.write(result.stdout ?? "");
  return result.stdout ?? "";
};
const phpFiles = (directory) => readdirSync(path.join(root, directory), { recursive: true, withFileTypes: true })
  .filter((entry) => entry.isFile() && entry.name.endsWith(".php"))
  .map((entry) => path.join(entry.parentPath, entry.name));

for (const file of [...phpFiles("packages/php/src"), ...phpFiles("extensions/github/php")]) run("php", ["-l", file], { quiet: true });
run("composer", ["validate", "--strict", "packages/php/composer.json"]);
run("composer", ["validate", "--strict", "extensions/github/php/composer.json"]);
const corePack = JSON.parse(run("npm", ["pack", "--dry-run", "--json", "--workspace", "@elqora/artifacts"], { quiet: true }));
const githubPack = JSON.parse(run("npm", ["pack", "--dry-run", "--json"], { quiet: true, cwd: path.join(root, "extensions", "github", "typescript") }));
if (corePack[0]?.name !== "@elqora/artifacts" || githubPack[0]?.name !== "@elqora/artifact-github") {
  throw new Error("TypeScript package dry run produced an unexpected package identity.");
}
process.stdout.write(`Package validation passed: ${corePack[0].name}, ${githubPack[0].name}.\n`);
