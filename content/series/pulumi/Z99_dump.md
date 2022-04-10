# Pulumi

Most web applications used today are deployed on an Infrastructure-as-a-Service (IaaS) provider (i.e. 'cloud provider') rather than on-premises. All major cloud providers provide a web-based graphical user interface (GUI) for administrators to manage their resources. For example, an administrator can use such GUI to create a new Linux server, configure a load balancer, add a user, etc.

Pulumi is an Infrastructure-as-Code (IaC) platform, similar to Terraform. But the biggest difference is it allows you to write the resource configuration in a programming language like TypeScript, JavaScript, Python, Go, and .NET, instead of the [HashiCorp Configuration Language](https://github.com/hashicorp/hcl) (HCL). This means developers don't need to learn a new language in order to work with infrastructure resources, lowering the barrier to entry for developers and making it easier for them to own the entire cycle from developing features to solving production issues.

## Concepts

![](https://www.pulumi.com/images/docs/pulumi-programming-model-diagram.svg)

Pulumi code are written in files called _programs_ (e.g. `core.ts`). Programs 

Programs are gathered together into their own directory (e.g. `acme-infra/`) called a _project_. Each project directory must have a `Pulumi.yaml` file that specifies the metadata for the project.

Programs are deployed to a _stack_, which is an isolated, independently configurable instance of the program. Stacks allows you to re-use the same program but with different parameters. For example, you may have a program that lays out the standard .

You can have many stacks within a project, but only one of them will be active at any time. You can see the list of all stacks in the project by running `pulumi stack ls`. The active stack will have an asterisk next to the name. You can change the active stack by running `pulumi stack select <stack>`. To view details of a stack, run `pulumi stack`.

When you run `pulumi up`, the current project will be deployed to the active stack. When this happens, programs within the project are executed by a _language host_ to compute the desired state of the stack's infrastructure. The _deployment engine_ then uses _resource providers_ to communicate with cloud providers and external services to find the difference between the desired state and the actual state. Then, the deployment engine will come up with a plan - a list of modifications needed to get the infrastructure to the desired state.

If you accept the plan, then the deployment engine will use the same resource providers to instruct them to create, update, or delete actual infrastructure resources.

Whenever any operations are pending, succeeds, or fail, the deployment engine communicates with the _backend_ to update the _state_. Each stack has its own state. Pulumi stores state in a backend.

Backends:

- Hosted Service (default)
  - Pulumi Service hosted at `app.pulumi.com`. encrypts secrets using encryption provider, Concurrent state locking
- [Self-hosted Service](https://www.pulumi.com/docs/guides/self-hosted/) (requires Pulumi Enterprise)
  - Pulumi Service application
- Self-hosted Self-Managed object store. State stored as simple JSON files
  - [AWS S3](https://www.pulumi.com/docs/intro/concepts/state/#logging-into-the-aws-s3-backend) (`pulumi login s3://<bucket-name>`)
  - Azure Blob Storage
  - Google Cloud Storage
  - AWS S3 compatible server
    - Minio
    - Ceph
  - Local filesystem

Stacks - dev, prod, or feature. Each stack has its own configuration file at `Pulumi.<stack>.yaml` (e.g. `Pulumi.dev.yaml`)

By splitting stacks into more granular level:
  pros:
    - diffing the states is faster (less to diff)
    - can deploy resources concurrently
    - in an organization (using Pulumi Service) you can give different groups of users different permissions to deploy different resources - e.g. devs can't change core infrastructure stuff. See https://www.pulumi.com/docs/intro/console/teams/

permanent stacks - `Pulumi.<stack>.yaml` checked in
ephemeral stacks - `Pulumi.<stack>.yaml` not checked in

## Getting Started

CLI, runtime, libraries, and a hosted service

Run `pulumi up` within the project directory. This commands creates a _stack_ - an instance of your program.

## Alternatives

ARM Bicep
[AWS Cloud Development Kit](https://aws.amazon.com/cdk/)

https://www.hashicorp.com/blog/announcing-cdk-for-terraform-0-4
https://docs.microsoft.com/en-us/azure/azure-resource-manager/bicep/overview
