# AssumeRole proxy regression

This example creates one temporary VPC in `cn-beijing`, checks that a refresh produces no changes, destroys it, and verifies absence with the VPC API. Use a test account and provide source credentials through `VOLCENGINE_ACCESS_KEY` / `VOLCENGINE_SECRET_KEY`. No credentials are stored in the example.

The provider pins SDK commit `410d17f17d99`, containing the STS HTTP/2 transport fix. Source builds require access to `code.byted.org` and that host in `GOPRIVATE`.

## Prepare

Run from this directory with Go 1.25.8, Terraform, Python 3, and your source credentials already in the environment:

```bash
export GOPRIVATE="${GOPRIVATE:+$GOPRIVATE,}code.byted.org"
./scripts/build.sh
python3 scripts/prepare.py
./bin/cloud create-role @artifacts/create-role.json
./bin/cloud create-policy @artifacts/create-policy.json
./bin/cloud attach-policy @artifacts/binding.json
```

`prepare.py` generates a unique role restricted to the current caller's TRN and a custom policy covering the VPC handler permissions plus `cloudcontrol:*`. It does not create resources. If provisioning partially fails, clean up the successfully created objects before starting a new run.

Start the loopback-only proxy in a separate terminal. It allows only CONNECT to `sts.volcengineapi.com:443` and logs destinations and byte counts, without decrypting TLS or logging headers:

```bash
python3 scripts/connect_proxy.py --port 18080 --log artifacts/proxy.jsonl
```

## Run both configurations

```bash
./scripts/cycle.sh environment
./scripts/cycle.sh explicit
```

The first mode omits `assume_role.proxy_url` and uses `HTTP_PROXY=http://127.0.0.1:18080`. The second uses the hardcoded proxy in `explicit.tfvars` while setting `HTTP_PROXY=http://127.0.0.1:1`, proving explicit configuration wins. The script requires an empty initial state and attempts destruction on failure. If destruction fails, retain the role and state for recovery.

`dev_overrides` loads the locally built provider; do not run `terraform init` to download a registry release. Artifacts are ignored by Git. Each mode overwrites its own logs, so archive them before another run when needed.

## Clean up IAM

After both deletion checks report `TotalCount=0` and Terraform state contains no resources:

```bash
./bin/cloud detach-policy @artifacts/binding.json
./bin/cloud delete-role @artifacts/role-key.json
./bin/cloud delete-policy @artifacts/policy-key.json
./bin/cloud get-role @artifacts/role-key.json
```

The final command should return `RoleNotExist`. Stop the proxy and clear source credentials from your shell. Run `prepare.py` again for a new test; deleted roles cannot be reused.

## Local SDK development

The committed `go.mod` uses a remote immutable dependency. To test edits in a local SDK checkout without changing it, use an alternate module file from the repository root:

```bash
cp go.mod /tmp/volcenginecc-local-sdk.mod
cp go.sum /tmp/volcenginecc-local-sdk.sum
go mod edit -modfile=/tmp/volcenginecc-local-sdk.mod \
  -replace=github.com/volcengine/volcengine-go-sdk=/absolute/path/to/volcengine-go-sdk
GOFLAGS=-modfile=/tmp/volcenginecc-local-sdk.mod \
  examples/assume-role-proxy/scripts/build.sh
go test -modfile=/tmp/volcenginecc-local-sdk.mod -race ./internal/provider
```

Check out `410d17f17d99` or a descendant containing the fix. The unpatched base commit `81fa3aa1ffe0` can fail against HTTP/2 STS endpoints after CONNECT. Remove the temporary module files and rebuild without `GOFLAGS` to restore the committed dependency.
