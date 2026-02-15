# Pokedex CLI
A high-performance, command-line interactive Pokemon encyclopedia built with Go. This isn't just a list; it's a fully functional exploration tool where you can navigate the Pokemon world, encounter wild creatures, and build your own Pokedex.

## Features
- **Map Exploration:** Browse areas with `map` and `mapb`.
- **Catch & Inspect:** Catch Pokemon and view their stats.
- **In-Memory Cache:** Fast responses with a custom cache system.

## Commands

| Command | Description |
| :--- | :--- |
| **help** | Displays all available commands |
| **map** | Shows next 20 location areas |
| **mapb** | Shows previous 20 location areas |
| **explore** | Lists Pokemon in a specific area |
| **catch** | Attempts to catch a Pokemon |
| **inspect** | Shows stats of a caught Pokemon |
| **pokedex** | Lists your collection |
| **exit** | Closes the application |

## Technical Details
- **Mutex:** Thread-safe cache operations.
- **Ticker:** Background cleanup for expired cache data.
- **PokeAPI:** Real-time data integration.

## Project Structure
```text
.
├── command_catch.go
├── command_exit.go
├── command_explore.go
├── command_help.go
├── command_inspect.go
├── command_map.go
├── command_pokedex.go
├── go.mod
├── internal
│   ├── pokeapi
│   │   ├── client.go
│   │   ├── location_get.go
│   │   ├── location_list.go
│   │   ├── pokeapi.go
│   │   ├── pokemon_get.go
│   │   ├── types_locations.go
│   │   └── types_pokemons.go
│   └── pokecache
│       ├── pokecache_test.go
│       └── pokecache.go
├── main.go
├── repl_test.go
└── repl.go
```
