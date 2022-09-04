package main

type quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

var Quotes = []quote{
	{"Welcome to the HEV Mark 4 Protective System, for use in hazardous environment conditions.", "HEV Suit"},
	{"High-impact reactive armor activated.", "HEV Suit"},
	{"Atmospheric contaminant sensors activated.", "HEV Suit"},
	{"Vital signs monitoring activated.", "HEV Suit"},
	{"Automatic medical systems engaged.", "HEV Suit"},
	{"Defensive weapon selection system activated.", "HEV Suit"},
	{"Munition level monitoring activated.", "HEV Suit"},
	{"Communications interface online.", "HEV Suit"},
	{"Have a very safe day.", "HEV Suit"},
	{"Minor fracture detected.", "HEV Suit"},
	{"Morphine administered.", "HEV Suit"},

	{"", "Gordon Freeman"}, // Yes guys, I did that.
	{"Have you heard? They are killing scientists... our own race, turning against us.", "Unnamed scientist"},
	{"Gordon doesn't need to hear all this, he's a highly trained professional!", "Unnamed scientist"}, // Of course he is.
	{"For god's sake open the silo doors!", "Unnamed scientist"},
	{"Oh, my GOD, WE'RE DOOMED!", "Unnamed scientist"},

	{"Morning Mr.Freeman, looks like you're running a late", "Barney Calhoun"},

	{"Forget about Freeman, we are cutting our losses and pulling out! Anyone left down there now is on his own.", "HECU Captain"},
	{"I didn't sign on for this shit. Monsters, sure, but civilians?", "HECU soldier"},
	{"All I know for sure is he's been killing my buddies", "HECU soldier"},
	{"Body? What body?", "HECU soldier"},

	{"I have recommended your services to my … employers, and they have authorized me to offer you a job.", "G-Man"},
	{"Quite a nasty piece of work you managed over there, I am impressed.", "G-Man"},
	{"Gordon Freeman, in the flesh - or rather, in the hazard suit.", "G-Man"},
}
