# aggreGATOR

An RSS feed aggregator (aggre gator 🐊) guided project from [Boot.dev](https://boot.dev)

A CLI tool that:

- Add RSS feeds from across the internet to be collected
- Store the collected posts in a PostgreSQL database
- Follow and unfollow RSS feeds that other users have added
- View summaries of the aggregated posts in the terminal, with a link to the full post

## Installation

You need [PostgreSQL](https://www.postgresql.org/download/) and [Go toolchain](https://webinstall.dev/golang/) installed for this to work.

Use the following command to install the CLI on your machine:

`go install github.com/burush0/gator`

You will need to create a database in PostgreSQL to store the data that the CLI will fetch and use.

e.g.
```
sudo -u postgres psql # to get into the postgresql CLI
create database gator; -- to create a database
```

You need a file called `.gatorconfig.json` in your `$HOME` directory . It should have the following structure:
```
{"db_url":"postgres://postgres:postgres@localhost:5432/gator?sslmode=disable"}
```

This is under the assumption that you're on Linux using default postgres user with password also set to postgres, on default port, and the database you created in psql is called "gator".

A more generic database connection string looks something like this:

`postgres://<user>:<password>@<ip>:<port>/<table_name>?sslmode=disable`

~~The last thing you need to do is to set up the correct tables in the database, thankfully there are migration files that do that for you, you just need to run this command:~~

~~`goose postgres postgres://postgres:postgres@localhost:5432/gator up`~~

(Uhh how do you run db migrations when using go install, there won't be a root folder, right?)

Then you can use the CLI by calling `gator` directly in your terminal.

Here are some commands you can run:

- reset -- Resets the app to its "clean" state
- register \<username> -- This will register a user in the CLI
- login \<username> -- Switches to another user
- addfeed \<name> \<url> -- This will add an RSS feed
- follow/unfollow \<url> -- Allows to unfollow a feed or follow another user's feed
- following -- Lists the feeds that the user follows
- agg \<time_between_requests> (e.g. 1m) -- Will continuosly check for new posts in the feed
- browse \<limit> (optional, 2 by default) -- Will show the last 2 posts from the feeds that the user is following