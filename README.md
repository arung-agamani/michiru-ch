# Michiru Ch.

_Turn Discord into a CI/CD Developer Platform_

## About

This is a web application that manages incoming webhooks for various git provider (currently only on Github) and forwards it into a target channel in a Discord server. Leveraging various features that a Discord application/bot able to do, such as context menu, modal, dialog, slash commands, and message commands, manage your deployments from Discord with more ease.

## Why

Kinda annoying to switch context and window every time I need to do deployment and I kinda want to do it in Discord. Does anyone ever do this, idk, but hey it seems cool.

Also because I want to further apply the knowledge I learned about building automation tools, and this time, with Golang.

## Requirements

-   Postgres Database. Currently no plugins used so a plain installation is enough
-   Discord application. Register one through Discord developer portal and make a bot user, and also retrieve the bot token.

The following table shows the environment variable used.
| Variable Name | Description | Required |
|--------------------|-----------------------------------|----------|
| `DATABASE_URL` | URL for connecting to Postgres DB | Yes |
| `DISCORD_BOT_TOKEN`| Token for authenticating the bot | Yes |

## Installation

To install Michiru Ch., follow these steps:

1. Clone the repository:
    ```sh
    git clone https://github.com/arung-agamani/michiru-ch.git
    ```
2. Navigate to the project directory:
    ```sh
    cd michiru-ch
    ```
3. Build the application:
    ```sh
    go build
    ```

## Usage

1. Run the application:
    ```sh
    ./michiru-ch
    ```
2. Configure your Discord bot and set up the necessary webhooks from your git provider.
3. Invite the bot to your Discord server and configure the target channel for notifications.

## What's with the name?

### Answer with philosophy

Michiru (満ちる) in Japanese means "to fill" "to become full" "to satisfy". This project is made to satisfy my needs of an automation tool from Discord itself.

### Answer with not much of philosophy

Michiru refers is a character from gacha game Blue Archive with name Chidori Michiru. Previously I made a Discord bot Izuna, which in lore is a member of Ninjutsu Research Club, together with Michiru. I thought Izuna (bot) is kinda lonely, so I just add her companion, Michiru (bot).

There is also Tsukuyo, but more on that later... when I have ideas on what kind of bot she will become... hehe.

## Closing Remarks

Thank you for using Michiru Ch.! If you encounter any issues or have suggestions for improvements, feel free to open an issue or submit a pull request on our [GitHub repository](https://github.com/arung-agamani/michiru-ch).

Happy deploying!
