# PoC-Apex

PoC showing an [Apex](http://apex.run/) function that directly uses `gobuster`/`libgobuster`, running on AWS lambda.

## Setup

* Make sure your AWS credentials have the following policy
    * http://apex.run/#minimum-iam-policy

```
export AWS_PROFILE=devalias

# Setup Apex
curl https://raw.githubusercontent.com/apex/apex/master/install.sh | sh
apex upgrade

# Init new project
mkdir foo
cd foo
apex init
```

Note: You will have to update the `role` ARN in `project.json` if you want to use this example directly. You can use the value created when init'ing a new project.

## Usage

```
# Create a new function
# See http://apex.run/#structuring-functions

# Deploy
apex deploy

# List deployed functions
apex list

# Run
apex invoke gobuster
apex invoke libgobuster

# View logs
apex logs gobuster
apex logs libgobuster

# View metrics/costs
apex metrics --since 8760h
```

## Destroy

You will need to manually clean up the following
* [AWS Roles](https://console.aws.amazon.com/iam/home?region=ap-southeast-2#/roles): `FUNCNAME_lambda_function`
* [AWS Policies](https://console.aws.amazon.com/iam/home?region=ap-southeast-2#/policies): `FUNCNAME_lambda_logs`

```
export AWS_PROFILE=devalias

# Delete deployed functions
apex delete

# Cleanup any AWS roles/policies
#aws iam delete-role --role-name FUNCNAME_lambda_function
#aws iam delete-policy --policy-arn THE_POLICY_ARN
```

## Future Improvements

* Do we need to use Apex? Could probably build our own more lightweight system around this, build in docker containers, export artificats, etc. I believe there are some really good shims out there..

## Apex

* http://apex.run
* https://github.com/apex/apex

## AWS Lambda

* https://aws.amazon.com/lambda/
* https://aws.amazon.com/lambda/pricing/
* http://docs.aws.amazon.com/lambda/latest/dg/limits.html
