#!/usr/bin/env python3
"""Prepare unique role/policy inputs. Does not create cloud resources."""
import datetime,json,pathlib,subprocess
root=pathlib.Path(__file__).resolve().parents[1]
state=root/'terraform.tfstate'
if state.exists() and json.loads(state.read_text()).get('resources'):
 raise SystemExit('Existing resources in state: destroy them before preparing a new run.')
identity=json.loads(subprocess.check_output([str(root/'bin/cloud'),'identity'],text=True))
run=datetime.datetime.now(datetime.timezone.utc).strftime('%Y%m%d%H%M%S')
role='cc-proxy-regression-'+run
trust={'Statement':[{'Effect':'Allow','Principal':{'IAM':[identity['Trn']]},'Action':['sts:AssumeRole']}]}
schema=json.loads((root/'../../internal/service/cloudcontrol/schemas/Volcengine_VPC_Vpc.json').read_text())
permissions=sorted({x for h in schema['handlers'].values() for x in h['permissions']})
policy={'Statement':[{'Effect':'Allow','Action':permissions,'Resource':['*']},{'Effect':'Allow','Action':['cloudcontrol:*'],'Resource':['*']}]}
artifacts=root/'artifacts'; artifacts.mkdir(exist_ok=True)
def save(name,value): (artifacts/name).write_text(json.dumps(value,indent=2))
save('create-role.json',{'RoleName':role,'Description':'Temporary AssumeRole proxy regression','MaxSessionDuration':3600,'TrustPolicyDocument':json.dumps(trust)})
save('create-policy.json',{'PolicyName':role,'Description':'Temporary VPC proxy regression permissions','PolicyDocument':json.dumps(policy)})
save('binding.json',{'RoleName':role,'PolicyName':role,'PolicyType':'Custom'})
save('role-key.json',{'RoleName':role})
save('policy-key.json',{'PolicyName':role})
save('run.json',{'role':role,'policy':role,'account_id':identity['AccountId'],'principal':identity['Trn']})
(root/'run.auto.tfvars.json').write_text(json.dumps({'role_trn':f'trn:iam::{identity["AccountId"]}:role/{role}','vpc_name':role},indent=2))
print('Prepared '+role+'; trust principal: '+identity['Trn'])
