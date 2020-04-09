package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	test := []string{"nozzle", "punjabi", "waterlogged", "imprison", "crux", "numismatists", "sultans", "rambles", "deprecating", "aware", "outfield", "marlborough", "guardrooms", "roast", "wattage", "shortcuts", "confidential", "reprint", "foxtrot", "dispossession", "floodgate", "unfriendliest", "semimonthlies", "dwellers", "walkways", "wastrels", "dippers", "engrossing", "undertakings", "unforeseen", "oscilloscopes", "pioneers", "geller", "neglects", "cultivates", "mantegna", "elicit", "couriered", "shielded", "shrew", "heartening", "lucks", "teammates", "jewishness", "documentaries", "subliming", "sultan", "redo", "recopy", "flippancy", "rothko", "conductor", "e", "carolingian", "outmanoeuvres", "gewgaw", "saki", "sarah", "snooping", "hakka", "highness", "mewling", "spender", "blockhead", "detonated", "cognac", "congaing", "prissy", "loathes", "bluebell", "involuntary", "aping", "sadly", "jiving", "buffalo", "chided", "instalment", "boon", "ashikaga", "enigmas", "recommenced", "snell", "parsley", "buns", "abracadabra", "forewomen", "persecuted", "carsick", "janitorial", "neonate", "expeditiously", "porterhouse", "bussed", "charm", "tinseled", "pencils", "inherits", "crew", "estimate", "blacktop", "mythologists", "essequibo", "dusky", "fends", "pithily", "positively", "participants", "brew", "tows", "pentathlon", "misdiagnoses", "paraphrase", "telephoning", "engining", "anglo", "duisburg", "shorthorns", "physical", "enquiries", "grudging", "floodlit", "safflower", "asphalts", "representing", "airbrush", "bedevilling", "fulminations", "peacefuller", "hurl", "unequalled", "wiser", "vinson", "paglia", "doggones", "optimist", "rulering", "katmandu", "flutists", "sterling", "oregonians", "boosts", "slaver", "straightedges", "stendhal", "defaulters", "stylize", "chucking", "adulterate", "partaking", "omelettes", "monochrome", "bitched", "foxhound", "tapir", "vocalizing", "manifolding", "northerner", "ineptly", "dunce", "matchbook", "locutions", "docudrama", "sinkers", "paralegal", "sip", "maliced", "lechers", "zippy", "tillman", "penknives", "olympias", "designates", "mossiest", "leanne", "lavishing", "understate", "underwriting", "showered", "belittle", "propounded", "gristly", "toxicity", "trike", "baudelaire", "sheers", "annmarie", "poultices", "therapeutics", "inputs", "bailed", "minutest", "pynchon", "jinx", "jackets", "subsections", "harmonizes", "caesareans", "freshened", "haring", "disruption", "buckle", "per", "pined", "solemnity", "recombined", "chamber", "tangling", "pitiful", "authoritarians", "oort", "ingratiate", "refreshed", "bavarian", "generically", "rescheduled", "typewritten", "level", "magnetism", "socialists", "oligocene", "resentful", "lambast", "counteroffer", "firefight", "phil", "attenuates", "teary", "demarcated", "moralities", "electrified", "pettiness", "unpacking", "hungary", "heavies", "tenancies", "tirade", "solaria", "scarcity", "prettiest", "carrillo", "yodel", "cantilever", "ridiculously", "tagalog", "schismatics", "ossification", "hezbollah", "downscaling", "calking", "tapped", "girl", "alba", "lavishness", "stepparents", "integrator", "overact", "father", "fobbing", "pb", "require", "toes", "sweats", "prisoners", "mbabane", "hatch", "motleyer", "worlds", "decentralize", "ingrained", "shekels", "directorship", "negotiating", "hiawatha", "busying", "reciprocate", "spinsterhood", "supervened", "scrimmage", "decolonized", "buildups", "sedative", "swats", "despotic", "driblets", "redoubting", "stoic", "xeroxes", "satellited", "exteriors", "deregulates", "lawful", "flunk", "broached", "energetics", "moodily", "popinjays", "shoshone", "misleads", "abduct", "nonevent", "flees", "harry", "cleverness", "manipulative", "shoplifts", "tom", "junk", "poniard", "transmute", "stricter", "trochees", "snack", "relations", "edger", "culminate", "implication", "carjacked", "kissers", "federate", "offsetting", "sutures", "wakened", "axis", "boxcars", "grinds", "scenting", "cordoba", "lumberyards", "incendiary", "antiphonal", "decipherable", "gilliam", "redecorates", "plum", "nitpickers", "linefeed", "awakes", "embittering", "spewing", "annul", "filial", "scarlet", "connors", "sanctum", "scotsman", "isobar", "activity", "overthrowing", "unseasoned", "tarantulae", "outtake", "diego", "mars", "stunted", "hunted", "sublimation", "barbadian", "barbarisms", "epidemic", "assesses", "imposture", "freaks", "detroit", "bloodiest", "avails", "prenatal", "connecticut", "guardsmen", "betwixt", "windsock", "neutralized", "psychoanalysis", "rubberized", "overproduces", "narcissism", "tallow", "cringes", "resinous", "paintbrushes", "duality", "paints", "deactivated", "expertly", "speedsters", "coward", "bass", "psychiatrist", "curies", "betrays", "bubble", "mellow", "showy", "retarding", "radishes", "coy", "unreservedly", "larks", "apportioned", "flaccid", "relabelling", "alphabeted", "anointment", "helms", "gillian", "trophying", "breakage", "underbrush", "directest", "wingtips", "pretence", "preshrink", "remarries", "addle", "brouhaha", "mobbing", "g", "dings", "gains", "stockade", "ouch", "particulates", "listens", "habituation", "kill", "crouped", "hyperbolae", "hutching", "stoney", "rightness", "davids", "questioned", "ethiopians", "courtliness", "delays", "navahos", "devils", "keeling", "accelerators", "investigator", "spindling", "illegality", "extremer", "revlon", "purity", "bradly", "jut", "machs", "liquidated", "informant", "smartly", "disfigure", "parliaments", "croup", "teletypes", "impression", "trainee", "implications", "embed", "podiatrists", "jewelled", "brokenhearted", "spaceman", "unsteadier", "kitchen", "twirling", "conurbations", "pygmies", "lourdes", "watertight", "reassessing", "dempsey", "matriarch", "alas", "abscissae", "decanter", "commentated", "sandy", "idler", "soybean", "cutoff", "dictate", "credibility", "indeterminable", "release", "blank", "curitiba", "pericardia", "probably", "indisposition", "hesitantly", "duff", "ratty", "derivative", "decals", "explication", "cockier", "monoxides", "hyperventilate", "genially", "polluter", "divan", "may", "convalesces", "morpheme", "pupa", "prospered", "tagging", "nerdier", "detached", "spearing", "hilbert", "russeted", "amanuensis", "periwinkles", "jujube", "guarantors", "premises", "descanting", "baned", "deviance", "sidearms", "lamentable", "barristers", "climes", "succulence", "mauve", "oceanographers", "migraine", "bisexual", "peruvians", "fatheads", "parsimony", "pilaf", "portly", "conniving", "insidiously", "inventing", "constabulary", "cling", "stunting", "accessioned", "deadliness", "overthrow", "expectorant", "agamemnon", "blindfold", "striker", "shrugging", "jibes", "appropriateness", "annihilates", "hairs", "proselytes", "goads", "rankling", "predominated", "hart", "enemies", "culpability", "drank", "martinets", "prospering", "dominick", "complemented", "invention", "foolscap", "tolerances", "lorelei", "profits", "awarer", "ungodlier", "victoriously", "mistrusts", "princes", "drudge", "moderator", "transversed", "disco", "japed", "loomed", "incense", "ponds", "gumbel", "disarranges", "coaxes", "technology", "hyde", "debentures", "confidantes", "hankered", "savant", "styes", "croupy", "unapproachable", "cisterns", "unto", "duds", "conglomerating", "clio", "negroid", "looked", "methodism", "hilario", "balloon", "thesauruses", "integuments", "thermometer", "slacks", "wigwagging", "gaping", "incensed", "misquotes", "chocking", "patrols", "upcoming", "insists", "livings", "thoth", "uselessness", "vibrated", "potluck", "starboard", "uniquer", "boone", "scintillates", "darker", "massey", "arbitrariness", "miniaturized", "rousseau", "chiffon", "consortia", "coral", "finesses", "half", "biked", "unlikeliest", "hilarious", "acrid", "twinkles", "galileo", "outsmarted", "ostentation", "cradle", "frats", "misidentifies", "uncleaner", "bacardi", "smoothest", "antwan", "warren", "jingling", "stocks", "daumier", "paranoid", "pantaloons", "dishing", "receive", "underpays", "kane", "variation", "beset", "disobliged", "dreadlocks", "psychics", "twofers", "lieutenants", "pebbling", "interposes", "shingles", "profanes", "machining", "dysfunctions", "wolfram", "brut", "nebraskan", "truculently", "copeland", "devonian", "fuller", "silvia", "philosophers", "cali", "adores", "disquiet", "savvies", "minatory", "blooms", "radiotelephones", "paradoxically", "competitions", "gandhi", "weddell", "occludes", "retracing", "kari", "dead", "lagoons", "menfolk", "abundant", "enacts", "conferences", "procreation", "steadier", "cols", "rabble", "unquestioned", "stupefying", "whodunit", "dizzier", "paula", "riposte", "elixirs", "discontented", "zimbabweans", "assemblage", "unsalted", "genders", "caldwell", "pulleys", "pureness", "kingston", "vests", "hearken", "abuses", "scull", "hussar", "solace", "gondwanaland", "surfacing", "vivienne", "subculture", "reciprocal", "expediencies", "projectiles", "segregationist", "prickle", "pooped", "telecommutes", "axes", "scenery", "peppery", "parenthesizing", "checked", "trademarked", "unreasonable", "curtly", "dynamically", "vulcanize", "airtight", "blotch", "edmund", "stoicism", "scrambles", "whirled", "chasing", "millstones", "helplessly", "permalloy", "remanding", "duplicate", "broadsided", "readjustment", "buggers", "quaked", "grapples", "democrats", "landfalls", "apprehensively", "turmoiling", "railing", "lothario", "modishly", "faggoted", "deflecting", "interment", "dodo", "recreants", "baywatch", "frescoes", "temblor", "brigade", "handgun", "bradstreet", "caspar", "godsend", "cochleae", "queered", "unevenness", "hairnet", "millimeters", "flawless", "plumbing", "disciplinarian", "orbiting", "foothill", "serviettes", "peseta", "windmills", "myrdal", "provides", "slowdowns", "clouting", "gainsays", "dishpans", "mediates", "weaker", "shoestrings", "gerunds", "potsdam", "chips", "disqualifications", "focus", "quarry", "dwarfs", "laurels", "coverall", "reconsidered", "exploded", "distending", "bronzes", "apollonian", "sweeper", "couperin", "gourmets", "irreconcilable", "goldbricking", "emotes", "demilitarizes", "lambkin", "grouper", "anyways", "hugs", "quizzed", "misstatement", "spectrums", "frigates", "plenipotentiaries", "parasites", "tacitly", "savvying", "treks", "dissociating", "departing", "resins", "psychiatric", "tablespoonfuls", "aught", "makeup", "copping", "interwove", "selling", "fantasize", "flamingos", "smolders", "stript", "laverne", "extremely", "chattering", "imminent", "vaulting", "slackly", "pasteurizes", "goody", "pearls", "conceptualization", "fins", "brogues", "muskogee", "naziism", "stromboli", "sunflower", "tosca", "luridness", "booing", "zaniness", "creel", "bootblacks", "attendants", "swordplay", "impinging", "premiere", "sedimentation", "traffickers", "carrel", "observatories", "telltales", "cuckolded", "ampler", "alternately", "shovel", "tasmania", "whooping", "methodologies", "pickling", "overseer"}

	groupAnagrams(test)
}

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	m := make(map[string]int)
	var sorted []string
	temp := ""
	for _, str := range strs {

		sorted = StringToRuneSlice(str)
		sort.Strings(sorted)
		temp = strings.Join(sorted, "")
		v, ok := m[temp]
		if !ok {
			m[string(temp)] = len(m)
			res = append(res, []string{str})

		} else {
			res[v] = append(res[v], str)
		}
	}
	fmt.Println(res)
	return res
}

// type ByRune []rune

// func (r ByRune) Len() int           { return len(r) }
// func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
// func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []string {
	var r []string
	for _, runeValue := range s {
		r = append(r, string(runeValue))
	}
	return r
}

// func IsAnagram(s1, s2 string) bool {
// 	var r1 ByRune = StringToRuneSlice(s1)
// 	var r2 ByRune = StringToRuneSlice(s2)

// 	sort.Sort(r1)
// 	sort.Sort(r2)

// 	return string(r1) == string(r2)
// }

// func groupAnagrams(strs []string) [][]string {
// 	var res [][]string
// 	var m map[string]int
// 	var sorted string

// 	appFlag := 0
// 	for i, str := range strs {
// 		if i == 0 {
// 			res = append(res, []string{str})
// 		} else {
// 			for i, s := range res {
// 				if IsAnagram(s[0], str) {
// 					res[i] = append(res[i], str)
// 					fmt.Println(i)
// 					appFlag = 1
// 				} else {
// 					continue
// 				}
// 			}
// 			if appFlag == 0 {
// 				res = append(res, []string{str})
// 				fmt.Println(str)
// 			}
// 			appFlag = 0
// 		}
// 		fmt.Println(res)

// 	}
// 	fmt.Println(IsAnagram("12", "3"))

// 	return res
// }
