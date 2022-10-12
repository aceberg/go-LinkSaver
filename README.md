<h1><a href="https://github.com/aceberg/go-LinkSaver">
    <img src="https://raw.githubusercontent.com/aceberg/go-LinkSaver/main/assets/logo.png" width="20" />
</a>go-LinkSaver</h1>
<br/>

[![Docker](https://github.com/aceberg/go-LinkSaver/actions/workflows/docker-publish.yml/badge.svg)](https://github.com/aceberg/go-LinkSaver/actions/workflows/docker-publish.yml)
![Docker Image Size (latest semver)](https://img.shields.io/docker/image-size/aceberg/go-linksaver)

Simple and lightweight link saver   
https://github.com/aceberg/go-LinkSaver


![Screenshot1](https://raw.githubusercontent.com/aceberg/go-LinkSaver/main/assets/Screenshot%202022-09-03%20at%2023-15-44%20Link%20Saver.png)
![Screenshot2](https://raw.githubusercontent.com/aceberg/go-LinkSaver/main/assets/Screenshot%202022-09-03%20at%2023-17-28%20Link%20Saver.png)
![Screebshot3](https://raw.githubusercontent.com/aceberg/go-LinkSaver/main/assets/Screenshot%202022-09-03%20at%2023-16-31%20Link%20Saver.png)

## Quick start

```sh
docker run --name linksaver \
-e "TZ=Asia/Novosibirsk" \
-v ~/.dockerdata/linksaver:/data/linksaver \
-p 8841:8841 \
aceberg/go-linksaver
```

## Config


Configuration can be done through config file or environment variables

| Variable  | Description | Default |
| --------  | ----------- | ------- |
| DBPATH    | Path to Database | /data/linksaver/db.sqlite |
| GUIPORT   | Port for web GUI | 8841 |
| THEME | Any theme name from https://bootswatch.com in lowcase | minty |
| TZ | Set your timezone for correct time | "" |

## Config file

Config file path is `/data/linksaver/config`. All variables could be set there.

## Thanks
- All go packages listed in [dependencies](https://github.com/aceberg/go-LinkSaver/network/dependencies)
- [Bootstrap](https://getbootstrap.com/)
- Themes: [Free themes for Bootstrap](https://bootswatch.com)