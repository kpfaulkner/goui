# UI library for creating desktop applications

Based off ebiten


## Design/thoughts

These will change a lot. I have *NO* experience in designing a UI library (and this will probably show :P )
There are few great looking GO UI libraries but most (not all) are based on CGo. This will be a pure Go library for
those who require it.

- Single window (for now)
- Panels are added to windows. ONLY panels are added to windows
- Widgets (buttons, menus, drop down combos etc) are all added to panels
- No customising looks (aka themeing)..  yet.


- When drawing all the widgets, should the main ebiten Draw function just go through and draw all widgets, or should it
  be delgated. For example, the ebiten Draw() draws the window, the window draws the panels (what it knows about),
  each panel draws the widgets inside it etc?  Would it make a difference?
