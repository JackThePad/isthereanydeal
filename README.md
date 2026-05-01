# Credits
Uses [isthereanydeal](https://isthereanydeal.com/)'s api to get game data
# Usage
This tool will take all of the games from your personal steam wish list and use isthereanydeal's api to find deals.
If any deals are found it will send a notification using ntfy.

Build and run with this `config.toml` format in the root dir.
---
```
[config]
itad_api_key="{ISTHEREANYDEAL API KEY}"
ntfy_url="{NTFY FULL URL}"
steam_api_key="{STEAM API KEY}"
steam_account_id="{YOUR STEAM ID}"
json_name="games.json"
```
---
# Credits
Uses [isthereanydeal](https://isthereanydeal.com/)'s api to get game data
---
Make sure anything based off this is open sourced. I do not feel like hunting down a license so honor system on this one.

There is a known bug that some urls change each time you see them so that ruins the unique url check. I may or may not fix this problem.
