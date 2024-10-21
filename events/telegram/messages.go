package telegram

const msgHelp = `I can save and keep you pages. Also I can offer you them to read- 

In order to save the page, just send me al Link to it.

To see the entire list of saved pages, send me the /list command.

To remove a page from your list, send me the /rm pageURL command.

In order to get a random page from your list, send me comnmand /rnd. 
Caution! After that, this page wiLL be removed from your list!`

const msgHello = "Hi there! \n\n" + msgHelp

const (
	msgUnknownCommand = "Unknown command"
	msgNoSavedPages   = "You have no saved pages"
	msgNoSavedPagesRm = "You don't have a saved page: "
	msgSaved          = "Saved!"
	msgRemove         = "Removed!"
	msgAlreadyExists  = "You have already have this page in your list"
)
