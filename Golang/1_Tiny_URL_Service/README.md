# TinyURL

TinyURL service to provide short URL when user requests by providing arguement as long url and vice versa.

## Tech Stack

- Go
- Redis
- Postgresql

## Algorithm

To provide short url we're using `Base58` encoding. Folks can place the argument why we're not using `Base64` ? 
Reasons for using `Base58`:

1. The characters `o,O,I,1` are confusing when used in certain fonts and difficult to diffrentiate for people with visuality issues.
2. Removing punctuations characters prevent confusion for line breakers.
3. Double-clicking selects the whole number as one word if it's all alpha numeric.

## References

Below references have helped me a lot in writing URL shortener.

- https://www.eddywm.com/lets-build-a-url-shortener-in-go/
- https://www.codekarle.com/system-design/TinyUrl-system-design.html

