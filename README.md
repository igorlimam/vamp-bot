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

Each implemented entity follows the same CRUD structure, where there is a slash command for each of the operations. Commands outside of reading are reserved to the guild owner.

The language used in the slash commands is Brazilian Portuguese (since the original targeted users are all brazilians), feel free to fork and make one in any other language desired.

The following entities are implemented:
> Clan, Discipline, Info, Merit, Power, Skill.

### Create

Each create command opens up a modal to write the data regarding the specifics of the entity. In the case of dependencies, such as the discipline of a power or the disciplines of a clan, a selection component will be shown **before** the modal.

Ex:
```
/add-clan
|
"Escolha pelo menos uma disciplina para o novo clã"
|
~Disciplines selected~
|
Opens modal
|
"Disciplina cadastrada com sucesso"
```

> In the case of Caitiff that don't have their own disciplines, it is advised to create a placeholder discipline for them.

### Read

Each read command functions as an autocomplete component, since some entities may hold more than 25 (discord listing limit) entries.

After selecting the desired entity, an embed message component will open up to show the data related to that specific entity.

Ex:
```
/clan Nosferatu
|
~Opens up embed component containing information about the clan~
```

> Embed components cannot exceed 6000 characters, nor a single paragraph may exceed 1024 characters.

### Update

The update command opens up an autocomplete component in order to select the desired entity. After that, the same process as the create is instanciated, but with the old data placed, ready to be edited.

Ex:
```
/update-clan Nosferatu
|
"Escolha as disciplinas para o novo clã"
|
~Disciplines changed~
|
~opens up filled modal~
|
"Clã atualizado com sucesso"
```

> In the case of disciplines of a clan, at least one change must be made in the selection component to acknowledge the need to go further for the modal. Which means that to update a clan, one must deselect a discipline, change it, then select it again.

### Delete

Delete is made using an autocomplete component, much like Update. However, in this case, after selecting the entity, a confirmation button will appear.

If confirmed, the entity will be deleted, along with all of its dependencies.

Ex:
```
/delete-clan Nosferatu
|
"Tem certeza que deseja deletar o clã **Nosferatu**"
|                           |
~Selects SIM~               ~Selects NÃO~
|                           |
"Clan deletado com sucesso" "Interação cancelada"
```

> In regards to entity dependencies, their deletion will happen in cascade mode. For example, deleting a discipline remove all of its powers as well.
