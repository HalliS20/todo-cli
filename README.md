# Todo CLI

Todo List program for the CLI.
Developed using go and Charmbracelets Bubbletea library

## Controls

### Base

- o : Add
- i : Edit
- d : Delete
- q : Quit

### Nav

- J/Down : Down
- k/Up : Up
- ctrl+j : moves item down
- ctrl+k : moves item up

### Specific

#### List View

- Enter : enters list

#### Todo View

- Enter : marks done
- "-" : Navigates to List view

#### Editing/adding

- ESC : cancels
- Enter : finishes adding

- ### Extra

- ctrl+c : Quits anywhere anytime
- backspace : also deltes

## Current Status

- Supports all operations adequately

## Future Work

By order of implementation

### Combining list and Todo types

- allows easily implementing subLists
- makes code neater and faster

### allow word scrolling when editing

- requires registering left/h and right/l correctly
- as well as small work on the look of the view for a cursor

### Cleaner controls view

- as is all controls are in a footer
- while some may be acceptable when i've added more it may be a bit too much

### customizable controls using a config file

- should not be too difficult since all bindings are made via strings
- its more of a personalization feature so it comes after functionality
