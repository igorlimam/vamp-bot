# LEIAME MUCHO

Discord bot implemented using Go.

Its main purpose is to streamline the search for quick info such as the system of certain powers, what disciplines a clan have, how certain merits work, etc.

## Disclaimer:
This bot is a fan-made project and is not affiliated with or endorsed by Paradox Interactive or any official publishers of *Vampire: The Masquerade*. The database file included with the bot is not distributed here due to copyright concerns. Users are encouraged to create their own content in accordance with copyright and trademark law.


To open the **sqlite** database `vtmgo.db` use:
```
sqlite3 vtmgo.db
```

> Sqlite can be installed with `apt install sqlite3`

The Entity-relationship diagram for the current bot implementation can be found [here](https://lucid.app/lucidchart/9e82f6a2-1e70-4920-bb38-c13e93a83713/edit?viewport_loc=-68%2C67%2C2217%2C1052%2C0_0&invitationId=inv_44bc8a4c-fd8c-4768-8281-e7191f6c64af) using LucidChart.

## Slash Commands

Commands regarding Disciplines, Powers, Clans and Merits are available in the form of a simple CRUD. Commands outside of reading are reserved to the guild owner.

### Discipline

`/add-disciplina`

Opens up a Modal component to register the proper information regarding the desired disicipline, such as its name, description, kind and masquerade threat.

`/disciplina [disciplina]`

Reads info about the discipline, the Embed component renders the information in a nice and contained block of readable text.

`/update-disciplina [disciplina]`

Updates a previously added discipline.

`/delete-disciplina [disciplina]`

Deletes the chosen disicpline under confirmation of desire to do so.

### Power

`add-poder [disciplina]`

Adds a new power related to a previsouly added discipline. Deletion on discipline also deletes all of the related powers to it.

After selecting the discipline, a modal is openned to place the power's info.

> Some fields will be separated by '|' due to discord's limitation to eight fiels. In those cases there should be an example indicating how to fill the field.

`poder [disciplina] [poder]`

Reads the info about the selected power from the selected discipline. Since some disciplines have more than 25 powers combined, an AutoComplete component was used to help search for the desired power. An Embed component will load the info of the selected power.

`update-poder [disciplina] [poder]`

Updates the selected power.

`delete-poder [disciplina] [poder]`

Deletes the selected power.

### Info

Some mechanical informations, such as how damage is calculated, what constitutes a critical roll, are often relevant to the player. For that, the `info` slash command whas made to hold information regarding several trivia for the players.

`add-info`

Registers a new information, containing the title, such as "Damage", a subject, such as "Gameplay", and a description fitting the info related to the title.

`info [info-title]`

Reads a previously added information in an autocomplete manner.

`update-info [info-title]`

Updates the info.

`delete-info [info-title]`

Deletes the info.