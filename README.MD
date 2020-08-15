# UI library for creating desktop applications

Core graphics components powered by Ebiten (fantastic lib!!!)


## Design/thoughts

These will change a lot. I have *NO* experience in designing a UI library (and this will probably show :P )
There are few great looking GO UI libraries but most (not all) are based on CGo. This will be a pure Go library for
those who require it.


## Widgets/Panels:

- Image : DONE
- TextButton : DONE
- ImageButton : DONE
- Checkbox : DONE
- List
- DropboxCombo
- MenuBar
- ToolBar : Basics done.
- TextInput : DONE
- Horizontal Panels : DONE
- Vertical Panels : DONE
- Radio Buttons : Mostly complete
- Add scroll bars to widgets (text input?) or panels?

## Optimisations

- Only load images once and keep common/global. Don't load once per widget!
- Font memory leaks...  keep eye on them.
