package main

import tgfun "github.com/Sagleft/tgfun/lib"

func getFunnelScript() tgfun.FunnelScript {
	return tgfun.FunnelScript{
		"/start": tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text:  "Greetings, I am a bot from the Utopia ecosystem, you are lucky to run into me!\nLet's play together! You need to choose one of the two options below.",
				Image: "play.jpg",
				Buttons: []tgfun.MessageButton{
					tgfun.MessageButton{
						Text:          "I don't get it!",
						NextMessageID: "dontget",
					},
					tgfun.MessageButton{
						Text:          "Let's play!",
						NextMessageID: "letsplay",
					},
				},
			},
		},
		"dontget": tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: "I don't understand much either, but I can see the future\nAnd I'm ready to show it to you, heh.",
				Buttons: []tgfun.MessageButton{
					tgfun.MessageButton{
						Text:          "come on!",
						NextMessageID: "comeon",
					},
				},
			},
		},
		"letsplay": tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: "Are you interested? Well, I'll continue.\nShall I show you the future?",
				Buttons: []tgfun.MessageButton{
					tgfun.MessageButton{
						Text:          "hmm, sure!",
						NextMessageID: "comeon",
					},
				},
			},
		},
		"comeon": tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: "For several years now, I... I mean, we've been trying to avoid the fate of a lack of privacy and confidentiality in the inevitable future. And we're doing a great job at it. It's difficult, but we're trying specifically for those who care about protecting themselves. Protect you?",
				Buttons: []tgfun.MessageButton{
					tgfun.MessageButton{
						Text:          "Give me protection!",
						NextMessageID: "getprotect",
					},
					tgfun.MessageButton{
						Text:          "How?",
						NextMessageID: "getprotect",
					},
				},
			},
		},
		"getprotect": tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: "Just sign up for the Utopia ecosystem. It's a p2p system that will completely preserve your privacy and confidentiality, built from the ground up. You don't need anything to register, no email, no subscription, no advertising or selling services inside.\n\nSounds interesting? I think so, too.",
				Buttons: []tgfun.MessageButton{
					tgfun.MessageButton{
						Text:          "What to do?",
						NextMessageID: "whatodo",
					},
					tgfun.MessageButton{
						Text:          "I want to go there!",
						NextMessageID: "whatodo",
					},
				},
			},
		},
		"whatodo": tgfun.FunnelEvent{
			Message: tgfun.EventMessage{
				Text: "I'll tell you everything!\n" +
					"I'm very glad you care about your personal life.\n" +
					"Here's everything you need to get to Utopia!\n" +
					"I'm waiting for you there right now!",
				Image: "end.jpg",
				Buttons: []tgfun.MessageButton{
					tgfun.MessageButton{
						Text: "Visit Utopia",
						URL:  "https://u.is/",
					},
				},
			},
		},
	}
}
