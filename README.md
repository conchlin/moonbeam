# Moonbeam

Moonbeam is a small Discord app written in Go, primarily used to organize parties within gaming communities.

## Guild Event Feed

The primary function of the Moonbeam app is to transform a Discord server into a live feed of Guild events. We achieve this by making timed API calls and generating posts based on character changes. This feature provides a convenient way to stay updated on in-game events directly from Discord!

![Moonbeam feed](https://i.imgur.com/xSrhlPz.png)

## Party Organization

Moonbeam utilizes a command-line interface to organize gaming events. The available commands are as follows:

![Moonbeam commands](https://i.imgur.com/4WjrqX6.png)

Examples:

```bash
$displayparties

# To use createparty, we must specify the type, date, and time.
# Date accepts the structure 2023-10-28 or the today keyword.
# Time accepts a UTC+0 time.
$createparty zakum today 8:30am

# To use joinparty, we must specify the character and party ID.
# Party ID can be found by using showparties.
$joinparty 1 GameCharacter

# Party ID can be found by using showparties.
$expel 1 GameCharacter

# We must use the Discord name instead of the character name.
$invite DiscordName 1

# You must be the party creator to use this command.
$deleteparty 1
```

## Miscellaneous Commands

```bash
# Chooses one of the entries at random.
$random 1 2 3

# Moonbeam uses the MapleLegends API to pull character data.
# This way, you can track your character stats from Discord.
$character CharacterName
```

![character command](https://i.imgur.com/KoQXrsS.png)