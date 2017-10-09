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

# Setup dep
brew install dep

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

# Build/Deploy
cd functions/libgobuster/
dep ensure
cd ../..
apex deploy

# List deployed functions
apex list

# Run locally (for testing)
cd functions/libgobuster/
go run main.go local
go run main.go local '{"url": "http://devalias.net/", "wordlist":"words.txt"}'

# Run
apex invoke gobuster
apex invoke libgobuster

echo '{"url": "http://devalias.net/", "wordlist":"words.txt"}' | apex invoke libgobuster | jq
echo '{"url": "http://devalias.net/", "wordlist":"words.txt", "sliceStart": 1, "sliceEnd": 1}' | apex invoke libgobuster | jq

# View logs
apex logs gobuster
apex logs libgobuster

# View metrics/costs
apex metrics --since 8760h

# Advanced Usage

# Create some sample files to match on our test bucket
./make-sample-files.sh

# Example of slicing a dictionary across multiple calls
./invoke-multi.sh
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
* Apex uses STDOUT for control, so sometimes it breaks if random things are printed to STDOUT that it doesn't expect.
    * Some parts of libgobuster have been updated to address this, but not all. Eg. if you hit a 'wildcard redirect' it will break currently, with no useful log messages in `apex logs` aside from `SyntaxError: Unexpected token ]`

## Apex

* http://apex.run
* https://github.com/apex/apex

## AWS Lambda

* https://aws.amazon.com/lambda/
* https://aws.amazon.com/lambda/pricing/
* http://docs.aws.amazon.com/lambda/latest/dg/limits.html

## Docker Lambda

* https://github.com/lambci/docker-lambda : Docker images and test runners that replicate the live AWS Lambda environment
* `docker run -it lambci/lambda bash`

## S3 Bucket Policy

The following policy allows the files uploaded by `./make-sample-files.sh` to be accessed by anyone.

```
{
  "Id": "Policy1507512286573",
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "Stmt1507512284907",
      "Action": [
        "s3:GetObject"
      ],
      "Effect": "Allow",
      "Resource": "arn:aws:s3:::gopherblazer-test-discovery/*",
      "Principal": "*"
    }
  ]
}
```
