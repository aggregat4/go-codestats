Starting from a base directory, gather code LOC statistics recursively and visualize the results in a navigable treemap. Allow filtering by language and type of line.

Solution sketch:

- Start with root folder and walk all subdirs concurrently and recursively
- Basic datastructure to generate:
    - Keep stats for each folder: name, children,
    - Keep stats on each file: name, language enum, lines code, lines blank, lines comment 
- Caclculate stats on the fly always (depending on language filter and line type filter)
