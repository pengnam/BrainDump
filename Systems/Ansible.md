# Ansible

## Introduction

Ansible is an automation tool, with strong focus on security and reliability. Use of OpenSSH for transport.

Ansible manages machines in an agent-less manner. Decentralized by using existing OS credentials to control access to remote machines. It will try to use native OpenSSH for remote communication when possible. 

This enables ControlPersist, Kerberos, and JumpHost. Ansible does not expose a channel to allow communication between user and the ssh process. 


## Two modes

1. Ad-hoc task execution

2. Playbook execution

Playbook are the basis for really simple configuration management and multi-machine deployment system. 

While I might run the main program for ad-hoc tasks, playbooks are more likely to be kept in source control. 

## Playbooks

Consists of one or more plays.

A play maps a group of hosts to some well defined roles, represented by thing sansible calls tasks. 

A task is a call to an ansible module. 


By composing a playbook of multiple 'plays', it is possible to orchestrate multi-machine deployments. 


### Basics

#### Hosts

For each play, I get to define which machines are targeted, and what the remote user should do to complete the tasks. 

`hosts` line is a list of 1 or more groups or host patterns, separated by colons. `remote_user` is name of user account. 

#### Tasks list

Each play contains a list of tasks. Tasks executed in order, one at a time, against all machines matched by a host pattern. 

Hosts with failed tasks are taken out of rotation for entire playbook. 

Modules should be idempotent. Should always check if desired final state has been achieved, and if state has been achieved, to exit without performing any actions. 

`name` is a human readable output.

#### Handlers: Running operations on change

Handlers list of tasks referenced by a globally unique name, notified by notifiers. 


#### Notable commands

ansible-pull : small script that checksout a repo of configuration instructions from git

ansible-lint : runs a detailed check of playbook


