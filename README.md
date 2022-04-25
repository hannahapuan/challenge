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

### Installing and Executing

- Clone repo
```
git clone github.com/hannahapuan/challenge && cd challenge
```

- Build docker image
```
docker build -t challenge .
```

- Run image, specify environment variables (port), and expose port
```
 docker run  -p 8081:8081 -e port=8081 -it challenge

```
* note that the port environment variable if not specified will default to 8081

- Querying is avaliable on the `localhost:8081/graphql` endpoint
