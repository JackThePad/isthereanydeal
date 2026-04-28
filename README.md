# Usage
This tool will take all of the games from your personal steam wish list and use isthereanydeal's api to find deals.
If any deals are found it will send a notification using ntfy.

Build and run with this `config.toml` format in the root dir.
---
```
[config]
itad_client_id="{ISTHEREANYDEAL CLIENT ID}"           # Currently not used
itad_client_secret="{ISTHEREANYDEAL CLIENT SECRET}"   # Currently not used
itad_api_key="{ISTHEREANYDEAL API KEY}"
ntfy_url="{NTFY FULL URL}"
steam_api_key="{STEAM API KEY}"
steam_account_id="{YOUR STEAM ID}"
```
---
# Credits
Uses [isthereanydeal](https://isthereanydeal.com/)'s api to get game data