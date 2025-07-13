<h1 align="center">🗺️ OSM Changesets Bot</h1>

> [!NOTE]
> The initial version (in Python) is still available in the [python branch](https://github.com/GHOSTsama2503/OSM-Changesets-Bot/tree/python)

Easy way to see the [changesets](https://wiki.openstreetmap.org/wiki/Changeset) in a certain area in [Telegram](http://telegram.org). Send the messages directly to you, or create a channel and share it with other users.

### Set up the environment variables

Environment variables are necessary, to declare them just create a `.env` file at the root of the project (you can use another method if you prefer).

-   **BOT_TOKEN**: Telegram bot token, is obtained from [@BotFather](https://t.me/BotFather) (Create a new bot if required)
-   **CHANNEL_ID**: Unique identifier for the target chat.
-   **FEED_URL**: URL of the [OSMCha](https://osmcha.org) filter that you want to use.

-   **DATABASE_URL**: Custom dstabase URL (default `database.db`).
-   **CHECKS_INTERVAL**: Time in seconds between each feed parse (default 60s).
-   **RETRIES_INTERVAL**: Time in seconds to wait after each error before running the task again (default 5s).
