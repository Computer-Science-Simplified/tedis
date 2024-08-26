# Tedis
A toy in-memory key-value store for trees.

## Data types

**Binary search trees**
- `BSTADD <key> <value>` adds _value_ to the tree stored at _key_.
- `BSTEXISTS <key> <value>` returns _true_ if _value_ exists in the tree stored at _key_. False otherwise.
- `BSTGETALL <key>` returns all values from the tree at _key_. 
- `BSTREM <key> <value>` removes a value from the tree stored at _key_.

**Binary trees**
- `BTADD <key> <value>` adds _value_ to the tree stored at _key_.
- `BTEXISTS <key> <value>` returns _true_ if _value_ exists in the tree stored at _key_. False otherwise.
- `BTGETALL <key>` returns all values from the tree at _key_.
- `BTREM <key> <value>` removes a value from the tree stored at _key_.