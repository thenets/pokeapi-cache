# PokeAPI Cache

PokeAPI-Cache is up to 4x faster than access PokeAPI directly.

## How to use

It's the same as PokeAPI official documentation:
- https://pokeapi.co/

You will need to change `pokeapi.co` to `pokeapi-cache.herokuapp.com`. Nothing more.

Examples:
- https://pokeapi-cache.herokuapp.com/api/v2/pokemon/pikachu
- https://pokeapi-cache.herokuapp.com/api/v2/move/34

## Benchmark

### PokeAPI

```python
import datetime as dt
import requests

def pokeapiBenchmark():
  for i in range(500):
    url = "https://pokeapi-cache.herokuapp.com/api/v2/move/%d" % (i+1)
    r = requests.get(url)
    
t_before = dt.datetime.now()
pokeapiBenchmark()
t_after = dt.datetime.now()

print(t_after - t_before)
```

Result: `0:04:12.039194`

### PokeAPI-Cache

```python
import datetime as dt
import requests

def pokeapiBenchmark():
  for i in range(500):
    url = "https://pokeapi-cache.herokuapp.com/api/v2/move/%d" % (i+1)
    r = requests.get(url)
    
t_before = dt.datetime.now()
pokeapiBenchmark()
t_after = dt.datetime.now()

print(t_after - t_before)
```

Result: `0:01:22.681270`

## Credits

Thanks for everyone from [PokeAPI](https://github.com/PokeAPI/pokeapi) community! 💖

## License

MIT
