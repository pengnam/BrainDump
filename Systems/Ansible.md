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

A play maps a group of hosts to some well defined roles, represented by things ansible calls tasks. 

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

Handlers list of tasks referenced by a globally unique name, notified by notifiers. Handlers are listed in the `notify` section. 


#### Notable commands

ansible-pull : small script that checksout a repo of configuration instructions from git

ansible-lint : runs a detailed check of playbook

#### Variables

Should be specified in the config files. There should be a concern when replaying as the variables might not be the same/not declared

#### Register Variables

Register variables are used when we want to capture output of a task to a register. Some modules have specific return values.


### Lookup Plugins
Lookup plugins allow Ansible to access data from outside sources. This means that ansible can read file system, contact external datastores, services. Evaluated on the Ansible control machine and not the target or remote.

Lookup plugins can be used anywhere in the templating in Ansible: in a play, variables file, Jinja2 template for the template module. 

### Query Plugins

`query` is used for lookup plugins. 
