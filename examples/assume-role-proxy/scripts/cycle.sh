#!/usr/bin/env bash
# Creates one test VPC, checks refresh, destroys it, and verifies cloud-side absence.
set -euo pipefail
mode="${1:-environment}"
case "$mode" in environment|explicit) ;; *) echo 'usage: cycle.sh environment|explicit' >&2; exit 2;; esac
example_dir="$(cd "$(dirname "$0")/.." && pwd)"
cd "$example_dir"
: "${VOLCENGINE_ACCESS_KEY:?Set credentials in the environment}"
: "${VOLCENGINE_SECRET_KEY:?Set credentials in the environment}"
export TF_CLI_CONFIG_FILE="$example_dir/terraform.rc"
unset VOLCENGINE_ASSUME_ROLE_PROXY_URL HTTP_PROXY http_proxy HTTPS_PROXY https_proxy ALL_PROXY all_proxy
unset VOLCENGINE_PROXY_URL VOLCENGINE_PROXY_AUTHORIZATION VOLCENGINE_NO_PROXY VOLCENGINE_PROXY_INCLUDE_DOMAINS
unset TF_LOG TF_LOG_PATH
if [[ "$mode" == environment ]]; then
  export HTTP_PROXY=http://127.0.0.1:18080
else
  export HTTP_PROXY=http://127.0.0.1:1
fi
python3 - <<'PY'
import json,pathlib
p=pathlib.Path('terraform.tfstate')
if p.exists() and json.loads(p.read_text()).get('resources'):
 raise SystemExit('Refusing to run over existing resources. Review/destroy the previous test first.')
PY
cleanup_on_failure() {
  status=$?
  if (( status != 0 )); then
    echo 'Test failed; attempting to destroy the VPC. Check state before removing its role.' >&2
    terraform destroy -no-color -input=false -auto-approve -var-file="$mode.tfvars" || true
  fi
  exit "$status"
}
trap cleanup_on_failure EXIT
terraform apply -no-color -input=false -auto-approve -var-file="$mode.tfvars" | tee "artifacts/$mode-apply.log"
terraform output -json > "artifacts/$mode-outputs.json"
terraform plan -no-color -input=false -detailed-exitcode -var-file="$mode.tfvars" | tee "artifacts/$mode-refresh.log"
terraform destroy -no-color -input=false -auto-approve -var-file="$mode.tfvars" | tee "artifacts/$mode-destroy.log"
python3 - "$mode" <<'PY'
import json,pathlib,subprocess,sys
mode=sys.argv[1]
vpc=json.loads(pathlib.Path(f'artifacts/{mode}-outputs.json').read_text())['vpc_id']['value']
result=json.loads(subprocess.check_output(['./bin/cloud','describe-vpcs',json.dumps({'VpcIds':[vpc]})],text=True))
pathlib.Path(f'artifacts/{mode}-deleted.json').write_text(json.dumps(result,indent=2))
assert result['Result']['TotalCount']==0,result
assert not json.loads(pathlib.Path('terraform.tfstate').read_text()).get('resources')
print('Verified deleted: '+vpc)
PY
