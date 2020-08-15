# UI library for creating desktop applications

Core graphics components powered by [Ebiten](https://ebiten.org/) (fantastic lib!!!)

## Aim

Go is a hugely productive general purpose programming language (IMHO), which primarily seems to sit in the services/server/infrastructure domain. Although not traditionally used in the graphical application domain it seems to be making some decent headway. Now, we could use a graphics library (of which Ebiten to me is the best one out for Go) but that is more about pushing pixels to the screen and handling some core peripherals I/O. It is NOT (IMO) a UI framework targetting application development which would provide the usual set of 'widgets', ie buttons, text entry, menus etc etc. 

There is already a few UI libraries/frameworks available for Go. The most mature one I'm aware of is [Fyne](https://fyne.io/) . It seems very feature rich and has great documentation and various video tutorials. It does appear to rely on cgo (or at least requires gcc). I'm fairly petantic about wanting to avoid cgo where I can (thanks to Dave Cheney and others for implanting that thought process). So although Fyne is probably great for everyone, I personally want stick to something that is pure Go. Hence GoUI (great naming, eh?)


## Design/thoughts

These will change a lot. I have *NO* experience in designing a UI library (and this will probably show :P )

## Documentation
Still lacking...   since the core widgets are still being developed. Best option right now is to look through the examples (see below)

## Examples

There are various examples in the [examples directory](https://github.com/kpfaulkner/goui/tree/master/examples) . As new widgets are created, the plan is to include at least one piece of sample code in the examples directory. Due to the young age of this project the look/feel is still being established so not a lot of effort has been put towards purely aesthetic side of things. This doesn't mean that you cannot customise how items looks. For example ImageButtons can be customised to have their own "pressed" and "non-pressed" states. The default look/themes aren't established yet.


### Calculator 

Simple calculator



<img src="docs\images\calculator.png" alt="calculator" />



Simple widget test
<img src="docs\images\simple1.png" alt="simple1"/>





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


