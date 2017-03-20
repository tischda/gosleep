package main

// Dumb implementation of a progress bar...
// see https://en.wikipedia.org/wiki/Block_Elements for unicode block elements
var bar = []string{
	"\r[\u2591         ]  10%",
	"\r[\u2591\u2591        ]  20%",
	"\r[\u2591\u2591\u2591       ]  30%",
	"\r[\u2591\u2591\u2591\u2591      ]  40%",
	"\r[\u2591\u2591\u2591\u2591\u2591     ]  50%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591    ]  60%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591\u2591   ]  70%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591  ]  80%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591 ]  90%",
	"\r[\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591\u2591] 100%",
}
