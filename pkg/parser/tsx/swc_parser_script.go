package tsx

// swcParserScript is the Node wrapper that calls @swc/core#parseFile and
// writes the AST JSON to stdout. swc's standalone CLI does not expose a
// `parse` subcommand, so the wrapper is the minimal supported entry point.
// It is materialized on disk on first use under $TMPDIR and reused.
const swcParserScript = `
const path = require('path');
const file = process.argv[2];
if (!file) { console.error('usage: swc-parse <file>'); process.exit(2); }
let core;
try {
  core = require('@swc/core');
} catch (e) {
  console.error('yongol: @swc/core not installed. Run one of:');
  console.error('  npm install --save-dev @swc/core');
  console.error('  (at the project root or in YONGOL_SWC_PROJECT_DIR)');
  console.error('underlying error: ' + e.message);
  process.exit(3);
}
core.parseFile(file, {
  syntax: 'typescript',
  tsx: true,
  decorators: false,
  dynamicImport: true,
  comments: false,
}).then(ast => {
  process.stdout.write(JSON.stringify(ast));
}).catch(err => {
  console.error(err && err.message ? err.message : String(err));
  process.exit(1);
});
`

// swcRunnerNotice is the fail-fast installation hint surfaced when the swc
// toolchain is absent. Kept as a single const so CLI integrations can wrap
// it uniformly.
const swcRunnerNotice = "install Node.js (>=18) and run `npm install --save-dev @swc/core` in your frontend project (or set YONGOL_SWC_PROJECT_DIR)"
