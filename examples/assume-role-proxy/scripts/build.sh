#!/usr/bin/env bash
set -euo pipefail
example_dir="$(cd "$(dirname "$0")/.." && pwd)"
provider_dir="$(cd "$example_dir/../.." && pwd)"
mkdir -p "$example_dir/bin" "$example_dir/artifacts"
(cd "$provider_dir" && go build -o "$example_dir/bin/terraform-provider-volcenginecc" .)
(cd "$provider_dir" && go build -o "$example_dir/bin/cloud" ./examples/assume-role-proxy/tools/cloud)
python3 - "$example_dir" <<'PY'
from pathlib import Path
import json,sys
root=Path(sys.argv[1])
(root/'terraform.rc').write_text('provider_installation {\n  dev_overrides {\n    "volcengine/volcenginecc" = '+json.dumps(str(root/'bin'))+'\n  }\n  direct {}\n}\n')
PY
