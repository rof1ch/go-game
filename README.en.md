# Go-Game 🎮

Go-Game is a small console text-based walker game written in Go. The player moves between locations, explores zones, interacts with objects, fights monsters and discovers new places.

## Languages

- [🇬🇧 English version](README.en.md)
- [🇷🇺 Русская версия](README.md)

## Features 🎲

- Exploration of various locations
- Interaction with items (keys, weapons, potions, artifacts)
- Meetings with NPCs that can help the player
- Fighting monsters
- Opening new zones using found items
- Ability to create your own maps

## Installation and launch 🚀

### Requirements

- Installed [Go](https://go.dev/dl/)

### Installation

1. Clone the repository:

```sh
git clone https://github.com/rof1ch/go-game.git
cd go-game
```

2. Build and launch the game:

```sh
go run cmd/game/main.go
```

Or compile the binary:

```sh
go build -o go-game cmd/game/main.go
./go-game
```

## Configuration 🛠

The game uses JSON files to configure locations, items, NPCs and monsters. You can add your own maps!

### JSON config structure:

```json
[
	{
		"name": "Home",
		"description": "My home",
		"is_open": true,
		"zones": [
			{
				"name": "Kitchen",
				"is_open": true,
				"items": [
					{
						"type": "key",
						"name": "Garden Key",
						"location_name": "Garden"
					},
					{
						"type": "health_potion",
						"name": "Health Potion",
						"health": 20
					}
				],
				"monster": {
					"name": "Orc",
					"health": 50,
					"damage": 10
				}
			},
			{
				"name": "Living Room",
				"items": [
					{
						"type": "weapon",
						"name": "Knife",
						"damage": 10
					}
				],
				"npc": {
					"name": "Soldier",
					"text": "Hello traveler, glad to see you here. Take a knife to fight monsters"
				}
			}
		],
		"locations": [{ "name": "Garden" }]
	}
]
```

### Field descriptions:

- `name` — Name of the location, zone, item, NPC or monster.
- `description` — Brief description of the location.
- `is_open` — Flag that determines whether the location is initially accessible.
- `zones` — Zones within the location containing items, NPCs and monsters.
- `items` — List of items within the zone.
- `type` — Item type (`key`, `weapon`, `health_potion`, `damage_potion`, `artifact`).
- `health` — Amount of restored health (for potions) or monster health.
- `damage` — Damage (for weapons, damage potions or monsters).
- `location_name` — Location that the key opens.
- `monster` — Object describing the monster (`name`, `health`, `damage`).
- `npc` — Object describing the NPC (`name`, `text`).
- `locations` — List of available transitions to other locations. Only the location name is used to link between locations.

### Creating your own map

You can create your own map by adding a new JSON file with locations and zones. Just follow the structure above and upload a new file to the `config` folder with the name `game.json`.

Example of adding a new location:

```json
{
	"name": "Cave",
	"description": "A dark cave full of mysteries",
	"is_open": false,
	"zones": [
		{
			"name": "Cave Depths",
			"items": [
				{
					"type": "artifact",
					"name": "Stone of Wisdom"
				}
			]
		}
	],
	"locations": [{ "name": "Home" }]
}
```

## Basic commands 🎮

- `help` — Call the menu with commands
- `quit / exit` — Close the game
- `go` — Go to the specified location
- `attack` — Attack the monster
- `take` — Pick up the item
- `inventory` — Open the inventory
- `location` — Display information about the current location
- `me` — Display information about the character

## How to play? 🎮

1. Launch the game.
2. Follow the text instructions.
3. Move between locations, explore zones and interact with objects.
4. Interact with NPCs to get hints or items.
5. Fight monsters, improve your character!

## License 📚

This project is distributed under the MIT license. See [LICENSE](https://create.mit-license.org/) for details.
