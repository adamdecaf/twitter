# twitter

### Tasks

- Total number of tweets received
- Average tweets per hour/minute/second
- Top emojis in tweets
- Percent of tweets that contains emojis
- Top hashtags
- Percent of tweets that contain a url
- Percent of tweets that contain a photo url (pic.twitter.com or instagram)
- Top domains of urls in tweets

### Projects

**Deployables**

- ingest
- parse/emojis
- parse/hashtags
- parse/urls

**Libraries**

- storage/cassandra
- storage/kafka

### Development

- Building: `make build [proj=<project>]`
- Testing: `make test [proj=<project>]`
