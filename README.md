# AWS Account Manager

This is tool to manage AWS accounts.
You can sign in an account and switch to a role.

## Usage

**add**

Manage new AWS account.

```sh
$ awsman add 123456789012 --alias pj-prod --role OpeRole
```

**remove**

Unmanage AWS account.

```sh
$ awsman remove pj-prod
```

**list**

List up AWS accounts managed by this tool.

```sh
$ awsman list
AWS Accounts
├── pj-prod
│   ├── [ID  ]  123456789012
│   └── [Role]  OpeRole
└── pj-stag
    ├── [ID  ]  987654321098
    └── [Role]  DevRole
```

**signin**

Sign-in to the account and switch to a role.

```ah
# Sign-in '123456789012' and switch to 'OpeRole'
awsman signin pj-prod
```

## How to Install

```sh
git clone https://github.com/x-color/awsman
cd awsman
go install
```