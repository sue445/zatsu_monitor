# Zatsu(雑) Monitor
*zatsu_monitor* is simple url monitoring tool

Notify message to chatroom only if http status is changed from just before (e.g. 500 -> 200)

![slack](img/slack.png)

## Getting
Download latest binary from here

https://github.com/sue445/zatsu_monitor/releases

## Usage
```sh
/path/to/zatsu_monitor -config /path/to/config.yml -data /path/to/data_dir
```

* `-config`: Path to config yaml file (see after)
* `-data`: Cache dir. latest http status is stored to here

## Help
```sh
$ ./zatsu_monitor --help
Usage of ./zatsu_monitor:
  -config string
    	Path to config file
  -data string
    	Path to data dir for cache
  -version
    	Whether showing version
```

## Config yaml file format
### Example
```yaml
# zatsu_monitor.yml
google:
  type: slack
  check_url: "https://www.google.com/"
  api_token: "AAAAAAAA"
  channel: "#general"
  user_name: "zatsu_monitor"
github:
  type: chatwork
  check_url: "https://github.com/"
  api_token: "AAAAAAAA"
  room_id: "111111"
```

### Detail
*zatsu_monitor* supports some notifier

#### [Slack](https://slack.com/)
* `type` : *slack* (fixed value)
* `check_url` : URL for checking (required)
* `api_token` : Slack API Token (required)
* `channel` : Channel for post (required)
* `user_name` : Name for post (optional. default is *zatsu_monitor*)

#### [ChatWork](http://www.chatwork.com/)
* `type` : *chatwork* (fixed value)
* `check_url` : URL for checking (required)
* `api_token` : ChatWork API Token (required)
* `room_id` : RoomID for post (required)

## ProTip
### Inherit values
```yml
slack: &common
  type: slack
  channel: "#general"
  api_token: "xoxp-0000000000-0000000000-0000000000-000000"

github:
  # inherit common values
  <<: *common

  # override common values
  check_url: "https://github.com/"
  chennel: "#github"

google:
  <<: *common
  check_url: "https://www.google.com/"
  chennel: "#google"
```

## Development
### Setup
```sh
go get -u ./...
go get -u -t ./...
cp .env.example .env
vi .env
```
