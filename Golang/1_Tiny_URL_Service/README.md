# TinyURL

TinyURL service to provide short URL when user requests by providing arguement as long url and vice versa.

## Pre-requisite

1. Go `go1.18 linux/amd64`
2. Redis `sudo docker run -d -p 6379:6379 --name my-redis redis`
   - `redis-cli`
   - Fire `ping` output should be `pong`

## Tech Stack

- Go
- Redis
- Postgresql

## Algorithm

To provide short url we're using `Base58` encoding. Folks can place the argument why we're not using `Base64` ? (<--  ) Reasons for using `Base58`:

1. The characters `o,O,I,1` are confusing when used in certain fonts and difficult to diffrentiate for people with visuality issues.
2. Removing punctuations characters prevent confusion for line breakers.
3. Double-clicking selects the whole number as one word if it's all alpha numeric.

## APIs
### 1.  Welcome Message
`curl -XGET localhost:9808/`
### Response
`"message":"Hey Go URL Shortener !"}`
### 2.  POST Request
`curl --request POST \
--data '{
    "long_url": "https://stackoverflow.com/users/3649496/infinitelearner",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b11"
}' \
  http://localhost:9808/create-short-url`
### Response
`"message":"short url created successfully","short_url":"http://localhost:9808/NdBiifkV"}`
### 3. Retrieving short URL
`curl -XGET http://localhost:9808/NdBiifkV`
### Response
`<a href="https://stackoverflow.com/users/3649496/infinitelearner">Found</a>`


## References

Below references have helped me in writing URL shortener.

- https://www.eddywm.com/lets-build-a-url-shortener-in-go/
- https://www.codekarle.com/system-design/TinyUrl-system-design.html
- https://www.thinksoftwarelearning.com/pages/tiny-url-design

