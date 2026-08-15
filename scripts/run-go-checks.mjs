import { spawnSync } from "node:child_process";
import path from "node:path";
import { fileURLToPath } from "node:url";

const root = path.resolve(path.dirname(fileURLToPath(import.meta.url)), "..");
const run = (command, args, options = {}) => {
  const result = spawnSync(command, args, { cwd: root, stdio: "inherit", ...options });
  if (result.error?.code === "ENOENT") return false;
  if (result.status !== 0) process.exit(result.status ?? 1);
  return true;
};

if (run("go", ["version"])) {
  for (const directory of ["packages/go", "extensions/github/go"]) {
    run("go", ["-C", directory, "test", "./..."]);
    run("go", ["-C", directory, "vet", "./..."]);
  }
} else {
  const mount = `${root}:/work`;
  const script = [
    "cd /work/packages/go && /usr/local/go/bin/go test ./... && /usr/local/go/bin/go vet ./...",
    "cd /work/extensions/github/go && /usr/local/go/bin/go test ./... && /usr/local/go/bin/go vet ./...",
  ].join(" && ");
  if (!run("docker", ["run", "--rm", "-v", mount, "-w", "/work", "golang:1.22", "sh", "-lc", script])) {
    throw new Error("Go 1.22+ or Docker is required for Go conformance checks.");
  }
}
