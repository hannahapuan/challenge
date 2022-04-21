# Challenge

GraphQL application which utilizes the ORM [ent](https://entgo.io/) and [gqlgen](https://gqlgen.com/) to query [SpamhausRBL](https://www.spamhaus.org/) for blacklisted IPs and stores in an SQLite datastore.

## Mutations
- `enqueue(ip: ["ip1", "ip2"])`
  - Kicks off a background job to do the DNS lookup and store it in the database for each IP passed in for future lookups
  - If the lookup has already happened, enqueue queues it up again and updates the `response` and `updated_at` fields

## Queries
- `getIPDetails(ip: "ip address here")`
  - Looks up the IP Address in the database
  - Response format:
    - `uuid` ID
    - `created_at` time
    - `updated_at` time
    - `response_code` string
    - `ip_address` string

## Getting Started

### Dependencies

* Docker
* Go 1.18

#### External Packages used
- [Ent](https://entgo.io/) 
  - Used as an ORM to manage the schema for the IP address CRUD actions and interactions with SQLite
- [Ajnasz/dnsbl-check](github.com/Ajnasz/dnsbl-check)
  - This fork included in the repo changes the command line tool to have exported functions that can be used as a dependency to look up DNSBL using a provided provider (e.g. zen.spamhaus.org).

### Installing

* How/where to download your program
* Any modifications needed to be made to files/folders

### Executing program

* How to run the program
* Step-by-step bullets
```
code blocks for commands
```

## Help

Any advise for common problems or issues.
```
command to run if program contains helper info
```

## Authors

Contributors names and contact info

ex. Dominique Pizzie  
ex. [@DomPizzie](https://twitter.com/dompizzie)

## Version History

* 0.2
    * Various bug fixes and optimizations
    * See [commit change]() or See [release history]()
* 0.1
    * Initial Release

## License

This project is licensed under the [NAME HERE] License - see the LICENSE.md file for details

## Acknowledgments

Inspiration, code snippets, etc.
* [awesome-readme](https://github.com/matiassingers/awesome-readme)
* [PurpleBooth](https://gist.github.com/PurpleBooth/109311bb0361f32d87a2)
* [dbader](https://github.com/dbader/readme-template)
* [zenorocha](https://gist.github.com/zenorocha/4526327)
* [fvcproductions](https://gist.github.com/fvcproductions/1bfc2d4aecb01a834b46)