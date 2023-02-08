---
title: AWS
slug: aws
date: 2023-02-07T20:44:43.142Z
publishDate:
chapter: z
order: 99
tags:
    - aws
draft: true
---

# VPCs

VPCs allows you to isolate resources into completely different networks.

VPCs are region-specific. When you create a VPC, it spans all availability zone within the region.

## Default VPC

## Availability Zones

"""AZs are distinct locations that are engineered to be isolated from failures in other Availability Zones. By launching instances in separate Availability Zones, you can protect your applications from the failure of a single location"""(https://www.pulumi.com/docs/reference/pkg/nodejs/pulumi/awsx/ec2/#AllTcpPorts)

# Subnets

Subnets (short for sub-networks) is a subsection of a network (in AWS, a subnet is a subsection of a VPC).
In AWS, a subnet is AZ-specific, which means it must be associated within one of the VPC's region's availability zones and cannot span across multiple AZs.

- Public
  - in a VPC with an Internet Gateway attached to it
  - associated with a route table that has a route to the IGW
- Private
  - associated with a route table that does not have a route to the IGW
  - it can access the Internet via a NAT gateway in the public subnet

# Security Groups (SGs)

"""
- You can specify allow rules, but not deny rules.
- You can specify separate rules for inbound and outbound traffic.
- Security group rules enable you to filter traffic based on protocols and port numbers.
"""(https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html)

- Group together EC2 instances with similar functions (e.g. bastion host, web servers, log servers, monitoring servers, databases, etc.)
- An EC2 instance can only be associated with multiple SGs; likewise, an SG can be associated with multiple EC2 instances
- Act as a stateful application-layer firewalls for EC2 instances within each SG, where you can specify firewall rules based on TCP/UDP port
- Stateful - """if you send a request from your instance, the response traffic for that request is allowed to flow in regardless of inbound security group rules. Responses to allowed inbound traffic are allowed to flow out, regardless of outbound rules."""(https://docs.aws.amazon.com/vpc/latest/userguide/VPC_SecurityGroups.html)
- rules can be restricted by source CIDR block as well as other security groups (e.g. database SG only allows traffic from EC2 instances in the backend SG)
- SGs are VPC-specific (they cannot be used across multiple VPCs)

# NACLs

- Network-layer firewall for a subnet
- A subnet can only be associated with a single NACL
- Stateless
- define different rules for inbound vs. outbound traffic to other subnets as well as to the Internet
- rules can be restricted by source CIDR block
- NACL rules are evaluated in order, based on the rule number (from lowest to highest)
  - If a rule applies, it is used immediately and subsequent rules are disregarded
  - there's a catchall/failsafe rule at the end (with rule 'number' `*`) that blocks all traffic, which means if no rules matches, the traffic will be denied
  - the default NACL that comes with a VPC has a rule (`#100`) that allows all traffic, newly-created NACLs don't have this rule and denies everything by default
- NACLs are VPC-specific (they cannot be used across multiple VPCs)

1024 - 65535

## AWS CodeBuild

`buildspec.yml`

## AWS CodeDeploy

`appspec.yml`

## AWS CodeStar

An AWS CodeStar _project_ uses AWS CodePipeline with AWS CodeCommit and AWS CodeDeploy to store and deploy your code. Everything is set up for you using AWS CodeStar (you can think of CodeStar as a combination of all the `Code*` services on AWS).
Develop - Cloud9 IDE (or your own IDE) + CodeCommit
Build - CodeBuild
Deploy - CodeDeploy
Orchestrating CI/CD - CodePipeline

The AWS CodeStar service needs to create a service role that allows it to interact with the AWS `Code*` services.

CodeStar has integrations with Visual Studio, Eclipse, and the command line interface.

## AWS Cloud9

AWS Cloud9 is an online, in-browser IDE for writing code.
